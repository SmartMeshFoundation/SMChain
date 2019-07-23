package tribe

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/SmartMeshFoundation/Spectrum/core/state"

	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/SmartMeshFoundation/Spectrum/crypto"
	"github.com/SmartMeshFoundation/Spectrum/log"
	"github.com/SmartMeshFoundation/Spectrum/params"
)

func NewTribeStatus() *TribeStatus {
	ts := &TribeStatus{
		Signers:     make([]*Signer, 0),
		SignerLevel: LevelNone,
	}
	return ts
}

func (self *TribeStatus) setTribe(tribe *Tribe) {
	self.tribe = tribe
}

func (self *TribeStatus) getNodekey() *ecdsa.PrivateKey {
	if self.nodeKey == nil {
		panic(errors.New("GetNodekey but nodekey not ready"))
	}
	return self.nodeKey
}

func (self *TribeStatus) GetMinerAddress() common.Address {
	if self.nodeKey == nil {
		panic(errors.New("GetMinerAddress but nodekey not ready"))
	}
	pub := self.nodeKey.PublicKey
	add := crypto.PubkeyToAddress(pub)
	return add
}

func (self *TribeStatus) GetMinerAddressByChan(rtn chan common.Address) {
	go func() {
		for {
			if self.nodeKey != nil && self.tribe.isInit {
				break
			}
			<-time.After(time.Second)
		}
		pub := self.nodeKey.PublicKey
		rtn <- crypto.PubkeyToAddress(pub)
	}()
}

func (self *TribeStatus) GetSignersFromChiefByHash(hash common.Hash, number *big.Int) ([]*Signer, error) {
	rtn := params.SendToMsgBoxWithHash("GetStatus", hash, number)
	r := <-rtn
	if !r.Success {
		return nil, r.Entity.(error)
	}
	cs := r.Entity.(params.ChiefStatus)
	signers := cs.SignerList
	scores := cs.ScoreList
	sl := make([]*Signer, 0, len(signers))
	for i, signer := range signers {
		score := scores[i]
		sl = append(sl, &Signer{signer, score.Int64()})
	}
	return sl, nil
}

// 在 加载完所有 node.service 后，需要主动调用一次
func (self *TribeStatus) LoadSignersFromChief(hash common.Hash, number *big.Int) error {
	//log.Info(fmt.Sprintf("LoadSignersFromChief hash=%s,number=%s", hash.String(), number))
	cs, err := params.TribeGetStatus(number, hash)
	if err != nil {
		log.Warn("TribeGetStatusError", "err", err, "num", number, "hash", hash.Hex())
		return err
	}
	signers := cs.SignerList
	scores := cs.ScoreList
	sl := make([]*Signer, 0, len(signers))
	for i, signer := range signers {
		score := scores[i]
		sl = append(sl, &Signer{signer, score.Int64()})
	}
	self.LeaderLimit = cs.LeaderLimit
	self.Leaders = cs.LeaderList
	if len(self.Leaders) == 0 {
		panic(fmt.Sprintf("LoadSignersFromChief err ,hash=%s,number=%s,cs=%#v", hash.String(), number, cs))
	}
	self.Number = cs.Number.Int64()
	self.blackList = cs.BlackList
	err = self.loadSigners(sl)
	if err != nil {
		return err
	}
	self.Epoch, self.SignerLimit = cs.Epoch, cs.SignerLimit
	go self.resetSignersLevel(hash, number)
	return nil
}

func (self *TribeStatus) resetSignersLevel(hash common.Hash, number *big.Int) {
	m := self.GetMinerAddress()
	for _, s := range self.Signers {
		if s.Address == m {
			self.SignerLevel = LevelSigner
			return
		}
	}
	for _, s := range self.blackList {
		if s == m {
			self.SignerLevel = LevelSinner
			return
		}
	}

	for _, s := range self.Leaders {
		if s == m {
			self.SignerLevel = LevelSigner
			return
		}
	}

	ci := params.GetChiefInfo(number)
	switch ci.Version {
	case "0.0.6":
		// if filterVolunteer return 1 then is volunteer
		rtn := params.SendToMsgBoxForFilterVolunteer(hash, number, m)
		r := <-rtn
		if r.Success {
			if fr := r.Entity.(*big.Int); fr != nil && fr.Int64() == 0 {
				self.SignerLevel = LevelVolunteer
				return
			}
		}
	}
	// default none
	self.SignerLevel = LevelNone
}

func (self *TribeStatus) loadSigners(sl []*Signer) error {
	self.Signers = append(self.Signers[:0], sl...)
	return nil
}

func (self *TribeStatus) GetSigners() []*Signer {
	return self.Signers
}

//InTurnForCalcChief100 计算规则参考inTurnForCalcChief100
func (self *TribeStatus) InTurnForCalcChief100(signer common.Address, parent *types.Header) *big.Int {
	return self.inTurnForCalcChief100(parent.Number.Int64()+1, parent.Hash(), signer)
}

/*
inTurnForCalcChief100 计算如果当前出块节点是signer的话,它对应的难度是多少.
signers:[0,...,16] 0号对应的是常委会节点,1-16对应的是普通出块节点
场景1:
1. 当前应该出块节点应该是3,如果signer是3,那么难度就是6.
2. 如果singers[0]对应的是常委2, 这时候常委2出块,难度是5,常委3出块难度是4,...,常委1出块难度则是1
场景2:当前出块节点应该是singers[0],也就是某个常委会节点
1. 如果signers[0] 出块,那么难度就是6
2. 假设signers[0]是常委2,那么常委3替他出块难度是5,常委4出块就是4,...常委1出块难度则是2

这里的number参数主要是选定合约版本,而parentHash则是用来选择读取哪个block时候的合约状态
*/
func (self *TribeStatus) inTurnForCalcChief100(number int64, parentHash common.Hash, signer common.Address) *big.Int {
	var (
		signers, _ = self.GetSignersFromChiefByHash(parentHash, big.NewInt(number)) //self.GetSigners()
		sl         = len(signers)
	)
	//	log.Info(fmt.Sprintf("singers=%v,signer=%s,leaders=%v,number=%d,parentHash=%s", signers, signer.String(), self.Leaders, number, parentHash.String()))
	if idx, _, err := self.fetchOnSigners(signer, signers); err == nil {
		// main
		if sl > 0 && number%int64(sl) == idx.Int64() {
			return big.NewInt(diff)
		}
		// second
		if idx.Int64() == 0 {
			return big.NewInt(diff - 1)
		}

	} else if sl > 0 {
		if leaders, err := leaderSort(signers[0].Address, self.Leaders); err == nil {
			for i, leader := range leaders {
				if signer == leader && number%int64(sl) == 0 {
					return big.NewInt(diff - int64(i+1))
				} else if signer == leader {
					return big.NewInt(diff - int64(i+2))
				}
			}
		}
	}
	return diffNoTurn
}

//InTurnForVerifyChief100: 计算规则参考inTurnForCalcChief100
func (self *TribeStatus) InTurnForVerifyChief100(number int64, parentHash common.Hash, signer common.Address) *big.Int {
	return self.inTurnForCalcChief100(number, parentHash, signer)
}

/*
假设list=[1,2,3,4,5]
first=3,那么返回[4,5,1,2]
如果first=2,返回[3,4,5,1]
如果first=5,返回[1,2,3,4]
*/
func leaderSort(first common.Address, list []common.Address) ([]common.Address, error) {
	for i, o := range list {
		if first == o {
			return append(list[i+1:], list[:i]...), nil
		}
	}

	return list, nil
}

//InTurnForCalc 在0.6版本计算难度
func (self *TribeStatus) InTurnForCalc(signer common.Address, parent *types.Header) *big.Int {
	number := parent.Number.Int64() + 1
	signers := self.GetSigners()
	if idx, _, err := self.fetchOnSigners(signer, signers); err == nil {
		sl := len(signers)
		if params.IsSIP002Block(big.NewInt(number)) {
			if sl > 0 && number%int64(sl) == idx.Int64() {
				return diffInTurnMain
			} else if sl > 0 && (number+1)%int64(sl) == idx.Int64() {
				return diffInTurn
			}
		} else {
			if sl > 0 && number%int64(sl) == idx.Int64() {
				return diffInTurn
			}
		}
	}

	return diffNoTurn
}

func (self *TribeStatus) InTurnForVerify(number int64, parentHash common.Hash, signer common.Address) *big.Int {

	if ci := params.GetChiefInfo(big.NewInt(number)); ci != nil {
		switch ci.Version {
		case "1.0.0":
			//TODO max value is a var ???
			return self.InTurnForVerifyChief100(number, parentHash, signer)
		}
	}

	var signers []*Signer
	if number > 3 {
		var err error
		signers, err = self.GetSignersFromChiefByHash(parentHash, big.NewInt(number))
		if err != nil {
			log.Warn("InTurn:GetSignersFromChiefByNumber:", "err", err)
		}
	} else {
		return diffInTurn
	}
	if idx, _, err := self.fetchOnSigners(signer, signers); err == nil {
		sl := len(signers)
		if params.IsSIP002Block(big.NewInt(number)) {
			if sl > 0 && number%int64(sl) == idx.Int64() {
				return diffInTurnMain
			} else if sl > 0 && (number+1)%int64(sl) == idx.Int64() {
				return diffInTurn
			}
		} else {
			if sl > 0 && number%int64(sl) == idx.Int64() {
				return diffInTurn
			}
		}
	}
	return diffNoTurn
}

func (self *TribeStatus) genesisSigner(header *types.Header) (common.Address, error) {
	extraVanity := extraVanityFn(header.Number)
	signer := common.Address{}
	copy(signer[:], header.Extra[extraVanity:])
	self.loadSigners([]*Signer{{signer, 3}})
	return signer, nil
}

//address对应的signer以及其在signers中的下标
func (self *TribeStatus) fetchOnSigners(address common.Address, signers []*Signer) (*big.Int, *Signer, error) {
	if signers == nil {
		signers = self.Signers
	}
	if l := len(signers); l > 0 {
		for i := 0; i < l; i++ {
			if s := signers[i]; s.Address == address {
				return big.NewInt(int64(i)), s, nil
			}
		}
	}
	return nil, nil, errors.New("not_found")
}

// called by end of WriteBlockAndState
// if miner then execute chief.update and chief.getStatus
// else execute chief.getStatus only
func (self *TribeStatus) Update(currentNumber *big.Int, hash common.Hash) {
	if currentNumber.Int64() >= CHIEF_NUMBER && atomic.LoadInt32(&self.mining) == 1 {
		// mining start
		log.Debug("<<TribeStatus.Update_begin>>", "num", currentNumber.Int64())
		success := <-params.SendToMsgBoxWithNumber("Update", currentNumber)
		log.Debug("<<TribeStatus.Update_end>>", "num", currentNumber.Int64(), "success", success.Success, "entity", success.Entity)
		err := self.LoadSignersFromChief(hash, currentNumber)
		if err != nil {
			log.Error("[TribeStatus.Update] LoadSignersFromChief ", "number", currentNumber, "err", err)
		}
	}
}

func verifyVrfNum(parent, header *types.Header) (err error) {
	var (
		np  = header.Extra[:extraVanityFn(header.Number)]
		sig = header.Extra[len(header.Extra)-extraSeal:]
		msg = append(parent.Number.Bytes(), parent.Extra[:32]...)
	)
	pubbuf, err := ecrecoverPubkey(header, sig)
	if err != nil {
		panic(err)
	}
	x, y := elliptic.Unmarshal(crypto.S256(), pubbuf)
	pubkey := ecdsa.PublicKey{crypto.S256(), x, y}
	err = crypto.SimpleVRFVerify(&pubkey, msg, np)
	log.Debug("[verifyVrfNum]", "err", err, "num", header.Number, "vrfn", new(big.Int).SetBytes(np[:32]), "parent", header.ParentHash.Bytes())
	return
}

/*
validateSigner:
1. 验证出块时间符合规则,具体规则见GetPeriodChief100描述
2.
*/
func (self *TribeStatus) validateSigner(parentHeader, header *types.Header, signer common.Address) bool {
	var (
		err        error
		signers    []*Signer
		number     = header.Number.Int64()
		parentHash = header.ParentHash
	)
	//if number > 1 && self.Number != parentNumber {
	if number <= CHIEF_NUMBER {
		return true
	}

	signers, err = self.GetSignersFromChiefByHash(parentHash, big.NewInt(number))
	if err != nil {
		log.Warn("TribeStatus.ValidateSigner : GetSignersFromChiefByNumber :", "err", err)
	}

	if params.IsSIP002Block(header.Number) {

		// second time of verification block time
		period := self.tribe.GetPeriod(header, signers)
		pt := parentHeader.Time.Uint64()
		if pt+period > header.Time.Uint64() {
			log.Error("[ValidateSigner] second time verification block time error", "num", header.Number, "pt", pt, "period", period, ", pt+period=", pt+period, " , ht=", header.Time.Uint64())
			log.Error("[ValidateSigner] second time verification block time error", "err", ErrInvalidTimestampSIP002)
			return false
		}
	}

	if params.IsSIP004Block(header.Number) && header.Coinbase == common.HexToAddress("0x") {
		log.Error("error_signer", "num", header.Number.String(), "miner", header.Coinbase.Hex(), "signer", signer.Hex())
		return false
	}

	idx, _, err := self.fetchOnSigners(signer, signers)
	if params.IsSIP100Block(header.Number) {
		if err == nil {
			// 轮到谁出就谁出的块
			idx_m := number % int64(len(signers))
			if idx_m == idx.Int64() {
				return true
			}
			// 其他只能有常委会节点替代
			if idx.Int64() == 0 {
				return true
			}
		} else {
			// other leader
			for _, leader := range self.Leaders {
				if signer == leader {
					return true
				}
			}
		}
	} else if err == nil {
		return true
	}
	return false
}

/*
VerifySignerBalance: 在chief1.0之前直接通过账号余额来判断是否具有出块资格,chief1.0之后只能通过抵押到poc合约中才具有资格.
*/
func (self *TribeStatus) VerifySignerBalance(state *state.StateDB, addr common.Address, header *types.Header) error {
	// SIP100 skip this verify
	if params.IsSIP100Block(header.Number) {
		return nil
	}
	var (
		pnum, num *big.Int
		f         common.Address
		nl        []common.Address
		err       error
	)
	if addr == common.HexToAddress("0x") {
		if _addr, err := ecrecover(header, self.tribe); err == nil {
			addr = _addr
		} else {
			return err
		}
	}
	if header != nil {
		num = header.Number
		pnum = new(big.Int).Sub(num, big.NewInt(1))
	} else {
		return errors.New("params of header can not be null")
	}
	// skip when v in meshbox.sol
	if params.IsReadyMeshbox(pnum) && params.MeshboxExistAddress(addr) {
		return nil
	}

	if params.IsReadyAnmap(pnum) {
		f, nl, err = params.AnmapBindInfo(addr, header.Hash())
	} else {
		err = errors.New("anmap_not_ready")
	}
	log.Debug("<<VerifySignerBalance_AnmapBindInfo>>", "num", num, "addr", addr.Hex(), "f", f.Hex(), "nl.len", len(nl), "err", err)
	if err == nil && f != common.HexToAddress("0x") && len(nl) > 0 {
		// exclude meshbox n in nl
		noBox := int64(0)
		for _, n := range nl {
			if !params.MeshboxExistAddress(n) {
				noBox++
			}
		}
		if noBox == 0 {
			return nil
		}
		fb := state.GetBalance(f)
		mb := new(big.Int).Mul(params.GetMinMinerBalance(), big.NewInt(noBox))
		log.Debug("<<VerifySignerBalance>> 0 :", "nl.len", len(nl), "nobox", noBox, "fb", fb, "mb", mb)
		//nb := state.GetBalance(n)
		if fb.Cmp(mb) < 0 {
			log.Debug("<<VerifySignerBalance>> 1 :", "f", f.Hex(), "fb", fb, "mb", mb)
			return ErrTribeChiefVolunteerLowBalance
		}
	} else if params.IsSIP004Block(num) {
		b := state.GetBalance(addr)
		if b.Cmp(params.GetMinMinerBalance()) < 0 {
			log.Debug("<<VerifySignerBalance>> 2 :", "n", addr.Hex(), "nb", b, "mb", params.GetMinMinerBalance())
			return ErrTribeChiefVolunteerLowBalance
		}
	}
	return nil
}

// every block
// sync download or mine
// check chief tx
func (self *TribeStatus) ValidateBlock(state *state.StateDB, parent, block *types.Block, validateSigner bool) error {
	if block.Number().Int64() <= 3 {
		return nil
	}
	err := self.LoadSignersFromChief(parent.Hash(), parent.Number())
	if err != nil {
		log.Error(fmt.Sprintf("[ValidateBlock] LoadSignersFromChief ,parent=%s,current=%s,currentNumber=%s", parent.Hash().String(), block.Hash().String(), block.Number()))
		return err
	}
	header := block.Header()
	number := header.Number.Int64()
	//number := block.Number().Int64()
	// add by liangc : seal call this func must skip validate signer
	if validateSigner {
		signer, err := ecrecover(header, self.tribe)
		// verify difficulty
		if number > 3 && !params.IsBeforeChief100block(header.Number) {
			difficulty := self.InTurnForVerify(number, header.ParentHash, signer)
			if difficulty.Cmp(header.Difficulty) != 0 {
				log.Error("** verifySeal ERROR **", "head.diff", header.Difficulty.String(), "target.diff", difficulty.String(), "err", errInvalidDifficulty)
				return errInvalidDifficulty
			}

		}
		// verify signer
		if err != nil {
			return err
		}
		if !self.validateSigner(parent.Header(), header, signer) {
			return errUnauthorized
		}

		// verify vrf num
		if params.IsSIP100Block(header.Number) {
			err = verifyVrfNum(parent.Header(), header)
			if err != nil {
				log.Error("vrf_num_fail", "num", number, "err", err)
				return err
			}
		}

	}
	// check first tx , must be chief.tx , and onely one chief.tx in tx list
	if block != nil && block.Transactions().Len() == 0 {
		return ErrTribeNotAllowEmptyTxList
	}

	// add by liangc 190412 : SIP004 if the sender in signerList now refuse and skip this tx
	signerMap := make(map[common.Address]struct{})
	if params.IsSIP004Block(header.Number) && !params.IsSIP100Block(header.Number) {
		for _, signer := range self.Signers {
			signerMap[signer.Address] = struct{}{}
		}
	}

	var total = 0
	for i, tx := range block.Transactions() {

		if params.IsSIP004Block(header.Number) && !params.IsSIP100Block(header.Number) {
			from := types.GetFromByTx(tx)
			//verify by anmap bindinfo
			_, nl, err := params.AnmapBindInfo(*from, parent.Hash())

			verifyBySignerMap := func(addr common.Address) error {
				if _, ok := signerMap[addr]; i > 0 && ok {
					return ErrTribeValdateTxSenderCannotInSignerList
				}
				return nil
			}
			if err == nil && len(nl) > 0 {
				// exclude meshbox first
				fnl := make([]common.Address, 0)
				for _, n := range nl {
					if !params.MeshboxExistAddress(n) {
						fnl = append(fnl[:], n)
					}
				}
				log.Debug("TODO<<TribeStatus.ValidateBlock>> exclude_meshbox_first", "num", number, "i", i, "from", from.Hex(), "to", tx.To(), "nl.len", len(nl), "fnl.len", len(fnl))
				for _, n := range fnl {
					if err := verifyBySignerMap(n); err != nil {
						return err
					}
				}
			} else {
				if err := verifyBySignerMap(*from); err != nil {
					return err
				}
			}
		}

		if tx.To() != nil && params.IsChiefAddress(*tx.To()) && params.IsChiefUpdate(tx.Data()) {
			//verify volunteer
			if state != nil {
				if params.IsSIP100Block(header.Number) {
					// TODO SIP100 check volunteer by vrfnp
					volunteerHex := common.Bytes2Hex(tx.Data()[4:])
					volunteer := common.HexToAddress(volunteerHex)
					vrfn := header.Extra[:32]
					if !params.VerifyMiner(header.ParentHash, volunteer, vrfn) {
						return errors.New("verify_volunteer_fail")
					}
				} else {
					volunteerHex := common.Bytes2Hex(tx.Data()[4:])
					volunteer := common.HexToAddress(volunteerHex)
					if volunteer != common.HexToAddress("0x") {
						log.Debug("<<TribeStatus.ValidateBlock>> verify_volunteer =>", "num", number, "v", volunteer.Hex())
						if err := self.VerifySignerBalance(state, volunteer, parent.Header()); err != nil {
							return err
						}
					}
				}
			}
			total++
		}
	}
	if total == 0 {
		return ErrTribeMustContainChiefTx
	} else if total > 1 {
		return ErrTribeChiefCannotRepeat
	}

	log.Debug("ValidateBlockp-->", "num", block.NumberU64(), "check_signer", validateSigner)
	return nil
}

func (self *TribeStatus) String() string {
	if b, e := json.Marshal(self); e != nil {
		return "error: " + e.Error()
	} else {
		return "status: " + string(b)
	}
}
