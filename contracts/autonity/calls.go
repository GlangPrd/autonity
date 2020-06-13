package autonity

import (
	"github.com/clearmatics/autonity/common"
	"github.com/clearmatics/autonity/core/state"
	"github.com/clearmatics/autonity/core/types"
	"github.com/clearmatics/autonity/core/vm"
	"github.com/clearmatics/autonity/log"
	"github.com/clearmatics/autonity/params"
	"math"
	"math/big"
	"reflect"
	"sort"
)

/*
 * ContractState is a unified structure to represent the autonity contract state.
 * By using a unified structure, the new state meta introduced in the Autonity.sol
 * should be synced with this structure.
 */
type ContractState struct {
	Users           []common.Address `abi:"users"`
	Enodes          []string         `abi:"enodes"`
	Types           []*big.Int       `abi:"types"`
	Stakes          []*big.Int       `abi:"stakes"`
	CommissionRates []*big.Int       `abi:"commisionrates"`
	Operator        common.Address   `abi:"operator"`
	Deployer        common.Address   `abi:"deployer"`
	MinGasPrice     *big.Int         `abi:"mingasprice"`
	BondingPeriod   *big.Int         `abi:"bondingperiod"`
}

type raw []byte

// deployContract deploys the contract contained within the genesis field bytecode
func (ac *evmContract) DeployAutonityContract(chain consensus.ChainReader, header *types.Header, statedb *state.StateDB) error {
	// Convert the contract bytecode from hex into bytes
	contractBytecode := common.Hex2Bytes(chain.Config().AutonityContractConfig.Bytecode)
	evm := ac.evmProvider.EVM(header, deployer, statedb)

	ln := len(chainConfig.AutonityContractConfig.GetValidatorUsers())
	validators := make(common.Addresses, 0, ln)
	enodes := make([]string, 0, ln)
	accTypes := make([]*big.Int, 0, ln)
	participantStake := make([]*big.Int, 0, ln)
	commissionRate := make([]*big.Int, 0, ln)

	// Default bond period is 100.
	defaultBondPeriod := big.NewInt(100)

	defaultCommitteeSize := big.NewInt(1000)
	defaultVersion := "v0.0.0"

	for _, v := range chainConfig.AutonityContractConfig.Users {
		validators = append(validators, *v.Address)
		enodes = append(enodes, v.Enode)
		accTypes = append(accTypes, big.NewInt(int64(v.Type.GetID())))
		participantStake = append(participantStake, big.NewInt(int64(v.Stake)))

		// TODO: default commission rate is 0, should use a config file...
		commissionRate = append(commissionRate, big.NewInt(0))
	}

	constructorParams, err := ac.contractABI.Pack("",
		validators,
		enodes,
		accTypes,
		participantStake,
		commissionRate,
		chainConfig.AutonityContractConfig.Operator,
		new(big.Int).SetUint64(chainConfig.AutonityContractConfig.MinGasPrice),
		defaultBondPeriod,
		defaultCommitteeSize,
		defaultVersion)
	if err != nil {
		log.Error("contractABI.Pack returns err", "err", err)
		return err
	}

	data := append(contractBytecode, constructorParams...)
	gas := uint64(0xFFFFFFFF)
	value := new(big.Int).SetUint64(0x00)

	// Deploy the Autonity contract
	_, _, _, vmerr := evm.Create(vm.AccountRef(deployer), data, gas, value)
	if vmerr != nil {
		log.Error("DeployAutonityContract evm.Create", "err", vmerr)
		return vmerr
	}
	log.Info("Deployed Autonity Contract", "Address", ContractAddress.String())

	return nil
}

func (ac *evmContract) updateAutonityContract(header *types.Header, statedb *state.StateDB, bytecode string, state []byte) error {
	evm := ac.evmProvider.EVM(header, deployer, statedb)
	contractBytecode := common.Hex2Bytes(bytecode)
	data := append(contractBytecode, state...)
	gas := uint64(0xFFFFFFFF)
	value := new(big.Int).SetUint64(0x00)
	_, _, _, vmerr := evm.CreateWithAddress(vm.AccountRef(deployer), data, gas, value, ContractAddress)
	if vmerr != nil {
		log.Error("updateAutonityContract evm.Create", "err", vmerr)
		return vmerr
	}
	return nil
}

func (ac *evmContract) AutonityContractCall(statedb *state.StateDB, header *types.Header, function string, result interface{}, args ...interface{}) error {
	gas := uint64(math.MaxUint64)
	evm := ac.evmProvider.EVM(header, deployer, statedb)

	input, err := ac.contractABI.Pack(function, args...)
	if err != nil {
		return err
	}

	ret, _, vmerr := evm.Call(vm.AccountRef(deployer), ContractAddress, input, gas, new(big.Int))
	if vmerr != nil {
		log.Error("Error Autonity Contract", "function", function)
		return vmerr
	}
	// if result's type is "raw" then bypass unpacking
	if reflect.TypeOf(result) == reflect.TypeOf(&raw{}) {
		rawPtr := result.(*raw)
		*rawPtr = raw(ret)
		return nil
	}

	if err := ac.contractABI.Unpack(result, function, ret); err != nil {
		log.Error("Could not unpack returned value", "function", function)
		return err
	}

	return nil
}

func (ac *evmContract) callGetWhitelist(state *state.StateDB, header *types.Header) (*types.Nodes, error) {
	var returnedEnodes []string
	err := ac.AutonityContractCall(state, header, "getWhitelist", &returnedEnodes)
	if err != nil {
		return nil, err
	}
	return types.NewNodes(returnedEnodes), nil
}

func (ac *evmContract) callGetMinimumGasPrice(state *state.StateDB, header *types.Header) (uint64, error) {
	minGasPrice := new(big.Int)
	err := ac.AutonityContractCall(state, header, "getMinimumGasPrice", &minGasPrice)
	if err != nil {
		return 0, err
	}
	return minGasPrice.Uint64(), nil
}

func (ac *evmContract) callFinalize(state *state.StateDB, header *types.Header, blockGas *big.Int) (bool, types.Committee, error) {

	var updateReady bool
	var committee types.Committee

	err := ac.AutonityContractCall(state, header, "finalize", &[]interface{}{&updateReady, &committee}, blockGas)
	if err != nil {
		return false, nil, err
	}
	sort.Sort(committee)
	// submit the final reward distribution metrics.
	//ac.metrics.SubmitRewardDistributionMetrics(&v, header.Number.Uint64())
	return updateReady, committee, nil
}

func (ac *evmContract) callRetrieveState(statedb *state.StateDB, header *types.Header) ([]byte, error) {
	var state raw

	err := ac.AutonityContractCall(statedb, header, "retrieveState", &state)
	if err != nil {
		return nil, err
	}

	return state, nil
}

func (ac *evmContract) callRetrieveContract(state *state.StateDB, header *types.Header) (string, string, error) {
	var bytecode string
	var abi string
	err := ac.AutonityContractCall(state, header, "retrieveContract", &[]interface{}{&bytecode, &abi})
	if err != nil {
		return "", "", err
	}
	return bytecode, abi, nil
}

func (ac *evmContract) callSetMinimumGasPrice(state *state.StateDB, header *types.Header, price *big.Int) error {
	// Needs to be refactored somehow
	gas := uint64(0xFFFFFFFF)
	evm := ac.evmProvider.EVM(header, deployer, state)

	input, err := ac.contractABI.Pack("setMinimumGasPrice")
	if err != nil {
		return err
	}

	_, _, vmerr := evm.Call(vm.AccountRef(deployer), ContractAddress, input, gas, price)
	if vmerr != nil {
		log.Error("Error Autonity Contract getMinimumGasPrice()")
		return vmerr
	}
	return nil
}
