// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Contract

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

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bcid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"enKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"cveHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"enHash\",\"type\":\"bytes\"}],\"name\":\"NewDataSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressToBcid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"getBCID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"getEnCVE\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"getEnHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"getEnKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"imageToOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bcid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"enKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"enHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"enCVE\",\"type\":\"bytes\"}],\"name\":\"setDockerImageData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// AddressToBcid is a free data retrieval call binding the contract method 0xb5f30f0a.
//
// Solidity: function addressToBcid(address , uint256 ) view returns(uint256)
func (_Contract *ContractCaller) AddressToBcid(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "addressToBcid", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToBcid is a free data retrieval call binding the contract method 0xb5f30f0a.
//
// Solidity: function addressToBcid(address , uint256 ) view returns(uint256)
func (_Contract *ContractSession) AddressToBcid(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.AddressToBcid(&_Contract.CallOpts, arg0, arg1)
}

// AddressToBcid is a free data retrieval call binding the contract method 0xb5f30f0a.
//
// Solidity: function addressToBcid(address , uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) AddressToBcid(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.AddressToBcid(&_Contract.CallOpts, arg0, arg1)
}

// GetBCID is a free data retrieval call binding the contract method 0x14629afe.
//
// Solidity: function getBCID(uint256 num) view returns(uint256)
func (_Contract *ContractCaller) GetBCID(opts *bind.CallOpts, num *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBCID", num)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBCID is a free data retrieval call binding the contract method 0x14629afe.
//
// Solidity: function getBCID(uint256 num) view returns(uint256)
func (_Contract *ContractSession) GetBCID(num *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetBCID(&_Contract.CallOpts, num)
}

// GetBCID is a free data retrieval call binding the contract method 0x14629afe.
//
// Solidity: function getBCID(uint256 num) view returns(uint256)
func (_Contract *ContractCallerSession) GetBCID(num *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetBCID(&_Contract.CallOpts, num)
}

// GetEnCVE is a free data retrieval call binding the contract method 0x3391257f.
//
// Solidity: function getEnCVE(uint256 num) view returns(bytes)
func (_Contract *ContractCaller) GetEnCVE(opts *bind.CallOpts, num *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEnCVE", num)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEnCVE is a free data retrieval call binding the contract method 0x3391257f.
//
// Solidity: function getEnCVE(uint256 num) view returns(bytes)
func (_Contract *ContractSession) GetEnCVE(num *big.Int) ([]byte, error) {
	return _Contract.Contract.GetEnCVE(&_Contract.CallOpts, num)
}

// GetEnCVE is a free data retrieval call binding the contract method 0x3391257f.
//
// Solidity: function getEnCVE(uint256 num) view returns(bytes)
func (_Contract *ContractCallerSession) GetEnCVE(num *big.Int) ([]byte, error) {
	return _Contract.Contract.GetEnCVE(&_Contract.CallOpts, num)
}

// GetEnHash is a free data retrieval call binding the contract method 0xd9a41e88.
//
// Solidity: function getEnHash(uint256 num) view returns(bytes)
func (_Contract *ContractCaller) GetEnHash(opts *bind.CallOpts, num *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEnHash", num)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEnHash is a free data retrieval call binding the contract method 0xd9a41e88.
//
// Solidity: function getEnHash(uint256 num) view returns(bytes)
func (_Contract *ContractSession) GetEnHash(num *big.Int) ([]byte, error) {
	return _Contract.Contract.GetEnHash(&_Contract.CallOpts, num)
}

// GetEnHash is a free data retrieval call binding the contract method 0xd9a41e88.
//
// Solidity: function getEnHash(uint256 num) view returns(bytes)
func (_Contract *ContractCallerSession) GetEnHash(num *big.Int) ([]byte, error) {
	return _Contract.Contract.GetEnHash(&_Contract.CallOpts, num)
}

// GetEnKey is a free data retrieval call binding the contract method 0x7809da4f.
//
// Solidity: function getEnKey(uint256 num) view returns(bytes)
func (_Contract *ContractCaller) GetEnKey(opts *bind.CallOpts, num *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEnKey", num)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEnKey is a free data retrieval call binding the contract method 0x7809da4f.
//
// Solidity: function getEnKey(uint256 num) view returns(bytes)
func (_Contract *ContractSession) GetEnKey(num *big.Int) ([]byte, error) {
	return _Contract.Contract.GetEnKey(&_Contract.CallOpts, num)
}

// GetEnKey is a free data retrieval call binding the contract method 0x7809da4f.
//
// Solidity: function getEnKey(uint256 num) view returns(bytes)
func (_Contract *ContractCallerSession) GetEnKey(num *big.Int) ([]byte, error) {
	return _Contract.Contract.GetEnKey(&_Contract.CallOpts, num)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256[])
func (_Contract *ContractCaller) GetNum(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getNum")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256[])
func (_Contract *ContractSession) GetNum() ([]*big.Int, error) {
	return _Contract.Contract.GetNum(&_Contract.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256[])
func (_Contract *ContractCallerSession) GetNum() ([]*big.Int, error) {
	return _Contract.Contract.GetNum(&_Contract.CallOpts)
}

// ImageToOwner is a free data retrieval call binding the contract method 0x42947657.
//
// Solidity: function imageToOwner(uint256 ) view returns(address)
func (_Contract *ContractCaller) ImageToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "imageToOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ImageToOwner is a free data retrieval call binding the contract method 0x42947657.
//
// Solidity: function imageToOwner(uint256 ) view returns(address)
func (_Contract *ContractSession) ImageToOwner(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.ImageToOwner(&_Contract.CallOpts, arg0)
}

// ImageToOwner is a free data retrieval call binding the contract method 0x42947657.
//
// Solidity: function imageToOwner(uint256 ) view returns(address)
func (_Contract *ContractCallerSession) ImageToOwner(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.ImageToOwner(&_Contract.CallOpts, arg0)
}

// SetDockerImageData is a paid mutator transaction binding the contract method 0x322aab34.
//
// Solidity: function setDockerImageData(uint256 bcid, bytes enKey, bytes enHash, bytes enCVE) returns()
func (_Contract *ContractTransactor) SetDockerImageData(opts *bind.TransactOpts, bcid *big.Int, enKey []byte, enHash []byte, enCVE []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setDockerImageData", bcid, enKey, enHash, enCVE)
}

// SetDockerImageData is a paid mutator transaction binding the contract method 0x322aab34.
//
// Solidity: function setDockerImageData(uint256 bcid, bytes enKey, bytes enHash, bytes enCVE) returns()
func (_Contract *ContractSession) SetDockerImageData(bcid *big.Int, enKey []byte, enHash []byte, enCVE []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetDockerImageData(&_Contract.TransactOpts, bcid, enKey, enHash, enCVE)
}

// SetDockerImageData is a paid mutator transaction binding the contract method 0x322aab34.
//
// Solidity: function setDockerImageData(uint256 bcid, bytes enKey, bytes enHash, bytes enCVE) returns()
func (_Contract *ContractTransactorSession) SetDockerImageData(bcid *big.Int, enKey []byte, enHash []byte, enCVE []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetDockerImageData(&_Contract.TransactOpts, bcid, enKey, enHash, enCVE)
}

// ContractNewDataSetIterator is returned from FilterNewDataSet and is used to iterate over the raw logs and unpacked data for NewDataSet events raised by the Contract contract.
type ContractNewDataSetIterator struct {
	Event *ContractNewDataSet // Event containing the contract specifics and raw log

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
func (it *ContractNewDataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewDataSet)
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
		it.Event = new(ContractNewDataSet)
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
func (it *ContractNewDataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewDataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewDataSet represents a NewDataSet event raised by the Contract contract.
type ContractNewDataSet struct {
	Num     *big.Int
	Bcid    *big.Int
	EnKey   []byte
	CveHash []byte
	EnHash  []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewDataSet is a free log retrieval operation binding the contract event 0xbd0ecdfa452b489f4cdeea99c3328dd6d4c8c15bd79b1da7be9bd9ca295eaba6.
//
// Solidity: event NewDataSet(uint256 num, uint256 bcid, bytes enKey, bytes cveHash, bytes enHash)
func (_Contract *ContractFilterer) FilterNewDataSet(opts *bind.FilterOpts) (*ContractNewDataSetIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewDataSet")
	if err != nil {
		return nil, err
	}
	return &ContractNewDataSetIterator{contract: _Contract.contract, event: "NewDataSet", logs: logs, sub: sub}, nil
}

// WatchNewDataSet is a free log subscription operation binding the contract event 0xbd0ecdfa452b489f4cdeea99c3328dd6d4c8c15bd79b1da7be9bd9ca295eaba6.
//
// Solidity: event NewDataSet(uint256 num, uint256 bcid, bytes enKey, bytes cveHash, bytes enHash)
func (_Contract *ContractFilterer) WatchNewDataSet(opts *bind.WatchOpts, sink chan<- *ContractNewDataSet) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewDataSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewDataSet)
				if err := _Contract.contract.UnpackLog(event, "NewDataSet", log); err != nil {
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

// ParseNewDataSet is a log parse operation binding the contract event 0xbd0ecdfa452b489f4cdeea99c3328dd6d4c8c15bd79b1da7be9bd9ca295eaba6.
//
// Solidity: event NewDataSet(uint256 num, uint256 bcid, bytes enKey, bytes cveHash, bytes enHash)
func (_Contract *ContractFilterer) ParseNewDataSet(log types.Log) (*ContractNewDataSet, error) {
	event := new(ContractNewDataSet)
	if err := _Contract.contract.UnpackLog(event, "NewDataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
