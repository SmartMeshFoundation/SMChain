// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"fmt"
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum/common"
)

var (
	MainnetGenesisHash = common.HexToHash("0xa8ab5ecb308dd061e6baa8c8a04a62b3e35fae7aae1393921f8d52deac2c5712") // enforce below configs on
	TestnetGenesisHash = common.HexToHash("0x28242478c4f01d9208a79d962b7f6383af488565f36aea7038a34020c07db697") // Testnet genesis hash to enforce below configs on
)

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainId:        big.NewInt(20180430),
		HomesteadBlock: big.NewInt(0),
		DAOForkBlock:   nil,
		DAOForkSupport: false,
		EIP150Block:    nil,
		EIP150Hash:     common.Hash{},
		EIP155Block:    big.NewInt(0),
		EIP158Block:    big.NewInt(0),
		ByzantiumBlock: big.NewInt(0),
		// add by liangc : change default consensus for dev
		// if skip this vsn please set 0 or nil to block and set common.Address{} to address
		// 0.0.5 : ready for release
		Chief005Block:   big.NewInt(2),
		Chief005Address: common.HexToAddress("0xf0d2ad4e0d25cfea98e9640329993bbc52396abd"),
		// 0.0.6 : ready for release
		Chief006Block:   big.NewInt(595888),
		Chief006Address: common.HexToAddress("0xba7f507d5aab3e931312512c234fbeb85cbd9dce"),

		MeshboxBlock:   big.NewInt(1870333),
		MeshboxAddress: common.HexToAddress("0xf0ced0b1ce8738eeac06fdca51e0ff398328634b"),

		NR001Block: big.NewInt(0), // new rules for chief.tx of gaspool
		NR002Block: big.NewInt(588888),
		NR003Block: big.NewInt(808888),
		Tribe:      &TribeConfig{},
	}

	// TestnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	TestnetChainConfig = &ChainConfig{
		ChainId:        big.NewInt(3),
		HomesteadBlock: big.NewInt(0),
		DAOForkBlock:   nil,
		DAOForkSupport: true,
		EIP150Block:    nil,
		EIP150Hash:     common.Hash{},
		EIP155Block:    big.NewInt(0),
		EIP158Block:    big.NewInt(0),
		ByzantiumBlock: big.NewInt(62678),
		//ByzantiumBlock: big.NewInt(1700000),
		//Clique: &CliqueConfig{ Period: 15, Epoch:  30000}, //Ethash: new(EthashConfig),
		// if skip this vsn please set 0 or nil to block and set common.Address{} to address
		// 0.0.2
		Chief002Block:   big.NewInt(2),
		Chief002Address: common.HexToAddress("0x9ec55c1dafd4a487e41da33e344aef86da41ab82"),
		// TODO : set chief 0.0.3 in chain config
		Chief003Block:   big.NewInt(113772), //hard fork testnet
		Chief003Address: common.HexToAddress("0xac28e532b3fac82554fc7b0b8b62549deeeb33a9"),
		// 0.0.4 : fix bug redeploy
		Chief004Block:   big.NewInt(120305), //hard fork testnet : fix bugs and debug chief
		Chief004Address: common.HexToAddress("0xe242e2bcf5b0da6518320210fab0a27458bc0674"),
		// 0.0.5 : ready for release
		Chief005Block:   big.NewInt(176244),
		Chief005Address: common.HexToAddress("0xe90da8175922925dfb40e6505b508f1042e807aa"),
		// 0.0.6 : ready for release
		Chief006Block:   big.NewInt(595710),
		Chief006Address: common.HexToAddress("0x53cb83888e6d28cf7ec168308c65172001f441aa"),

		MeshboxBlock:   big.NewInt(1976666),
		MeshboxAddress: common.HexToAddress("0x0f15e1e44322b2946215705d2ed60cba899f0b38"),

		// new rules for chief.tx of gaspool
		NR001Block: big.NewInt(176222),
		NR002Block: big.NewInt(525233),
		NR003Block: big.NewInt(917013),
		Tribe:      &TribeConfig{},
	}

	// RinkebyChainConfig contains the chain parameters to run a node on the Rinkeby test network.
	RinkebyChainConfig = &ChainConfig{
		ChainId:        big.NewInt(4),
		HomesteadBlock: big.NewInt(1),
		DAOForkBlock:   nil,
		DAOForkSupport: true,
		EIP150Block:    big.NewInt(2),
		EIP150Hash:     common.HexToHash("0x9b095b36c15eaf13044373aef8ee0bd3a382a5abb92e402afa44b8249c3a90e9"),
		EIP155Block:    big.NewInt(3),
		EIP158Block:    big.NewInt(3),
		ByzantiumBlock: big.NewInt(1035301),
		Clique: &CliqueConfig{
			Period: 15,
			Epoch:  30000,
		},
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Ethash consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{
		ChainId:        big.NewInt(1337),
		HomesteadBlock: big.NewInt(0),
		DAOForkBlock:   nil,
		DAOForkSupport: false,
		EIP150Block:    big.NewInt(0),
		EIP150Hash:     common.Hash{},
		EIP155Block:    big.NewInt(0),
		EIP158Block:    big.NewInt(0),
		ByzantiumBlock: big.NewInt(0),
		Ethash:         new(EthashConfig),
		Clique:         nil,
		Tribe:          nil,
	}

	// AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Clique consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllCliqueProtocolChanges = &ChainConfig{
		ChainId:        big.NewInt(1337),
		HomesteadBlock: big.NewInt(0),
		DAOForkBlock:   nil,
		DAOForkSupport: false,
		EIP150Block:    big.NewInt(0),
		EIP150Hash:     common.Hash{},
		EIP155Block:    big.NewInt(0),
		EIP158Block:    big.NewInt(0),
		ByzantiumBlock: big.NewInt(0),
		Ethash:         nil,
		Clique:         &CliqueConfig{Period: 0, Epoch: 30000},
		Tribe:          nil,
	}
	TestChainConfig = &ChainConfig{
		ChainId:        big.NewInt(1),
		HomesteadBlock: big.NewInt(0),
		DAOForkBlock:   nil,
		DAOForkSupport: false,
		EIP150Block:    big.NewInt(0),
		EIP150Hash:     common.Hash{},
		EIP155Block:    big.NewInt(0),
		EIP158Block:    big.NewInt(0),
		ByzantiumBlock: big.NewInt(0),
		Ethash:         new(EthashConfig),
		Clique:         nil,
		Tribe:          nil,
	}
	TestRules = TestChainConfig.Rules(new(big.Int))
)

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainId *big.Int `json:"chainId"` // Chain id identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock *big.Int `json:"byzantiumBlock,omitempty"` // Byzantium switch block (nil = no fork, 0 = already on byzantium)

	// Various consensus engines
	Ethash *EthashConfig `json:"ethash,omitempty"`
	Clique *CliqueConfig `json:"clique,omitempty"`
	Tribe  *TribeConfig  `json:"tribe,omitempty"` // add by liangc

	// >>> add by liangc : set chief start number >>>
	// chief.sol vsn 0.0.2
	Chief002Block   *big.Int       `json:"chief002Block,omitempty"`
	Chief002Address common.Address `json:"chief002Address,omitempty"`
	// chief.sol vsn 0.0.3
	Chief003Block   *big.Int       `json:"chief003Block,omitempty"`
	Chief003Address common.Address `json:"chief003Address,omitempty"`
	// chief.sol vsn 0.0.4
	Chief004Block   *big.Int       `json:"chief004Block,omitempty"`
	Chief004Address common.Address `json:"chief004Address,omitempty"`
	// chief.sol vsn 0.0.5
	Chief005Block   *big.Int       `json:"chief005Block,omitempty"`
	Chief005Address common.Address `json:"chief005Address,omitempty"`
	// chief.sol vsn 0.0.6
	Chief006Block   *big.Int       `json:"chief006Block,omitempty"`
	Chief006Address common.Address `json:"chief006Address,omitempty"`
	// <<< add by liangc : set chief start number <<<
	// add by liangc : new rules for chief.tx of gaspool
	NR001Block *big.Int `json:"nr001Block,omitempty"` // NR001 HF block
	// add by liangc : new rules for block period
	NR002Block *big.Int `json:"nr002Block,omitempty"` // NR002 HF block
	// add by liangc : 18-09-13 : incompatible HomesteadSigner begin at this number
	NR003Block *big.Int `json:"nr003Block,omitempty"`

	MeshboxBlock   *big.Int       `json:"chief006Block,omitempty"`
	MeshboxAddress common.Address `json:"chief006Address,omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// TribeConfig is the consensus engine configs.
type TribeConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
}

// String implements the stringer interface, returning the consensus engine details.
func (c *TribeConfig) String() string {
	return "tribe"
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	case c.Tribe != nil:
		engine = c.Tribe
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v Homestead: %v DAO: %v DAOSupport: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Engine: %v}",
		c.ChainId,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		engine,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsDAO returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isForked(c.DAOForkBlock, num)
}

func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	// add by liangc : set default byzantium
	return isForked(c.ByzantiumBlock, num)
}

// GasTable returns the gas table corresponding to the current phase (homestead or homestead reprice).
//
// The returned GasTable's fields shouldn't, under any circumstances, be changed.
func (c *ChainConfig) GasTable(num *big.Int) GasTable {
	if num == nil {
		return GasTableHomestead
	}
	switch {
	case c.IsEIP158(num):
		return GasTableEIP158
	case c.IsEIP150(num):
		return GasTableEIP150
	default:
		return GasTableHomestead
	}
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
		return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, head) {
		return newCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if c.IsDAOFork(head) && c.DAOForkSupport != newcfg.DAOForkSupport {
		return newCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
		return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
		return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(head) && !configNumEqual(c.ChainId, newcfg.ChainId) {
		return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
		return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntatic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainId                                   *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158 bool
	IsByzantium                               bool
}

func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainId := c.ChainId
	if chainId == nil {
		chainId = new(big.Int)
	}
	return Rules{ChainId: new(big.Int).Set(chainId), IsHomestead: c.IsHomestead(num), IsEIP150: c.IsEIP150(num), IsEIP155: c.IsEIP155(num), IsEIP158: c.IsEIP158(num), IsByzantium: c.IsByzantium(num)}
}
