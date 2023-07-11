package autonity

import (
	"fmt"
	"github.com/autonity/autonity/common"
	"github.com/autonity/autonity/core/state"
	"github.com/autonity/autonity/core/types"
	"github.com/autonity/autonity/core/vm"
	"github.com/autonity/autonity/log"
	"github.com/autonity/autonity/params"
	"github.com/autonity/autonity/params/generated"
	"math"
	"math/big"
	"reflect"
)

type raw []byte

func DeployContracts(genesisConfig *params.ChainConfig, evm *vm.EVM) error {
	if err := DeployAutonityContract(genesisConfig.AutonityContractConfig, evm); err != nil {
		return fmt.Errorf("error %w when deploying the autonity contract", err)
	}
	if err := DeployAccountabilityContract(evm); err != nil {
		return fmt.Errorf("error %w when deploying the accountability contract", err)
	}
	if err := DeployOracleContract(genesisConfig, evm); err != nil {
		return fmt.Errorf("error %w when deploying the oracle contract", err)
	}
	return nil
}

func DeployAccountabilityContract(evm *vm.EVM) error {
	constructorParams, err := generated.AccountabilityAbi.Pack("", AutonityContractAddress)
	if err != nil {
		log.Error("Accountability contract err", "err", err)
		return err
	}

	data := append(generated.AccountabilityBytecode, constructorParams...)
	gas := uint64(0xFFFFFFFF)
	value := new(big.Int).SetUint64(0x00)

	// Deploy the Accountability contract
	_, _, _, vmerr := evm.Create(vm.AccountRef(DeployerAddress), data, gas, value)
	if vmerr != nil {
		log.Error("DeployAutonityContract evm create", "err", vmerr)
		return vmerr
	}
	log.Info("Deployed Accountability contract", "Address", AccountabilityContractAddress.String())

	return nil
}

func DeployAutonityContract(genesisConfig *params.AutonityContractGenesis, evm *vm.EVM) error {
	contractConfig := AutonityConfig{
		OperatorAccount:        genesisConfig.Operator,
		TreasuryAccount:        genesisConfig.Treasury,
		OracleContract:         OracleContractAddress,
		TreasuryFee:            new(big.Int).SetUint64(genesisConfig.TreasuryFee),
		MinBaseFee:             new(big.Int).SetUint64(genesisConfig.MinBaseFee),
		DelegationRate:         new(big.Int).SetUint64(genesisConfig.DelegationRate),
		EpochPeriod:            new(big.Int).SetUint64(genesisConfig.EpochPeriod),
		UnbondingPeriod:        new(big.Int).SetUint64(genesisConfig.UnbondingPeriod),
		CommitteeSize:          new(big.Int).SetUint64(genesisConfig.MaxCommitteeSize),
		ContractVersion:        big.NewInt(1),
		BlockPeriod:            new(big.Int).SetUint64(genesisConfig.BlockPeriod),
		AccountabilityContract: AccountabilityContractAddress,
	}

	validators := make([]params.Validator, 0, len(genesisConfig.Validators))
	for _, v := range genesisConfig.Validators {
		validators = append(validators, *v)
	}
	constructorParams, err := genesisConfig.ABI.Pack("", validators, contractConfig)
	if err != nil {
		log.Error("contractABI.Pack returns err", "err", err)
		return err
	}

	data := append(genesisConfig.Bytecode, constructorParams...)
	gas := uint64(0xFFFFFFFF)
	value := new(big.Int).SetUint64(0x00)

	// Deploy the Autonity contract
	_, _, _, vmerr := evm.Create(vm.AccountRef(DeployerAddress), data, gas, value)
	if vmerr != nil {
		log.Error("DeployAutonityContract evm.Create", "err", vmerr)
		return vmerr
	}
	log.Info("Deployed Autonity contract", "Address", AutonityContractAddress.String())

	return nil
}

func DeployOracleContract(genesisConfig *params.ChainConfig, evm *vm.EVM) error {
	if genesisConfig.OracleContractConfig == nil {
		log.Info("Using default genesis parameters for the Oracle Contract")
		genesisConfig.OracleContractConfig = params.DefaultGenesisOracleConfig()
	}
	if err := genesisConfig.OracleContractConfig.Prepare(); err != nil {
		log.Crit("Error with Oracle Contract configuration", "err", err)
	}

	voters := make([]common.Address, len(genesisConfig.AutonityContractConfig.Validators))
	for _, val := range genesisConfig.AutonityContractConfig.Validators {
		voters = append(voters, val.OracleAddress)
	}

	constructorParams, err := genesisConfig.OracleContractConfig.ABI.Pack("",
		voters,
		AutonityContractAddress,
		genesisConfig.AutonityContractConfig.Operator, // same operator as autonity
		genesisConfig.OracleContractConfig.Symbols,
		new(big.Int).SetUint64(genesisConfig.OracleContractConfig.VotePeriod))
	if err != nil {
		return err
	}

	data := append(genesisConfig.OracleContractConfig.Bytecode, constructorParams...)
	gas := uint64(0xFFFFFFFF)
	value := new(big.Int).SetUint64(0x00)

	// Deploy the Oracle contract
	if _, _, _, err = evm.Create(vm.AccountRef(DeployerAddress), data, gas, value); err != nil {
		log.Error("DeployOracleContract evm.Create", "err", err)
		return err
	}

	log.Info("Deployed Oracle Contract", "Address", OracleContractAddress)
	return nil
}

func (c *Contracts) replaceAutonityBytecode(header *types.Header, statedb *state.StateDB, bytecode []byte) error {
	evm := c.evmProvider(header, DeployerAddress, statedb)
	_, _, _, vmerr := evm.Replace(vm.AccountRef(DeployerAddress), bytecode, AutonityContractAddress)
	if vmerr != nil {
		log.Error("replaceAutonityBytecode evm.Create", "err", vmerr)
		return vmerr
	}
	return nil
}

// AutonityContractCall calls the specified function of the autonity contract
// with the given args, and returns the output unpacked into the result
// interface.
func (c *Contracts) AutonityContractCall(statedb *state.StateDB, header *types.Header, function string, result any, args ...any) error {
	packedArgs, err := c.contractABI.Pack(function, args...)
	if err != nil {
		return err
	}
	ret, err := c.CallContractFunc(statedb, header, packedArgs)
	if err != nil {
		return err
	}
	// if result's type is "raw" then bypass unpacking
	if reflect.TypeOf(result) == reflect.TypeOf(&raw{}) {
		rawPtr := result.(*raw)
		*rawPtr = ret
		return nil
	}
	if err := c.contractABI.UnpackIntoInterface(result, function, ret); err != nil {
		log.Error("Could not unpack returned value", "function", function)
		return err
	}

	return nil
}

// CallContractFunc creates an evm object, uses it to call the
// specified function of the autonity contract with packedArgs and returns the
// packed result. If there is an error making the evm call it will be returned.
// Callers should use the autonity contract ABI to pack and unpack the args and
// result.
func (c *Contracts) CallContractFunc(statedb *state.StateDB, header *types.Header, packedArgs []byte) ([]byte, error) {
	gas := uint64(math.MaxUint64)
	evm := c.evmProvider(header, DeployerAddress, statedb)
	packedResult, _, err := evm.Call(vm.AccountRef(DeployerAddress), AutonityContractAddress, packedArgs, gas, new(big.Int))
	return packedResult, err
}

func (c *Contracts) callGetCommitteeEnodes(state *state.StateDB, header *types.Header) (*types.Nodes, error) {
	var returnedEnodes []string
	err := c.AutonityContractCall(state, header, "getCommitteeEnodes", &returnedEnodes)
	if err != nil {
		return nil, err
	}
	return types.NewNodes(returnedEnodes), nil
}

func (c *Contracts) callGetMinimumBaseFee(state *state.StateDB, header *types.Header) (uint64, error) {
	minBaseFee := new(big.Int)
	err := c.AutonityContractCall(state, header, "getMinimumBaseFee", &minBaseFee)
	if err != nil {
		return 0, err
	}
	return minBaseFee.Uint64(), nil
}

/*
func (c *Contracts) callGetProposer(state *state.StateDB, header *types.Header, height uint64, round int64) common.Address {
	var proposer common.Address
	h := new(big.Int).SetUint64(height)
	r := new(big.Int).SetInt64(round)
	err := c.AutonityContractCall(state, header, "getProposer", &proposer, h, r)
	if err != nil {
		log.Error("get proposer failed from contract.", "error", err)
		return common.Address{}
	}
	return proposer
}
*/

func (c *Contracts) callFinalize(state *state.StateDB, header *types.Header) (bool, types.Committee, error) {
	var updateReady bool
	var committee types.Committee
	if err := c.AutonityContractCall(state, header, "finalize", &[]any{&updateReady, &committee}); err != nil {
		return false, nil, err
	}
	return updateReady, committee, nil
}

func (c *Contracts) callRetrieveContract(state *state.StateDB, header *types.Header) ([]byte, string, error) {
	var bytecode []byte
	var abi string
	if err := c.AutonityContractCall(state, header, "getNewContract", &[]any{&bytecode, &abi}); err != nil {
		return nil, "", err
	}
	return bytecode, abi, nil
}
