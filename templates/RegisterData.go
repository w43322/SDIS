// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RegisterData

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RegisterDataABI is the input ABI used to generate the binding from.
const RegisterDataABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bcid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"enKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"enHash\",\"type\":\"string\"}],\"name\":\"NewDataSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressToBcid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBCID\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bcid\",\"type\":\"uint256\"}],\"name\":\"getEnHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bcid\",\"type\":\"uint256\"}],\"name\":\"getEnKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"imageToOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"enKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"enHash\",\"type\":\"string\"}],\"name\":\"setDockerImageData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RegisterDataBin is the compiled bytecode used for deploying new contracts.
var RegisterDataBin = "0x608060405234801561001057600080fd5b50610ba1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806305619eb6146100675780630c3cf0781461008357806342947657146100a15780637809da4f146100d1578063b5f30f0a14610101578063d9a41e8814610131575b600080fd5b610081600480360381019061007c9190610712565b610161565b005b61008b6102ee565b604051610098919061089e565b60405180910390f35b6100bb60048036038101906100b6919061077e565b610383565b6040516100c89190610883565b60405180910390f35b6100eb60048036038101906100e6919061077e565b6103b6565b6040516100f891906108c0565b60405180910390f35b61011b600480360381019061011691906106d6565b610493565b60405161012891906108e2565b60405180910390f35b61014b6004803603810190610146919061077e565b6104c4565b60405161015891906108c0565b60405180910390f35b600080805490509050600060405180606001604052808381526020018581526020018481525090806001815401808255809150506001900390600052602060002090600302016000909190919091506000820151816000015560208201518160010190805190602001906101d69291906105a1565b5060408201518160020190805190602001906101f39291906105a1565b505050336001600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f602e1c3daaa19bfeac2d00dbc7048627e616488cdbc5e5747a684f40cc707d3581848460405161027b939291906108fd565b60405180910390a1600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819080600181540180825580915050600190039060005260206000200160009091909190915055505050565b6060600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561037957602002820191906000526020600020905b815481526020019060010190808311610365575b5050505050905090565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060600082815481106103f2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060030201600101805461040e90610a6b565b80601f016020809104026020016040519081016040528092919081815260200182805461043a90610a6b565b80156104875780601f1061045c57610100808354040283529160200191610487565b820191906000526020600020905b81548152906001019060200180831161046a57829003601f168201915b50505050509050919050565b600260205281600052604060002081815481106104af57600080fd5b90600052602060002001600091509150505481565b606060008281548110610500577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060030201600201805461051c90610a6b565b80601f016020809104026020016040519081016040528092919081815260200182805461054890610a6b565b80156105955780601f1061056a57610100808354040283529160200191610595565b820191906000526020600020905b81548152906001019060200180831161057857829003601f168201915b50505050509050919050565b8280546105ad90610a6b565b90600052602060002090601f0160209004810192826105cf5760008555610616565b82601f106105e857805160ff1916838001178555610616565b82800160010185558215610616579182015b828111156106155782518255916020019190600101906105fa565b5b5090506106239190610627565b5090565b5b80821115610640576000816000905550600101610628565b5090565b600061065761065284610967565b610942565b90508281526020810184848401111561066f57600080fd5b61067a848285610a29565b509392505050565b60008135905061069181610b3d565b92915050565b600082601f8301126106a857600080fd5b81356106b8848260208601610644565b91505092915050565b6000813590506106d081610b54565b92915050565b600080604083850312156106e957600080fd5b60006106f785828601610682565b9250506020610708858286016106c1565b9150509250929050565b6000806040838503121561072557600080fd5b600083013567ffffffffffffffff81111561073f57600080fd5b61074b85828601610697565b925050602083013567ffffffffffffffff81111561076857600080fd5b61077485828601610697565b9150509250929050565b60006020828403121561079057600080fd5b600061079e848285016106c1565b91505092915050565b60006107b38383610865565b60208301905092915050565b6107c8816109ed565b82525050565b60006107d9826109a8565b6107e381856109cb565b93506107ee83610998565b8060005b8381101561081f57815161080688826107a7565b9750610811836109be565b9250506001810190506107f2565b5085935050505092915050565b6000610837826109b3565b61084181856109dc565b9350610851818560208601610a38565b61085a81610b2c565b840191505092915050565b61086e81610a1f565b82525050565b61087d81610a1f565b82525050565b600060208201905061089860008301846107bf565b92915050565b600060208201905081810360008301526108b881846107ce565b905092915050565b600060208201905081810360008301526108da818461082c565b905092915050565b60006020820190506108f76000830184610874565b92915050565b60006060820190506109126000830186610874565b8181036020830152610924818561082c565b90508181036040830152610938818461082c565b9050949350505050565b600061094c61095d565b90506109588282610a9d565b919050565b6000604051905090565b600067ffffffffffffffff82111561098257610981610afd565b5b61098b82610b2c565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60006109f8826109ff565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610a56578082015181840152602081019050610a3b565b83811115610a65576000848401525b50505050565b60006002820490506001821680610a8357607f821691505b60208210811415610a9757610a96610ace565b5b50919050565b610aa682610b2c565b810181811067ffffffffffffffff82111715610ac557610ac4610afd565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b610b46816109ed565b8114610b5157600080fd5b50565b610b5d81610a1f565b8114610b6857600080fd5b5056fea264697066735822122061b65846af0781fa95dacb4fd86321507b91fb096b4bc117d68d6f7f1f45e71a64736f6c63430008020033"

// DeployRegisterData deploys a new Ethereum contract, binding an instance of RegisterData to it.
func DeployRegisterData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RegisterData, error) {
	parsed, err := abi.JSON(strings.NewReader(RegisterDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegisterDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RegisterData{RegisterDataCaller: RegisterDataCaller{contract: contract}, RegisterDataTransactor: RegisterDataTransactor{contract: contract}, RegisterDataFilterer: RegisterDataFilterer{contract: contract}}, nil
}

// RegisterData is an auto generated Go binding around an Ethereum contract.
type re struct {
	RegisterDataCaller     // Read-only binding to the contract
	RegisterDataTransactor // Write-only binding to the contract
	RegisterDataFilterer   // Log filterer for contract events
}

// RegisterDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegisterDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegisterDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegisterDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegisterDataSession struct {
	Contract     *RegisterData     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegisterDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegisterDataCallerSession struct {
	Contract *RegisterDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RegisterDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegisterDataTransactorSession struct {
	Contract     *RegisterDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RegisterDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegisterDataRaw struct {
	Contract *RegisterData // Generic contract binding to access the raw methods on
}

// RegisterDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegisterDataCallerRaw struct {
	Contract *RegisterDataCaller // Generic read-only contract binding to access the raw methods on
}

// RegisterDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegisterDataTransactorRaw struct {
	Contract *RegisterDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegisterData creates a new instance of RegisterData, bound to a specific deployed contract.
func NewRegisterData(address common.Address, backend bind.ContractBackend) (*RegisterData, error) {
	contract, err := bindRegisterData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegisterData{RegisterDataCaller: RegisterDataCaller{contract: contract}, RegisterDataTransactor: RegisterDataTransactor{contract: contract}, RegisterDataFilterer: RegisterDataFilterer{contract: contract}}, nil
}

// NewRegisterDataCaller creates a new read-only instance of RegisterData, bound to a specific deployed contract.
func NewRegisterDataCaller(address common.Address, caller bind.ContractCaller) (*RegisterDataCaller, error) {
	contract, err := bindRegisterData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegisterDataCaller{contract: contract}, nil
}

// NewRegisterDataTransactor creates a new write-only instance of RegisterData, bound to a specific deployed contract.
func NewRegisterDataTransactor(address common.Address, transactor bind.ContractTransactor) (*RegisterDataTransactor, error) {
	contract, err := bindRegisterData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegisterDataTransactor{contract: contract}, nil
}

// NewRegisterDataFilterer creates a new log filterer instance of RegisterData, bound to a specific deployed contract.
func NewRegisterDataFilterer(address common.Address, filterer bind.ContractFilterer) (*RegisterDataFilterer, error) {
	contract, err := bindRegisterData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegisterDataFilterer{contract: contract}, nil
}

// bindRegisterData binds a generic wrapper to an already deployed contract.
func bindRegisterData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegisterDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegisterData *RegisterDataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegisterData.Contract.RegisterDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegisterData *RegisterDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegisterData.Contract.RegisterDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegisterData *RegisterDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegisterData.Contract.RegisterDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegisterData *RegisterDataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegisterData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegisterData *RegisterDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegisterData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegisterData *RegisterDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegisterData.Contract.contract.Transact(opts, method, params...)
}

// AddressToBcid is a free data retrieval call binding the contract method 0xb5f30f0a.
//
// Solidity: function addressToBcid(address , uint256 ) view returns(uint256)
func (_RegisterData *RegisterDataCaller) AddressToBcid(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RegisterData.contract.Call(opts, &out, "addressToBcid", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToBcid is a free data retrieval call binding the contract method 0xb5f30f0a.
//
// Solidity: function addressToBcid(address , uint256 ) view returns(uint256)
func (_RegisterData *RegisterDataSession) AddressToBcid(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _RegisterData.Contract.AddressToBcid(&_RegisterData.CallOpts, arg0, arg1)
}

// AddressToBcid is a free data retrieval call binding the contract method 0xb5f30f0a.
//
// Solidity: function addressToBcid(address , uint256 ) view returns(uint256)
func (_RegisterData *RegisterDataCallerSession) AddressToBcid(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _RegisterData.Contract.AddressToBcid(&_RegisterData.CallOpts, arg0, arg1)
}

// GetBCID is a free data retrieval call binding the contract method 0x0c3cf078.
//
// Solidity: function getBCID() view returns(uint256[])
func (_RegisterData *RegisterDataCaller) GetBCID(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _RegisterData.contract.Call(opts, &out, "getBCID")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBCID is a free data retrieval call binding the contract method 0x0c3cf078.
//
// Solidity: function getBCID() view returns(uint256[])
func (_RegisterData *RegisterDataSession) GetBCID() ([]*big.Int, error) {
	return _RegisterData.Contract.GetBCID(&_RegisterData.CallOpts)
}

// GetBCID is a free data retrieval call binding the contract method 0x0c3cf078.
//
// Solidity: function getBCID() view returns(uint256[])
func (_RegisterData *RegisterDataCallerSession) GetBCID() ([]*big.Int, error) {
	return _RegisterData.Contract.GetBCID(&_RegisterData.CallOpts)
}

// GetEnHash is a free data retrieval call binding the contract method 0xd9a41e88.
//
// Solidity: function getEnHash(uint256 bcid) view returns(string)
func (_RegisterData *RegisterDataCaller) GetEnHash(opts *bind.CallOpts, bcid *big.Int) (string, error) {
	var out []interface{}
	err := _RegisterData.contract.Call(opts, &out, "getEnHash", bcid)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEnHash is a free data retrieval call binding the contract method 0xd9a41e88.
//
// Solidity: function getEnHash(uint256 bcid) view returns(string)
func (_RegisterData *RegisterDataSession) GetEnHash(bcid *big.Int) (string, error) {
	return _RegisterData.Contract.GetEnHash(&_RegisterData.CallOpts, bcid)
}

// GetEnHash is a free data retrieval call binding the contract method 0xd9a41e88.
//
// Solidity: function getEnHash(uint256 bcid) view returns(string)
func (_RegisterData *RegisterDataCallerSession) GetEnHash(bcid *big.Int) (string, error) {
	return _RegisterData.Contract.GetEnHash(&_RegisterData.CallOpts, bcid)
}

// GetEnKey is a free data retrieval call binding the contract method 0x7809da4f.
//
// Solidity: function getEnKey(uint256 bcid) view returns(string)
func (_RegisterData *RegisterDataCaller) GetEnKey(opts *bind.CallOpts, bcid *big.Int) (string, error) {
	var out []interface{}
	err := _RegisterData.contract.Call(opts, &out, "getEnKey", bcid)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEnKey is a free data retrieval call binding the contract method 0x7809da4f.
//
// Solidity: function getEnKey(uint256 bcid) view returns(string)
func (_RegisterData *RegisterDataSession) GetEnKey(bcid *big.Int) (string, error) {
	return _RegisterData.Contract.GetEnKey(&_RegisterData.CallOpts, bcid)
}

// GetEnKey is a free data retrieval call binding the contract method 0x7809da4f.
//
// Solidity: function getEnKey(uint256 bcid) view returns(string)
func (_RegisterData *RegisterDataCallerSession) GetEnKey(bcid *big.Int) (string, error) {
	return _RegisterData.Contract.GetEnKey(&_RegisterData.CallOpts, bcid)
}

// ImageToOwner is a free data retrieval call binding the contract method 0x42947657.
//
// Solidity: function imageToOwner(uint256 ) view returns(address)
func (_RegisterData *RegisterDataCaller) ImageToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RegisterData.contract.Call(opts, &out, "imageToOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ImageToOwner is a free data retrieval call binding the contract method 0x42947657.
//
// Solidity: function imageToOwner(uint256 ) view returns(address)
func (_RegisterData *RegisterDataSession) ImageToOwner(arg0 *big.Int) (common.Address, error) {
	return _RegisterData.Contract.ImageToOwner(&_RegisterData.CallOpts, arg0)
}

// ImageToOwner is a free data retrieval call binding the contract method 0x42947657.
//
// Solidity: function imageToOwner(uint256 ) view returns(address)
func (_RegisterData *RegisterDataCallerSession) ImageToOwner(arg0 *big.Int) (common.Address, error) {
	return _RegisterData.Contract.ImageToOwner(&_RegisterData.CallOpts, arg0)
}

// SetDockerImageData is a paid mutator transaction binding the contract method 0x05619eb6.
//
// Solidity: function setDockerImageData(string enKey, string enHash) returns()
func (_RegisterData *RegisterDataTransactor) SetDockerImageData(opts *bind.TransactOpts, enKey string, enHash string) (*types.Transaction, error) {
	return _RegisterData.contract.Transact(opts, "setDockerImageData", enKey, enHash)
}

// SetDockerImageData is a paid mutator transaction binding the contract method 0x05619eb6.
//
// Solidity: function setDockerImageData(string enKey, string enHash) returns()
func (_RegisterData *RegisterDataSession) SetDockerImageData(enKey string, enHash string) (*types.Transaction, error) {
	return _RegisterData.Contract.SetDockerImageData(&_RegisterData.TransactOpts, enKey, enHash)
}

// SetDockerImageData is a paid mutator transaction binding the contract method 0x05619eb6.
//
// Solidity: function setDockerImageData(string enKey, string enHash) returns()
func (_RegisterData *RegisterDataTransactorSession) SetDockerImageData(enKey string, enHash string) (*types.Transaction, error) {
	return _RegisterData.Contract.SetDockerImageData(&_RegisterData.TransactOpts, enKey, enHash)
}

// RegisterDataNewDataSetIterator is returned from FilterNewDataSet and is used to iterate over the raw logs and unpacked data for NewDataSet events raised by the RegisterData contract.
type RegisterDataNewDataSetIterator struct {
	Event *RegisterDataNewDataSet // Event containing the contract specifics and raw log

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
func (it *RegisterDataNewDataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegisterDataNewDataSet)
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
		it.Event = new(RegisterDataNewDataSet)
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
func (it *RegisterDataNewDataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegisterDataNewDataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegisterDataNewDataSet represents a NewDataSet event raised by the RegisterData contract.
type RegisterDataNewDataSet struct {
	Bcid   *big.Int
	EnKey  string
	EnHash string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewDataSet is a free log retrieval operation binding the contract event 0x602e1c3daaa19bfeac2d00dbc7048627e616488cdbc5e5747a684f40cc707d35.
//
// Solidity: event NewDataSet(uint256 bcid, string enKey, string enHash)
func (_RegisterData *RegisterDataFilterer) FilterNewDataSet(opts *bind.FilterOpts) (*RegisterDataNewDataSetIterator, error) {

	logs, sub, err := _RegisterData.contract.FilterLogs(opts, "NewDataSet")
	if err != nil {
		return nil, err
	}
	return &RegisterDataNewDataSetIterator{contract: _RegisterData.contract, event: "NewDataSet", logs: logs, sub: sub}, nil
}

// WatchNewDataSet is a free log subscription operation binding the contract event 0x602e1c3daaa19bfeac2d00dbc7048627e616488cdbc5e5747a684f40cc707d35.
//
// Solidity: event NewDataSet(uint256 bcid, string enKey, string enHash)
func (_RegisterData *RegisterDataFilterer) WatchNewDataSet(opts *bind.WatchOpts, sink chan<- *RegisterDataNewDataSet) (event.Subscription, error) {

	logs, sub, err := _RegisterData.contract.WatchLogs(opts, "NewDataSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegisterDataNewDataSet)
				if err := _RegisterData.contract.UnpackLog(event, "NewDataSet", log); err != nil {
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

// ParseNewDataSet is a log parse operation binding the contract event 0x602e1c3daaa19bfeac2d00dbc7048627e616488cdbc5e5747a684f40cc707d35.
//
// Solidity: event NewDataSet(uint256 bcid, string enKey, string enHash)
func (_RegisterData *RegisterDataFilterer) ParseNewDataSet(log types.Log) (*RegisterDataNewDataSet, error) {
	event := new(RegisterDataNewDataSet)
	if err := _RegisterData.contract.UnpackLog(event, "NewDataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
