// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ArbitrumMetaData contains all meta data concerning the Arbitrum contract.
var ArbitrumMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040526040516200103238038062001032833981016040819052620000269162000487565b828162000036828260006200004d565b50620000449050826200008a565b505050620005ba565b6200005883620000e5565b600082511180620000665750805b1562000085576200008383836200012760201b6200028c1760201c565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f620000b562000156565b604080516001600160a01b03928316815291841660208301520160405180910390a1620000e2816200018f565b50565b620000f08162000244565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606200014f83836040518060600160405280602781526020016200100b60279139620002f8565b9392505050565b60006200018060008051602062000feb83398151915260001b620003de60201b6200022e1760201c565b546001600160a01b0316919050565b6001600160a01b038116620001fa5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b806200022360008051602062000feb83398151915260001b620003de60201b6200022e1760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b6200025a81620003e160201b620002b81760201c565b620002be5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401620001f1565b80620002237f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b620003de60201b6200022e1760201c565b60606001600160a01b0384163b620003625760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401620001f1565b600080856001600160a01b0316856040516200037f919062000567565b600060405180830381855af49150503d8060008114620003bc576040519150601f19603f3d011682016040523d82523d6000602084013e620003c1565b606091505b509092509050620003d4828286620003f0565b9695505050505050565b90565b6001600160a01b03163b151590565b60608315620004015750816200014f565b825115620004125782518084602001fd5b8160405162461bcd60e51b8152600401620001f1919062000585565b80516001600160a01b03811681146200044657600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b838110156200047e57818101518382015260200162000464565b50506000910152565b6000806000606084860312156200049d57600080fd5b620004a8846200042e565b9250620004b8602085016200042e565b60408501519092506001600160401b0380821115620004d657600080fd5b818601915086601f830112620004eb57600080fd5b8151818111156200050057620005006200044b565b604051601f8201601f19908116603f011681019083821181831017156200052b576200052b6200044b565b816040528281528960208487010111156200054557600080fd5b6200055883602083016020880162000461565b80955050505050509250925092565b600082516200057b81846020870162000461565b9190910192915050565b6020815260008251806020840152620005a681604085016020870162000461565b601f01601f19169190910160400192915050565b610a2180620005ca6000396000f3fe60806040526004361061005e5760003560e01c80635c60da1b116100435780635c60da1b146100a85780638f283970146100e6578063f851a440146101065761006d565b80633659cfe6146100755780634f1ef286146100955761006d565b3661006d5761006b61011b565b005b61006b61011b565b34801561008157600080fd5b5061006b610090366004610895565b610135565b61006b6100a33660046108b0565b61017f565b3480156100b457600080fd5b506100bd6101f3565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100f257600080fd5b5061006b610101366004610895565b610231565b34801561011257600080fd5b506100bd61025e565b6101236102d4565b61013361012e6103ab565b6103b5565b565b61013d6103d9565b73ffffffffffffffffffffffffffffffffffffffff1633036101775761017481604051806020016040528060008152506000610419565b50565b61017461011b565b6101876103d9565b73ffffffffffffffffffffffffffffffffffffffff1633036101eb576101e68383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525060019250610419915050565b505050565b6101e661011b565b60006101fd6103d9565b73ffffffffffffffffffffffffffffffffffffffff163303610226576102216103ab565b905090565b61022e61011b565b90565b6102396103d9565b73ffffffffffffffffffffffffffffffffffffffff1633036101775761017481610444565b60006102686103d9565b73ffffffffffffffffffffffffffffffffffffffff163303610226576102216103d9565b60606102b183836040518060600160405280602781526020016109c5602791396104a5565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6102dc6103d9565b73ffffffffffffffffffffffffffffffffffffffff163303610133576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f7879207461726760648201527f6574000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b60006102216105cd565b3660008037600080366000845af43d6000803e8080156103d4573d6000f35b3d6000fd5b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b5473ffffffffffffffffffffffffffffffffffffffff16919050565b610422836105f5565b60008251118061042f5750805b156101e65761043e838361028c565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61046d6103d9565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301520160405180910390a161017481610642565b606073ffffffffffffffffffffffffffffffffffffffff84163b61054b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60448201527f6e7472616374000000000000000000000000000000000000000000000000000060648201526084016103a2565b6000808573ffffffffffffffffffffffffffffffffffffffff16856040516105739190610957565b600060405180830381855af49150503d80600081146105ae576040519150601f19603f3d011682016040523d82523d6000602084013e6105b3565b606091505b50915091506105c382828661074e565b9695505050505050565b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6103fd565b6105fe816107a1565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b73ffffffffffffffffffffffffffffffffffffffff81166106e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016103a2565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b6060831561075d5750816102b1565b82511561076d5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a29190610973565b73ffffffffffffffffffffffffffffffffffffffff81163b610845576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e74726163740000000000000000000000000000000000000060648201526084016103a2565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc610708565b803573ffffffffffffffffffffffffffffffffffffffff8116811461089057600080fd5b919050565b6000602082840312156108a757600080fd5b6102b18261086c565b6000806000604084860312156108c557600080fd5b6108ce8461086c565b9250602084013567ffffffffffffffff808211156108eb57600080fd5b818601915086601f8301126108ff57600080fd5b81358181111561090e57600080fd5b87602082850101111561092057600080fd5b6020830194508093505050509250925092565b60005b8381101561094e578181015183820152602001610936565b50506000910152565b60008251610969818460208701610933565b9190910192915050565b6020815260008251806020840152610992816040850160208701610933565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212205f078eeb5690e33d91e7b90c18c8f4a8b449ac85285d1fee003c4e18e239c87764736f6c63430008100033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564000000000000000000000000c4ed0a9ea70d5bcc69f748547650d32cc219d882000000000000000000000000db216562328215e010f819b5abe947bad4ca961e00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000",
}

// ArbitrumABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbitrumMetaData.ABI instead.
var ArbitrumABI = ArbitrumMetaData.ABI

// ArbitrumBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ArbitrumMetaData.Bin instead.
var ArbitrumBin = ArbitrumMetaData.Bin

// DeployArbitrum deploys a new Ethereum contract, binding an instance of Arbitrum to it.
func DeployArbitrum(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, admin_ common.Address, _data []byte) (common.Address, *types.Transaction, *Arbitrum, error) {
	parsed, err := ArbitrumMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArbitrumBin), backend, _logic, admin_, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Arbitrum{ArbitrumCaller: ArbitrumCaller{contract: contract}, ArbitrumTransactor: ArbitrumTransactor{contract: contract}, ArbitrumFilterer: ArbitrumFilterer{contract: contract}}, nil
}

// Arbitrum is an auto generated Go binding around an Ethereum contract.
type Arbitrum struct {
	ArbitrumCaller     // Read-only binding to the contract
	ArbitrumTransactor // Write-only binding to the contract
	ArbitrumFilterer   // Log filterer for contract events
}

// ArbitrumCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbitrumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbitrumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbitrumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbitrumSession struct {
	Contract     *Arbitrum         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbitrumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbitrumCallerSession struct {
	Contract *ArbitrumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ArbitrumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbitrumTransactorSession struct {
	Contract     *ArbitrumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ArbitrumRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbitrumRaw struct {
	Contract *Arbitrum // Generic contract binding to access the raw methods on
}

// ArbitrumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbitrumCallerRaw struct {
	Contract *ArbitrumCaller // Generic read-only contract binding to access the raw methods on
}

// ArbitrumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbitrumTransactorRaw struct {
	Contract *ArbitrumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbitrum creates a new instance of Arbitrum, bound to a specific deployed contract.
func NewArbitrum(address common.Address, backend bind.ContractBackend) (*Arbitrum, error) {
	contract, err := bindArbitrum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arbitrum{ArbitrumCaller: ArbitrumCaller{contract: contract}, ArbitrumTransactor: ArbitrumTransactor{contract: contract}, ArbitrumFilterer: ArbitrumFilterer{contract: contract}}, nil
}

// NewArbitrumCaller creates a new read-only instance of Arbitrum, bound to a specific deployed contract.
func NewArbitrumCaller(address common.Address, caller bind.ContractCaller) (*ArbitrumCaller, error) {
	contract, err := bindArbitrum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCaller{contract: contract}, nil
}

// NewArbitrumTransactor creates a new write-only instance of Arbitrum, bound to a specific deployed contract.
func NewArbitrumTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrumTransactor, error) {
	contract, err := bindArbitrum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumTransactor{contract: contract}, nil
}

// NewArbitrumFilterer creates a new log filterer instance of Arbitrum, bound to a specific deployed contract.
func NewArbitrumFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrumFilterer, error) {
	contract, err := bindArbitrum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrumFilterer{contract: contract}, nil
}

// bindArbitrum binds a generic wrapper to an already deployed contract.
func bindArbitrum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbitrumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arbitrum *ArbitrumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arbitrum.Contract.ArbitrumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arbitrum *ArbitrumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrum.Contract.ArbitrumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arbitrum *ArbitrumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arbitrum.Contract.ArbitrumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arbitrum *ArbitrumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arbitrum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arbitrum *ArbitrumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arbitrum *ArbitrumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arbitrum.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Arbitrum *ArbitrumTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrum.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Arbitrum *ArbitrumSession) Admin() (*types.Transaction, error) {
	return _Arbitrum.Contract.Admin(&_Arbitrum.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Arbitrum *ArbitrumTransactorSession) Admin() (*types.Transaction, error) {
	return _Arbitrum.Contract.Admin(&_Arbitrum.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Arbitrum *ArbitrumTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Arbitrum.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Arbitrum *ArbitrumSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Arbitrum.Contract.ChangeAdmin(&_Arbitrum.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Arbitrum *ArbitrumTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Arbitrum.Contract.ChangeAdmin(&_Arbitrum.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Arbitrum *ArbitrumTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrum.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Arbitrum *ArbitrumSession) Implementation() (*types.Transaction, error) {
	return _Arbitrum.Contract.Implementation(&_Arbitrum.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Arbitrum *ArbitrumTransactorSession) Implementation() (*types.Transaction, error) {
	return _Arbitrum.Contract.Implementation(&_Arbitrum.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Arbitrum *ArbitrumTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Arbitrum.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Arbitrum *ArbitrumSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Arbitrum.Contract.UpgradeTo(&_Arbitrum.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Arbitrum *ArbitrumTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Arbitrum.Contract.UpgradeTo(&_Arbitrum.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Arbitrum *ArbitrumTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Arbitrum.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Arbitrum *ArbitrumSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Arbitrum.Contract.UpgradeToAndCall(&_Arbitrum.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Arbitrum *ArbitrumTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Arbitrum.Contract.UpgradeToAndCall(&_Arbitrum.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Arbitrum *ArbitrumTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Arbitrum.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Arbitrum *ArbitrumSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Arbitrum.Contract.Fallback(&_Arbitrum.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Arbitrum *ArbitrumTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Arbitrum.Contract.Fallback(&_Arbitrum.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Arbitrum *ArbitrumTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitrum.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Arbitrum *ArbitrumSession) Receive() (*types.Transaction, error) {
	return _Arbitrum.Contract.Receive(&_Arbitrum.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Arbitrum *ArbitrumTransactorSession) Receive() (*types.Transaction, error) {
	return _Arbitrum.Contract.Receive(&_Arbitrum.TransactOpts)
}

// ArbitrumAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Arbitrum contract.
type ArbitrumAdminChangedIterator struct {
	Event *ArbitrumAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbitrumAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbitrumAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbitrumAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbitrumAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbitrumAdminChanged represents a AdminChanged event raised by the Arbitrum contract.
type ArbitrumAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Arbitrum *ArbitrumFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ArbitrumAdminChangedIterator, error) {

	logs, sub, err := _Arbitrum.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ArbitrumAdminChangedIterator{contract: _Arbitrum.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Arbitrum *ArbitrumFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ArbitrumAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Arbitrum.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbitrumAdminChanged)
				if err := _Arbitrum.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Arbitrum *ArbitrumFilterer) ParseAdminChanged(log types.Log) (*ArbitrumAdminChanged, error) {
	event := new(ArbitrumAdminChanged)
	if err := _Arbitrum.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbitrumBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Arbitrum contract.
type ArbitrumBeaconUpgradedIterator struct {
	Event *ArbitrumBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbitrumBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbitrumBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbitrumBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbitrumBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbitrumBeaconUpgraded represents a BeaconUpgraded event raised by the Arbitrum contract.
type ArbitrumBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Arbitrum *ArbitrumFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ArbitrumBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Arbitrum.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumBeaconUpgradedIterator{contract: _Arbitrum.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Arbitrum *ArbitrumFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ArbitrumBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Arbitrum.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbitrumBeaconUpgraded)
				if err := _Arbitrum.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Arbitrum *ArbitrumFilterer) ParseBeaconUpgraded(log types.Log) (*ArbitrumBeaconUpgraded, error) {
	event := new(ArbitrumBeaconUpgraded)
	if err := _Arbitrum.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbitrumUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Arbitrum contract.
type ArbitrumUpgradedIterator struct {
	Event *ArbitrumUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbitrumUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbitrumUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbitrumUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbitrumUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbitrumUpgraded represents a Upgraded event raised by the Arbitrum contract.
type ArbitrumUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Arbitrum *ArbitrumFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ArbitrumUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Arbitrum.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumUpgradedIterator{contract: _Arbitrum.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Arbitrum *ArbitrumFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ArbitrumUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Arbitrum.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbitrumUpgraded)
				if err := _Arbitrum.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Arbitrum *ArbitrumFilterer) ParseUpgraded(log types.Log) (*ArbitrumUpgraded, error) {
	event := new(ArbitrumUpgraded)
	if err := _Arbitrum.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
