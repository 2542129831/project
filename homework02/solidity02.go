// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigen

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

// AbigenABI is the input ABI used to generate the binding from.
const AbigenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AddPerson\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Admin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"SeePersonip\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"SeeNumber\",\"type\":\"uint8\"}],\"name\":\"Empower\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PersonInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"addPerson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfPersons\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"}],\"name\":\"getPersonAge\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"SeeIp\",\"type\":\"address\"}],\"name\":\"getPersonInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isPersonExsist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ip\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"}],\"name\":\"setPersonAgeMem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Abigen is an auto generated Go binding around an Ethereum contract.
type Abigen struct {
	AbigenCaller     // Read-only binding to the contract
	AbigenTransactor // Write-only binding to the contract
	AbigenFilterer   // Log filterer for contract events
}

// AbigenCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbigenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbigenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbigenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbigenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbigenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbigenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbigenSession struct {
	Contract     *Abigen           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbigenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbigenCallerSession struct {
	Contract *AbigenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbigenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbigenTransactorSession struct {
	Contract     *AbigenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbigenRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbigenRaw struct {
	Contract *Abigen // Generic contract binding to access the raw methods on
}

// AbigenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbigenCallerRaw struct {
	Contract *AbigenCaller // Generic read-only contract binding to access the raw methods on
}

// AbigenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbigenTransactorRaw struct {
	Contract *AbigenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbigen creates a new instance of Abigen, bound to a specific deployed contract.
func NewAbigen(address common.Address, backend bind.ContractBackend) (*Abigen, error) {
	contract, err := bindAbigen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abigen{AbigenCaller: AbigenCaller{contract: contract}, AbigenTransactor: AbigenTransactor{contract: contract}, AbigenFilterer: AbigenFilterer{contract: contract}}, nil
}

// NewAbigenCaller creates a new read-only instance of Abigen, bound to a specific deployed contract.
func NewAbigenCaller(address common.Address, caller bind.ContractCaller) (*AbigenCaller, error) {
	contract, err := bindAbigen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbigenCaller{contract: contract}, nil
}

// NewAbigenTransactor creates a new write-only instance of Abigen, bound to a specific deployed contract.
func NewAbigenTransactor(address common.Address, transactor bind.ContractTransactor) (*AbigenTransactor, error) {
	contract, err := bindAbigen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbigenTransactor{contract: contract}, nil
}

// NewAbigenFilterer creates a new log filterer instance of Abigen, bound to a specific deployed contract.
func NewAbigenFilterer(address common.Address, filterer bind.ContractFilterer) (*AbigenFilterer, error) {
	contract, err := bindAbigen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbigenFilterer{contract: contract}, nil
}

// bindAbigen binds a generic wrapper to an already deployed contract.
func bindAbigen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbigenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abigen *AbigenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abigen.Contract.AbigenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abigen *AbigenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abigen.Contract.AbigenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abigen *AbigenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abigen.Contract.AbigenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abigen *AbigenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abigen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abigen *AbigenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abigen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abigen *AbigenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abigen.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0x33b005d0.
//
// Solidity: function Admin(address ) view returns(bool)
func (_Abigen *AbigenCaller) Admin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "Admin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0x33b005d0.
//
// Solidity: function Admin(address ) view returns(bool)
func (_Abigen *AbigenSession) Admin(arg0 common.Address) (bool, error) {
	return _Abigen.Contract.Admin(&_Abigen.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0x33b005d0.
//
// Solidity: function Admin(address ) view returns(bool)
func (_Abigen *AbigenCallerSession) Admin(arg0 common.Address) (bool, error) {
	return _Abigen.Contract.Admin(&_Abigen.CallOpts, arg0)
}

// PersonInfo is a free data retrieval call binding the contract method 0x44c7c373.
//
// Solidity: function PersonInfo(address ) view returns(address ip, uint8 id, uint8 age, string name, string info)
func (_Abigen *AbigenCaller) PersonInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Ip   common.Address
	Id   uint8
	Age  uint8
	Name string
	Info string
}, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "PersonInfo", arg0)

	outstruct := new(struct {
		Ip   common.Address
		Id   uint8
		Age  uint8
		Name string
		Info string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ip = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Id = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Age = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Name = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Info = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// PersonInfo is a free data retrieval call binding the contract method 0x44c7c373.
//
// Solidity: function PersonInfo(address ) view returns(address ip, uint8 id, uint8 age, string name, string info)
func (_Abigen *AbigenSession) PersonInfo(arg0 common.Address) (struct {
	Ip   common.Address
	Id   uint8
	Age  uint8
	Name string
	Info string
}, error) {
	return _Abigen.Contract.PersonInfo(&_Abigen.CallOpts, arg0)
}

// PersonInfo is a free data retrieval call binding the contract method 0x44c7c373.
//
// Solidity: function PersonInfo(address ) view returns(address ip, uint8 id, uint8 age, string name, string info)
func (_Abigen *AbigenCallerSession) PersonInfo(arg0 common.Address) (struct {
	Ip   common.Address
	Id   uint8
	Age  uint8
	Name string
	Info string
}, error) {
	return _Abigen.Contract.PersonInfo(&_Abigen.CallOpts, arg0)
}

// GetNumberOfPersons is a free data retrieval call binding the contract method 0xcaac5da7.
//
// Solidity: function getNumberOfPersons() view returns(uint256)
func (_Abigen *AbigenCaller) GetNumberOfPersons(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "getNumberOfPersons")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfPersons is a free data retrieval call binding the contract method 0xcaac5da7.
//
// Solidity: function getNumberOfPersons() view returns(uint256)
func (_Abigen *AbigenSession) GetNumberOfPersons() (*big.Int, error) {
	return _Abigen.Contract.GetNumberOfPersons(&_Abigen.CallOpts)
}

// GetNumberOfPersons is a free data retrieval call binding the contract method 0xcaac5da7.
//
// Solidity: function getNumberOfPersons() view returns(uint256)
func (_Abigen *AbigenCallerSession) GetNumberOfPersons() (*big.Int, error) {
	return _Abigen.Contract.GetNumberOfPersons(&_Abigen.CallOpts)
}

// GetPersonAge is a free data retrieval call binding the contract method 0x0cfce82e.
//
// Solidity: function getPersonAge(address ip) view returns(uint8)
func (_Abigen *AbigenCaller) GetPersonAge(opts *bind.CallOpts, ip common.Address) (uint8, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "getPersonAge", ip)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPersonAge is a free data retrieval call binding the contract method 0x0cfce82e.
//
// Solidity: function getPersonAge(address ip) view returns(uint8)
func (_Abigen *AbigenSession) GetPersonAge(ip common.Address) (uint8, error) {
	return _Abigen.Contract.GetPersonAge(&_Abigen.CallOpts, ip)
}

// GetPersonAge is a free data retrieval call binding the contract method 0x0cfce82e.
//
// Solidity: function getPersonAge(address ip) view returns(uint8)
func (_Abigen *AbigenCallerSession) GetPersonAge(ip common.Address) (uint8, error) {
	return _Abigen.Contract.GetPersonAge(&_Abigen.CallOpts, ip)
}

// IsPersonExsist is a free data retrieval call binding the contract method 0x974b6cfb.
//
// Solidity: function isPersonExsist(address ) view returns(bool)
func (_Abigen *AbigenCaller) IsPersonExsist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "isPersonExsist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPersonExsist is a free data retrieval call binding the contract method 0x974b6cfb.
//
// Solidity: function isPersonExsist(address ) view returns(bool)
func (_Abigen *AbigenSession) IsPersonExsist(arg0 common.Address) (bool, error) {
	return _Abigen.Contract.IsPersonExsist(&_Abigen.CallOpts, arg0)
}

// IsPersonExsist is a free data retrieval call binding the contract method 0x974b6cfb.
//
// Solidity: function isPersonExsist(address ) view returns(bool)
func (_Abigen *AbigenCallerSession) IsPersonExsist(arg0 common.Address) (bool, error) {
	return _Abigen.Contract.IsPersonExsist(&_Abigen.CallOpts, arg0)
}

// Empower is a paid mutator transaction binding the contract method 0xb1bd42e6.
//
// Solidity: function Empower(address SeePersonip, uint8 SeeNumber) returns(bool)
func (_Abigen *AbigenTransactor) Empower(opts *bind.TransactOpts, SeePersonip common.Address, SeeNumber uint8) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "Empower", SeePersonip, SeeNumber)
}

// Empower is a paid mutator transaction binding the contract method 0xb1bd42e6.
//
// Solidity: function Empower(address SeePersonip, uint8 SeeNumber) returns(bool)
func (_Abigen *AbigenSession) Empower(SeePersonip common.Address, SeeNumber uint8) (*types.Transaction, error) {
	return _Abigen.Contract.Empower(&_Abigen.TransactOpts, SeePersonip, SeeNumber)
}

// Empower is a paid mutator transaction binding the contract method 0xb1bd42e6.
//
// Solidity: function Empower(address SeePersonip, uint8 SeeNumber) returns(bool)
func (_Abigen *AbigenTransactorSession) Empower(SeePersonip common.Address, SeeNumber uint8) (*types.Transaction, error) {
	return _Abigen.Contract.Empower(&_Abigen.TransactOpts, SeePersonip, SeeNumber)
}

// AddPerson is a paid mutator transaction binding the contract method 0x1f600adb.
//
// Solidity: function addPerson(address ip, uint8 id, uint8 age, string name, string info) returns()
func (_Abigen *AbigenTransactor) AddPerson(opts *bind.TransactOpts, ip common.Address, id uint8, age uint8, name string, info string) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "addPerson", ip, id, age, name, info)
}

// AddPerson is a paid mutator transaction binding the contract method 0x1f600adb.
//
// Solidity: function addPerson(address ip, uint8 id, uint8 age, string name, string info) returns()
func (_Abigen *AbigenSession) AddPerson(ip common.Address, id uint8, age uint8, name string, info string) (*types.Transaction, error) {
	return _Abigen.Contract.AddPerson(&_Abigen.TransactOpts, ip, id, age, name, info)
}

// AddPerson is a paid mutator transaction binding the contract method 0x1f600adb.
//
// Solidity: function addPerson(address ip, uint8 id, uint8 age, string name, string info) returns()
func (_Abigen *AbigenTransactorSession) AddPerson(ip common.Address, id uint8, age uint8, name string, info string) (*types.Transaction, error) {
	return _Abigen.Contract.AddPerson(&_Abigen.TransactOpts, ip, id, age, name, info)
}

// GetPersonInfo is a paid mutator transaction binding the contract method 0xd26526b0.
//
// Solidity: function getPersonInfo(address SeeIp) returns(string)
func (_Abigen *AbigenTransactor) GetPersonInfo(opts *bind.TransactOpts, SeeIp common.Address) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "getPersonInfo", SeeIp)
}

// GetPersonInfo is a paid mutator transaction binding the contract method 0xd26526b0.
//
// Solidity: function getPersonInfo(address SeeIp) returns(string)
func (_Abigen *AbigenSession) GetPersonInfo(SeeIp common.Address) (*types.Transaction, error) {
	return _Abigen.Contract.GetPersonInfo(&_Abigen.TransactOpts, SeeIp)
}

// GetPersonInfo is a paid mutator transaction binding the contract method 0xd26526b0.
//
// Solidity: function getPersonInfo(address SeeIp) returns(string)
func (_Abigen *AbigenTransactorSession) GetPersonInfo(SeeIp common.Address) (*types.Transaction, error) {
	return _Abigen.Contract.GetPersonInfo(&_Abigen.TransactOpts, SeeIp)
}

// SetPersonAgeMem is a paid mutator transaction binding the contract method 0x25b9a2c8.
//
// Solidity: function setPersonAgeMem(address ip, uint8 age) returns()
func (_Abigen *AbigenTransactor) SetPersonAgeMem(opts *bind.TransactOpts, ip common.Address, age uint8) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "setPersonAgeMem", ip, age)
}

// SetPersonAgeMem is a paid mutator transaction binding the contract method 0x25b9a2c8.
//
// Solidity: function setPersonAgeMem(address ip, uint8 age) returns()
func (_Abigen *AbigenSession) SetPersonAgeMem(ip common.Address, age uint8) (*types.Transaction, error) {
	return _Abigen.Contract.SetPersonAgeMem(&_Abigen.TransactOpts, ip, age)
}

// SetPersonAgeMem is a paid mutator transaction binding the contract method 0x25b9a2c8.
//
// Solidity: function setPersonAgeMem(address ip, uint8 age) returns()
func (_Abigen *AbigenTransactorSession) SetPersonAgeMem(ip common.Address, age uint8) (*types.Transaction, error) {
	return _Abigen.Contract.SetPersonAgeMem(&_Abigen.TransactOpts, ip, age)
}

// AbigenAddPersonIterator is returned from FilterAddPerson and is used to iterate over the raw logs and unpacked data for AddPerson events raised by the Abigen contract.
type AbigenAddPersonIterator struct {
	Event *AbigenAddPerson // Event containing the contract specifics and raw log

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
func (it *AbigenAddPersonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbigenAddPerson)
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
		it.Event = new(AbigenAddPerson)
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
func (it *AbigenAddPersonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbigenAddPersonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbigenAddPerson represents a AddPerson event raised by the Abigen contract.
type AbigenAddPerson struct {
	Id        uint8
	Age       uint8
	Name      string
	Info      string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddPerson is a free log retrieval operation binding the contract event 0x92d2361235a243e50564c2a187ffd783e61ee1bd4e074795d323cbf6cb2a2fa9.
//
// Solidity: event AddPerson(uint8 id, uint8 age, string name, string info, uint256 timestamp)
func (_Abigen *AbigenFilterer) FilterAddPerson(opts *bind.FilterOpts) (*AbigenAddPersonIterator, error) {

	logs, sub, err := _Abigen.contract.FilterLogs(opts, "AddPerson")
	if err != nil {
		return nil, err
	}
	return &AbigenAddPersonIterator{contract: _Abigen.contract, event: "AddPerson", logs: logs, sub: sub}, nil
}

// WatchAddPerson is a free log subscription operation binding the contract event 0x92d2361235a243e50564c2a187ffd783e61ee1bd4e074795d323cbf6cb2a2fa9.
//
// Solidity: event AddPerson(uint8 id, uint8 age, string name, string info, uint256 timestamp)
func (_Abigen *AbigenFilterer) WatchAddPerson(opts *bind.WatchOpts, sink chan<- *AbigenAddPerson) (event.Subscription, error) {

	logs, sub, err := _Abigen.contract.WatchLogs(opts, "AddPerson")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbigenAddPerson)
				if err := _Abigen.contract.UnpackLog(event, "AddPerson", log); err != nil {
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

// ParseAddPerson is a log parse operation binding the contract event 0x92d2361235a243e50564c2a187ffd783e61ee1bd4e074795d323cbf6cb2a2fa9.
//
// Solidity: event AddPerson(uint8 id, uint8 age, string name, string info, uint256 timestamp)
func (_Abigen *AbigenFilterer) ParseAddPerson(log types.Log) (*AbigenAddPerson, error) {
	event := new(AbigenAddPerson)
	if err := _Abigen.contract.UnpackLog(event, "AddPerson", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
