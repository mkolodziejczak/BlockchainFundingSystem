// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// ProjectMetaData contains all meta data concerning the Project contract.
var ProjectMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"projectCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"milestone1Deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"milestone2Deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"milestone3Deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goalDate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumProject.State\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"AcceptanceVoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBacking\",\"type\":\"uint256\"}],\"name\":\"BackingReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OwnerPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundPaid\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"checkIfFullyRefunded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDetails\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"projectStarter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"enumProject.State\",\"name\":\"currentState\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"currentAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountGoal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goalBacking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"milestone1Date\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"milestone2Date\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"milestone3Date\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processProjectStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumProject.State\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBacking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteForMilestone1Acceptance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteForMilestone2Acceptance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600860006101000a81548160ff021916908360068111156200002d576200002c620000e4565b5b02179055503480156200003f57600080fd5b5060405162001919380380620019198339818101604052810190620000659190620001b8565b856000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600181905550806002819055508460038190555083600481905550826005819055506000600681905550600060078190555050505050505062000254565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001458262000118565b9050919050565b620001578162000138565b81146200016357600080fd5b50565b60008151905062000177816200014c565b92915050565b6000819050919050565b62000192816200017d565b81146200019e57600080fd5b50565b600081519050620001b28162000187565b92915050565b60008060008060008060c08789031215620001d857620001d762000113565b5b6000620001e889828a0162000166565b9650506020620001fb89828a01620001a1565b95505060406200020e89828a01620001a1565b94505060606200022189828a01620001a1565b93505060806200023489828a01620001a1565b92505060a06200024789828a01620001a1565b9150509295509295509295565b6116b580620002646000396000f3fe6080604052600436106100e85760003560e01c80639caf85581161008a578063d7bb99ba11610059578063d7bb99ba1461028a578063eb2cd25814610294578063ed846e75146102bf578063fbbf93a0146102ea576100e8565b80639caf8558146101f2578063ba6c25a914610209578063c19d93fb14610234578063ce845d1d1461025f576100e8565b806337bfaee8116100c657806337bfaee81461015a5780636ae0f3e9146101855780638622dc7b1461019c5780638da5cb5b146101c7576100e8565b806311381f20146100ed5780631ecd3d131461010457806322bcfdfe1461012f575b600080fd5b3480156100f957600080fd5b5061010261031a565b005b34801561011057600080fd5b5061011961053b565b60405161012691906111de565b60405180910390f35b34801561013b57600080fd5b50610144610541565b60405161015191906111de565b60405180910390f35b34801561016657600080fd5b5061016f610547565b60405161017c91906111de565b60405180910390f35b34801561019157600080fd5b5061019a61054d565b005b3480156101a857600080fd5b506101b161062f565b6040516101be91906111de565b60405180910390f35b3480156101d357600080fd5b506101dc610635565b6040516101e9919061123a565b60405180910390f35b3480156101fe57600080fd5b50610207610659565b005b34801561021557600080fd5b5061021e61073b565b60405161022b91906111de565b60405180910390f35b34801561024057600080fd5b50610249610741565b60405161025691906112cc565b60405180910390f35b34801561026b57600080fd5b50610274610754565b60405161028191906111de565b60405180910390f35b61029261075a565b005b3480156102a057600080fd5b506102a9610997565b6040516102b691906111de565b60405180910390f35b3480156102cb57600080fd5b506102d461099d565b6040516102e19190611302565b60405180910390f35b3480156102f657600080fd5b506102ff610ae3565b6040516103119695949392919061131d565b60405180910390f35b6000600681111561032e5761032d611255565b5b600860009054906101000a900460ff1660068111156103505761034f611255565b5b14801561035e575060025442115b156103b9576001546007541061037d576103786003610b3d565b6103b4565b6001600860006101000a81548160ff021916908360068111156103a3576103a2611255565b5b02179055506103b26064610c82565b505b610539565b600360068111156103cd576103cc611255565b5b600860009054906101000a900460ff1660068111156103ef576103ee611255565b5b1480156103fd575060035442115b1561043d576002600860006101000a81548160ff0219169083600681111561042857610427611255565b5b02179055506104376046610c82565b50610538565b6004600681111561045157610450611255565b5b600860009054906101000a900460ff16600681111561047357610472611255565b5b148015610481575060045442115b156104c1576002600860006101000a81548160ff021916908360068111156104ac576104ab611255565b5b02179055506104bb601e610c82565b50610537565b600560068111156104d5576104d4611255565b5b600860009054906101000a900460ff1660068111156104f7576104f6611255565b5b148015610505575060055442115b15610536576006600860006101000a81548160ff021916908360068111156105305761052f611255565b5b02179055505b5b5b5b565b60055481565b60035481565b60045481565b6003600681111561056157610560611255565b5b600860009054906101000a900460ff16600681111561058357610582611255565b5b148015610591575060045442105b61059a57600080fd5b6000600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054116105e657600080fd5b6105f3600b336004610f72565b7f1fcb19b3ef2fd74857ee8e72f0ca1142284686dc6573eb891fc2b0d427ebed0833600360405161062592919061139f565b60405180910390a1565b60015481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6004600681111561066d5761066c611255565b5b600860009054906101000a900460ff16600681111561068f5761068e611255565b5b14801561069d575060055442105b6106a657600080fd5b6000600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054116106f257600080fd5b6106ff600c336005610f72565b7f1fcb19b3ef2fd74857ee8e72f0ca1142284686dc6573eb891fc2b0d427ebed0833600460405161073192919061139f565b60405180910390a1565b60025481565b600860009054906101000a900460ff1681565b60065481565b6000600681111561076e5761076d611255565b5b600860009054906101000a900460ff1660068111156107905761078f611255565b5b1461079a57600080fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16036107f257600080fd5b34600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461083d91906113f7565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600a339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055503460065461094991906113f7565b6006819055506006546007819055507fe1fca9ddd1c84fa92c6e1c6f0399ce9f5c1ff6e090b101c59a43403e47ed8bf2333460065460405161098d9392919061144d565b60405180910390a1565b60075481565b6000600260068111156109b3576109b2611255565b5b600860009054906101000a900460ff1660068111156109d5576109d4611255565b5b1480610a145750600160068111156109f0576109ef611255565b5b600860009054906101000a900460ff166006811115610a1257610a11611255565b5b145b610a1d57600080fd5b60006001905060005b600a80549050811015610adb57600d6000600a8381548110610a4b57610a4a611484565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610ac857600091505b8080610ad3906114b3565b915050610a26565b508091505090565b60008060008060008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1695506002549450600860009054906101000a900460ff169350600654925060075490506001549150909192939495565b60036006811115610b5157610b50611255565b5b816006811115610b6457610b63611255565b5b03610ba957610b73601e610fe3565b15610ba4576003600860006101000a81548160ff02191690836006811115610b9e57610b9d611255565b5b02179055505b610c7f565b60046006811115610bbd57610bbc611255565b5b816006811115610bd057610bcf611255565b5b03610c1557610bdf6028610fe3565b15610c10576004600860006101000a81548160ff02191690836006811115610c0a57610c09611255565b5b02179055505b610c7e565b60056006811115610c2957610c28611255565b5b816006811115610c3c57610c3b611255565b5b03610c7d57610c4b601e610fe3565b15610c7c576005600860006101000a81548160ff02191690836006811115610c7657610c75611255565b5b02179055505b5b5b5b50565b6000806000905060005b600a80549050811015610f68576000606485610ca8919061152a565b60096000600a8581548110610cc057610cbf611484565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610d30919061155b565b9050600d6000600a8481548110610d4a57610d49611484565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610f5457600a8281548110610dd657610dd5611484565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015610f4e577f856b7cb4b8bd821fadaa17b6a737dbd02cce301a5b121d1b01f3967c080ac409600a8381548110610e6f57610e6e611484565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168683604051610eab9392919061144d565b60405180910390a16000600d6000600a8581548110610ecd57610ecc611484565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550610f53565b600192505b5b508080610f60906114b3565b915050610c8c565b5080915050919050565b82829080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610fde83826110e0565b505050565b600080606483610ff3919061152a565b600754611000919061155b565b905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050156110d5577fa2e9126e4fe352d19a2e0868453f19ed604825e843000dea5f0e24904153769260008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684836040516110af93929190611614565b60405180910390a1806006546110c5919061164b565b60068190555060019150506110db565b60009150505b919050565b6000805b838054905081101561118d576009600085838154811061110757611106611484565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548261117891906113f7565b91508080611185906114b3565b9150506110e4565b5060006064600754836111a0919061152a565b6111aa919061155b565b905060468111156111bf576111be83610b3d565b5b50505050565b6000819050919050565b6111d8816111c5565b82525050565b60006020820190506111f360008301846111cf565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611224826111f9565b9050919050565b61123481611219565b82525050565b600060208201905061124f600083018461122b565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6007811061129557611294611255565b5b50565b60008190506112a682611284565b919050565b60006112b682611298565b9050919050565b6112c6816112ab565b82525050565b60006020820190506112e160008301846112bd565b92915050565b60008115159050919050565b6112fc816112e7565b82525050565b600060208201905061131760008301846112f3565b92915050565b600060c082019050611332600083018961122b565b61133f60208301886111cf565b61134c60408301876112bd565b61135960608301866111cf565b61136660808301856111cf565b61137360a08301846111cf565b979650505050505050565b6000611389826111f9565b9050919050565b6113998161137e565b82525050565b60006040820190506113b46000830185611390565b6113c160208301846112bd565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611402826111c5565b915061140d836111c5565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611442576114416113c8565b5b828201905092915050565b60006060820190506114626000830186611390565b61146f60208301856111cf565b61147c60408301846111cf565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006114be826111c5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036114f0576114ef6113c8565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611535826111c5565b9150611540836111c5565b9250826115505761154f6114fb565b5b828204905092915050565b6000611566826111c5565b9150611571836111c5565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156115aa576115a96113c8565b5b828202905092915050565b6000819050919050565b60006115da6115d56115d0846111f9565b6115b5565b6111f9565b9050919050565b60006115ec826115bf565b9050919050565b60006115fe826115e1565b9050919050565b61160e816115f3565b82525050565b60006060820190506116296000830186611605565b61163660208301856111cf565b61164360408301846111cf565b949350505050565b6000611656826111c5565b9150611661836111c5565b925082821015611674576116736113c8565b5b82820390509291505056fea264697066735822122023508374c83fcb3408f5e9d1f5e8732dac3f53490272312c6a06538825e395c064736f6c634300080e0033",
}

// ProjectABI is the input ABI used to generate the binding from.
// Deprecated: Use ProjectMetaData.ABI instead.
var ProjectABI = ProjectMetaData.ABI

// ProjectBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProjectMetaData.Bin instead.
var ProjectBin = ProjectMetaData.Bin

// DeployProject deploys a new Ethereum contract, binding an instance of Project to it.
func DeployProject(auth *bind.TransactOpts, backend bind.ContractBackend, projectCreator common.Address, milestone1Deadline *big.Int, milestone2Deadline *big.Int, milestone3Deadline *big.Int, goalAmount *big.Int, goalDate *big.Int) (common.Address, *types.Transaction, *Project, error) {
	parsed, err := ProjectMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProjectBin), backend, projectCreator, milestone1Deadline, milestone2Deadline, milestone3Deadline, goalAmount, goalDate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Project{ProjectCaller: ProjectCaller{contract: contract}, ProjectTransactor: ProjectTransactor{contract: contract}, ProjectFilterer: ProjectFilterer{contract: contract}}, nil
}

// Project is an auto generated Go binding around an Ethereum contract.
type Project struct {
	ProjectCaller     // Read-only binding to the contract
	ProjectTransactor // Write-only binding to the contract
	ProjectFilterer   // Log filterer for contract events
}

// ProjectCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProjectCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProjectTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProjectFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProjectSession struct {
	Contract     *Project          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProjectCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProjectCallerSession struct {
	Contract *ProjectCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ProjectTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProjectTransactorSession struct {
	Contract     *ProjectTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ProjectRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProjectRaw struct {
	Contract *Project // Generic contract binding to access the raw methods on
}

// ProjectCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProjectCallerRaw struct {
	Contract *ProjectCaller // Generic read-only contract binding to access the raw methods on
}

// ProjectTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProjectTransactorRaw struct {
	Contract *ProjectTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProject creates a new instance of Project, bound to a specific deployed contract.
func NewProject(address common.Address, backend bind.ContractBackend) (*Project, error) {
	contract, err := bindProject(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Project{ProjectCaller: ProjectCaller{contract: contract}, ProjectTransactor: ProjectTransactor{contract: contract}, ProjectFilterer: ProjectFilterer{contract: contract}}, nil
}

// NewProjectCaller creates a new read-only instance of Project, bound to a specific deployed contract.
func NewProjectCaller(address common.Address, caller bind.ContractCaller) (*ProjectCaller, error) {
	contract, err := bindProject(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectCaller{contract: contract}, nil
}

// NewProjectTransactor creates a new write-only instance of Project, bound to a specific deployed contract.
func NewProjectTransactor(address common.Address, transactor bind.ContractTransactor) (*ProjectTransactor, error) {
	contract, err := bindProject(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectTransactor{contract: contract}, nil
}

// NewProjectFilterer creates a new log filterer instance of Project, bound to a specific deployed contract.
func NewProjectFilterer(address common.Address, filterer bind.ContractFilterer) (*ProjectFilterer, error) {
	contract, err := bindProject(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProjectFilterer{contract: contract}, nil
}

// bindProject binds a generic wrapper to an already deployed contract.
func bindProject(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProjectABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Project *ProjectRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Project.Contract.ProjectCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Project *ProjectRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Project.Contract.ProjectTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Project *ProjectRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Project.Contract.ProjectTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Project *ProjectCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Project.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Project *ProjectTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Project.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Project *ProjectTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Project.Contract.contract.Transact(opts, method, params...)
}

// CheckIfFullyRefunded is a free data retrieval call binding the contract method 0xed846e75.
//
// Solidity: function checkIfFullyRefunded() view returns(bool)
func (_Project *ProjectCaller) CheckIfFullyRefunded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "checkIfFullyRefunded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckIfFullyRefunded is a free data retrieval call binding the contract method 0xed846e75.
//
// Solidity: function checkIfFullyRefunded() view returns(bool)
func (_Project *ProjectSession) CheckIfFullyRefunded() (bool, error) {
	return _Project.Contract.CheckIfFullyRefunded(&_Project.CallOpts)
}

// CheckIfFullyRefunded is a free data retrieval call binding the contract method 0xed846e75.
//
// Solidity: function checkIfFullyRefunded() view returns(bool)
func (_Project *ProjectCallerSession) CheckIfFullyRefunded() (bool, error) {
	return _Project.Contract.CheckIfFullyRefunded(&_Project.CallOpts)
}

// CurrentBalance is a free data retrieval call binding the contract method 0xce845d1d.
//
// Solidity: function currentBalance() view returns(uint256)
func (_Project *ProjectCaller) CurrentBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "currentBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentBalance is a free data retrieval call binding the contract method 0xce845d1d.
//
// Solidity: function currentBalance() view returns(uint256)
func (_Project *ProjectSession) CurrentBalance() (*big.Int, error) {
	return _Project.Contract.CurrentBalance(&_Project.CallOpts)
}

// CurrentBalance is a free data retrieval call binding the contract method 0xce845d1d.
//
// Solidity: function currentBalance() view returns(uint256)
func (_Project *ProjectCallerSession) CurrentBalance() (*big.Int, error) {
	return _Project.Contract.CurrentBalance(&_Project.CallOpts)
}

// GetDetails is a free data retrieval call binding the contract method 0xfbbf93a0.
//
// Solidity: function getDetails() view returns(address projectStarter, uint256 deadline, uint8 currentState, uint256 currentAmount, uint256 amountGoal, uint256 totalAmount)
func (_Project *ProjectCaller) GetDetails(opts *bind.CallOpts) (struct {
	ProjectStarter common.Address
	Deadline       *big.Int
	CurrentState   uint8
	CurrentAmount  *big.Int
	AmountGoal     *big.Int
	TotalAmount    *big.Int
}, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "getDetails")

	outstruct := new(struct {
		ProjectStarter common.Address
		Deadline       *big.Int
		CurrentState   uint8
		CurrentAmount  *big.Int
		AmountGoal     *big.Int
		TotalAmount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProjectStarter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Deadline = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CurrentState = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.CurrentAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AmountGoal = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDetails is a free data retrieval call binding the contract method 0xfbbf93a0.
//
// Solidity: function getDetails() view returns(address projectStarter, uint256 deadline, uint8 currentState, uint256 currentAmount, uint256 amountGoal, uint256 totalAmount)
func (_Project *ProjectSession) GetDetails() (struct {
	ProjectStarter common.Address
	Deadline       *big.Int
	CurrentState   uint8
	CurrentAmount  *big.Int
	AmountGoal     *big.Int
	TotalAmount    *big.Int
}, error) {
	return _Project.Contract.GetDetails(&_Project.CallOpts)
}

// GetDetails is a free data retrieval call binding the contract method 0xfbbf93a0.
//
// Solidity: function getDetails() view returns(address projectStarter, uint256 deadline, uint8 currentState, uint256 currentAmount, uint256 amountGoal, uint256 totalAmount)
func (_Project *ProjectCallerSession) GetDetails() (struct {
	ProjectStarter common.Address
	Deadline       *big.Int
	CurrentState   uint8
	CurrentAmount  *big.Int
	AmountGoal     *big.Int
	TotalAmount    *big.Int
}, error) {
	return _Project.Contract.GetDetails(&_Project.CallOpts)
}

// GoalBacking is a free data retrieval call binding the contract method 0x8622dc7b.
//
// Solidity: function goalBacking() view returns(uint256)
func (_Project *ProjectCaller) GoalBacking(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "goalBacking")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GoalBacking is a free data retrieval call binding the contract method 0x8622dc7b.
//
// Solidity: function goalBacking() view returns(uint256)
func (_Project *ProjectSession) GoalBacking() (*big.Int, error) {
	return _Project.Contract.GoalBacking(&_Project.CallOpts)
}

// GoalBacking is a free data retrieval call binding the contract method 0x8622dc7b.
//
// Solidity: function goalBacking() view returns(uint256)
func (_Project *ProjectCallerSession) GoalBacking() (*big.Int, error) {
	return _Project.Contract.GoalBacking(&_Project.CallOpts)
}

// GoalDeadline is a free data retrieval call binding the contract method 0xba6c25a9.
//
// Solidity: function goalDeadline() view returns(uint256)
func (_Project *ProjectCaller) GoalDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "goalDeadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GoalDeadline is a free data retrieval call binding the contract method 0xba6c25a9.
//
// Solidity: function goalDeadline() view returns(uint256)
func (_Project *ProjectSession) GoalDeadline() (*big.Int, error) {
	return _Project.Contract.GoalDeadline(&_Project.CallOpts)
}

// GoalDeadline is a free data retrieval call binding the contract method 0xba6c25a9.
//
// Solidity: function goalDeadline() view returns(uint256)
func (_Project *ProjectCallerSession) GoalDeadline() (*big.Int, error) {
	return _Project.Contract.GoalDeadline(&_Project.CallOpts)
}

// Milestone1Date is a free data retrieval call binding the contract method 0x22bcfdfe.
//
// Solidity: function milestone1Date() view returns(uint256)
func (_Project *ProjectCaller) Milestone1Date(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "milestone1Date")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Milestone1Date is a free data retrieval call binding the contract method 0x22bcfdfe.
//
// Solidity: function milestone1Date() view returns(uint256)
func (_Project *ProjectSession) Milestone1Date() (*big.Int, error) {
	return _Project.Contract.Milestone1Date(&_Project.CallOpts)
}

// Milestone1Date is a free data retrieval call binding the contract method 0x22bcfdfe.
//
// Solidity: function milestone1Date() view returns(uint256)
func (_Project *ProjectCallerSession) Milestone1Date() (*big.Int, error) {
	return _Project.Contract.Milestone1Date(&_Project.CallOpts)
}

// Milestone2Date is a free data retrieval call binding the contract method 0x37bfaee8.
//
// Solidity: function milestone2Date() view returns(uint256)
func (_Project *ProjectCaller) Milestone2Date(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "milestone2Date")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Milestone2Date is a free data retrieval call binding the contract method 0x37bfaee8.
//
// Solidity: function milestone2Date() view returns(uint256)
func (_Project *ProjectSession) Milestone2Date() (*big.Int, error) {
	return _Project.Contract.Milestone2Date(&_Project.CallOpts)
}

// Milestone2Date is a free data retrieval call binding the contract method 0x37bfaee8.
//
// Solidity: function milestone2Date() view returns(uint256)
func (_Project *ProjectCallerSession) Milestone2Date() (*big.Int, error) {
	return _Project.Contract.Milestone2Date(&_Project.CallOpts)
}

// Milestone3Date is a free data retrieval call binding the contract method 0x1ecd3d13.
//
// Solidity: function milestone3Date() view returns(uint256)
func (_Project *ProjectCaller) Milestone3Date(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "milestone3Date")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Milestone3Date is a free data retrieval call binding the contract method 0x1ecd3d13.
//
// Solidity: function milestone3Date() view returns(uint256)
func (_Project *ProjectSession) Milestone3Date() (*big.Int, error) {
	return _Project.Contract.Milestone3Date(&_Project.CallOpts)
}

// Milestone3Date is a free data retrieval call binding the contract method 0x1ecd3d13.
//
// Solidity: function milestone3Date() view returns(uint256)
func (_Project *ProjectCallerSession) Milestone3Date() (*big.Int, error) {
	return _Project.Contract.Milestone3Date(&_Project.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Project *ProjectCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Project *ProjectSession) Owner() (common.Address, error) {
	return _Project.Contract.Owner(&_Project.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Project *ProjectCallerSession) Owner() (common.Address, error) {
	return _Project.Contract.Owner(&_Project.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Project *ProjectCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Project *ProjectSession) State() (uint8, error) {
	return _Project.Contract.State(&_Project.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Project *ProjectCallerSession) State() (uint8, error) {
	return _Project.Contract.State(&_Project.CallOpts)
}

// TotalBacking is a free data retrieval call binding the contract method 0xeb2cd258.
//
// Solidity: function totalBacking() view returns(uint256)
func (_Project *ProjectCaller) TotalBacking(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "totalBacking")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBacking is a free data retrieval call binding the contract method 0xeb2cd258.
//
// Solidity: function totalBacking() view returns(uint256)
func (_Project *ProjectSession) TotalBacking() (*big.Int, error) {
	return _Project.Contract.TotalBacking(&_Project.CallOpts)
}

// TotalBacking is a free data retrieval call binding the contract method 0xeb2cd258.
//
// Solidity: function totalBacking() view returns(uint256)
func (_Project *ProjectCallerSession) TotalBacking() (*big.Int, error) {
	return _Project.Contract.TotalBacking(&_Project.CallOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Project *ProjectTransactor) Contribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "contribute")
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Project *ProjectSession) Contribute() (*types.Transaction, error) {
	return _Project.Contract.Contribute(&_Project.TransactOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Project *ProjectTransactorSession) Contribute() (*types.Transaction, error) {
	return _Project.Contract.Contribute(&_Project.TransactOpts)
}

// ProcessProjectStatus is a paid mutator transaction binding the contract method 0x11381f20.
//
// Solidity: function processProjectStatus() returns()
func (_Project *ProjectTransactor) ProcessProjectStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "processProjectStatus")
}

// ProcessProjectStatus is a paid mutator transaction binding the contract method 0x11381f20.
//
// Solidity: function processProjectStatus() returns()
func (_Project *ProjectSession) ProcessProjectStatus() (*types.Transaction, error) {
	return _Project.Contract.ProcessProjectStatus(&_Project.TransactOpts)
}

// ProcessProjectStatus is a paid mutator transaction binding the contract method 0x11381f20.
//
// Solidity: function processProjectStatus() returns()
func (_Project *ProjectTransactorSession) ProcessProjectStatus() (*types.Transaction, error) {
	return _Project.Contract.ProcessProjectStatus(&_Project.TransactOpts)
}

// VoteForMilestone1Acceptance is a paid mutator transaction binding the contract method 0x6ae0f3e9.
//
// Solidity: function voteForMilestone1Acceptance() returns()
func (_Project *ProjectTransactor) VoteForMilestone1Acceptance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "voteForMilestone1Acceptance")
}

// VoteForMilestone1Acceptance is a paid mutator transaction binding the contract method 0x6ae0f3e9.
//
// Solidity: function voteForMilestone1Acceptance() returns()
func (_Project *ProjectSession) VoteForMilestone1Acceptance() (*types.Transaction, error) {
	return _Project.Contract.VoteForMilestone1Acceptance(&_Project.TransactOpts)
}

// VoteForMilestone1Acceptance is a paid mutator transaction binding the contract method 0x6ae0f3e9.
//
// Solidity: function voteForMilestone1Acceptance() returns()
func (_Project *ProjectTransactorSession) VoteForMilestone1Acceptance() (*types.Transaction, error) {
	return _Project.Contract.VoteForMilestone1Acceptance(&_Project.TransactOpts)
}

// VoteForMilestone2Acceptance is a paid mutator transaction binding the contract method 0x9caf8558.
//
// Solidity: function voteForMilestone2Acceptance() returns()
func (_Project *ProjectTransactor) VoteForMilestone2Acceptance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "voteForMilestone2Acceptance")
}

// VoteForMilestone2Acceptance is a paid mutator transaction binding the contract method 0x9caf8558.
//
// Solidity: function voteForMilestone2Acceptance() returns()
func (_Project *ProjectSession) VoteForMilestone2Acceptance() (*types.Transaction, error) {
	return _Project.Contract.VoteForMilestone2Acceptance(&_Project.TransactOpts)
}

// VoteForMilestone2Acceptance is a paid mutator transaction binding the contract method 0x9caf8558.
//
// Solidity: function voteForMilestone2Acceptance() returns()
func (_Project *ProjectTransactorSession) VoteForMilestone2Acceptance() (*types.Transaction, error) {
	return _Project.Contract.VoteForMilestone2Acceptance(&_Project.TransactOpts)
}

// ProjectAcceptanceVoteCastIterator is returned from FilterAcceptanceVoteCast and is used to iterate over the raw logs and unpacked data for AcceptanceVoteCast events raised by the Project contract.
type ProjectAcceptanceVoteCastIterator struct {
	Event *ProjectAcceptanceVoteCast // Event containing the contract specifics and raw log

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
func (it *ProjectAcceptanceVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectAcceptanceVoteCast)
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
		it.Event = new(ProjectAcceptanceVoteCast)
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
func (it *ProjectAcceptanceVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectAcceptanceVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectAcceptanceVoteCast represents a AcceptanceVoteCast event raised by the Project contract.
type ProjectAcceptanceVoteCast struct {
	Contributor common.Address
	State       uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAcceptanceVoteCast is a free log retrieval operation binding the contract event 0x1fcb19b3ef2fd74857ee8e72f0ca1142284686dc6573eb891fc2b0d427ebed08.
//
// Solidity: event AcceptanceVoteCast(address contributor, uint8 state)
func (_Project *ProjectFilterer) FilterAcceptanceVoteCast(opts *bind.FilterOpts) (*ProjectAcceptanceVoteCastIterator, error) {

	logs, sub, err := _Project.contract.FilterLogs(opts, "AcceptanceVoteCast")
	if err != nil {
		return nil, err
	}
	return &ProjectAcceptanceVoteCastIterator{contract: _Project.contract, event: "AcceptanceVoteCast", logs: logs, sub: sub}, nil
}

// WatchAcceptanceVoteCast is a free log subscription operation binding the contract event 0x1fcb19b3ef2fd74857ee8e72f0ca1142284686dc6573eb891fc2b0d427ebed08.
//
// Solidity: event AcceptanceVoteCast(address contributor, uint8 state)
func (_Project *ProjectFilterer) WatchAcceptanceVoteCast(opts *bind.WatchOpts, sink chan<- *ProjectAcceptanceVoteCast) (event.Subscription, error) {

	logs, sub, err := _Project.contract.WatchLogs(opts, "AcceptanceVoteCast")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectAcceptanceVoteCast)
				if err := _Project.contract.UnpackLog(event, "AcceptanceVoteCast", log); err != nil {
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

// ParseAcceptanceVoteCast is a log parse operation binding the contract event 0x1fcb19b3ef2fd74857ee8e72f0ca1142284686dc6573eb891fc2b0d427ebed08.
//
// Solidity: event AcceptanceVoteCast(address contributor, uint8 state)
func (_Project *ProjectFilterer) ParseAcceptanceVoteCast(log types.Log) (*ProjectAcceptanceVoteCast, error) {
	event := new(ProjectAcceptanceVoteCast)
	if err := _Project.contract.UnpackLog(event, "AcceptanceVoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectBackingReceivedIterator is returned from FilterBackingReceived and is used to iterate over the raw logs and unpacked data for BackingReceived events raised by the Project contract.
type ProjectBackingReceivedIterator struct {
	Event *ProjectBackingReceived // Event containing the contract specifics and raw log

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
func (it *ProjectBackingReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectBackingReceived)
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
		it.Event = new(ProjectBackingReceived)
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
func (it *ProjectBackingReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectBackingReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectBackingReceived represents a BackingReceived event raised by the Project contract.
type ProjectBackingReceived struct {
	Contributor  common.Address
	Amount       *big.Int
	TotalBacking *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBackingReceived is a free log retrieval operation binding the contract event 0xe1fca9ddd1c84fa92c6e1c6f0399ce9f5c1ff6e090b101c59a43403e47ed8bf2.
//
// Solidity: event BackingReceived(address contributor, uint256 amount, uint256 totalBacking)
func (_Project *ProjectFilterer) FilterBackingReceived(opts *bind.FilterOpts) (*ProjectBackingReceivedIterator, error) {

	logs, sub, err := _Project.contract.FilterLogs(opts, "BackingReceived")
	if err != nil {
		return nil, err
	}
	return &ProjectBackingReceivedIterator{contract: _Project.contract, event: "BackingReceived", logs: logs, sub: sub}, nil
}

// WatchBackingReceived is a free log subscription operation binding the contract event 0xe1fca9ddd1c84fa92c6e1c6f0399ce9f5c1ff6e090b101c59a43403e47ed8bf2.
//
// Solidity: event BackingReceived(address contributor, uint256 amount, uint256 totalBacking)
func (_Project *ProjectFilterer) WatchBackingReceived(opts *bind.WatchOpts, sink chan<- *ProjectBackingReceived) (event.Subscription, error) {

	logs, sub, err := _Project.contract.WatchLogs(opts, "BackingReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectBackingReceived)
				if err := _Project.contract.UnpackLog(event, "BackingReceived", log); err != nil {
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

// ParseBackingReceived is a log parse operation binding the contract event 0xe1fca9ddd1c84fa92c6e1c6f0399ce9f5c1ff6e090b101c59a43403e47ed8bf2.
//
// Solidity: event BackingReceived(address contributor, uint256 amount, uint256 totalBacking)
func (_Project *ProjectFilterer) ParseBackingReceived(log types.Log) (*ProjectBackingReceived, error) {
	event := new(ProjectBackingReceived)
	if err := _Project.contract.UnpackLog(event, "BackingReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectOwnerPaidIterator is returned from FilterOwnerPaid and is used to iterate over the raw logs and unpacked data for OwnerPaid events raised by the Project contract.
type ProjectOwnerPaidIterator struct {
	Event *ProjectOwnerPaid // Event containing the contract specifics and raw log

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
func (it *ProjectOwnerPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectOwnerPaid)
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
		it.Event = new(ProjectOwnerPaid)
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
func (it *ProjectOwnerPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectOwnerPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectOwnerPaid represents a OwnerPaid event raised by the Project contract.
type ProjectOwnerPaid struct {
	Recipient  common.Address
	Percentage *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOwnerPaid is a free log retrieval operation binding the contract event 0xa2e9126e4fe352d19a2e0868453f19ed604825e843000dea5f0e249041537692.
//
// Solidity: event OwnerPaid(address recipient, uint256 percentage, uint256 amount)
func (_Project *ProjectFilterer) FilterOwnerPaid(opts *bind.FilterOpts) (*ProjectOwnerPaidIterator, error) {

	logs, sub, err := _Project.contract.FilterLogs(opts, "OwnerPaid")
	if err != nil {
		return nil, err
	}
	return &ProjectOwnerPaidIterator{contract: _Project.contract, event: "OwnerPaid", logs: logs, sub: sub}, nil
}

// WatchOwnerPaid is a free log subscription operation binding the contract event 0xa2e9126e4fe352d19a2e0868453f19ed604825e843000dea5f0e249041537692.
//
// Solidity: event OwnerPaid(address recipient, uint256 percentage, uint256 amount)
func (_Project *ProjectFilterer) WatchOwnerPaid(opts *bind.WatchOpts, sink chan<- *ProjectOwnerPaid) (event.Subscription, error) {

	logs, sub, err := _Project.contract.WatchLogs(opts, "OwnerPaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectOwnerPaid)
				if err := _Project.contract.UnpackLog(event, "OwnerPaid", log); err != nil {
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

// ParseOwnerPaid is a log parse operation binding the contract event 0xa2e9126e4fe352d19a2e0868453f19ed604825e843000dea5f0e249041537692.
//
// Solidity: event OwnerPaid(address recipient, uint256 percentage, uint256 amount)
func (_Project *ProjectFilterer) ParseOwnerPaid(log types.Log) (*ProjectOwnerPaid, error) {
	event := new(ProjectOwnerPaid)
	if err := _Project.contract.UnpackLog(event, "OwnerPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectRefundPaidIterator is returned from FilterRefundPaid and is used to iterate over the raw logs and unpacked data for RefundPaid events raised by the Project contract.
type ProjectRefundPaidIterator struct {
	Event *ProjectRefundPaid // Event containing the contract specifics and raw log

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
func (it *ProjectRefundPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectRefundPaid)
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
		it.Event = new(ProjectRefundPaid)
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
func (it *ProjectRefundPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectRefundPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectRefundPaid represents a RefundPaid event raised by the Project contract.
type ProjectRefundPaid struct {
	Recipient  common.Address
	Percentage *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRefundPaid is a free log retrieval operation binding the contract event 0x856b7cb4b8bd821fadaa17b6a737dbd02cce301a5b121d1b01f3967c080ac409.
//
// Solidity: event RefundPaid(address recipient, uint256 percentage, uint256 amount)
func (_Project *ProjectFilterer) FilterRefundPaid(opts *bind.FilterOpts) (*ProjectRefundPaidIterator, error) {

	logs, sub, err := _Project.contract.FilterLogs(opts, "RefundPaid")
	if err != nil {
		return nil, err
	}
	return &ProjectRefundPaidIterator{contract: _Project.contract, event: "RefundPaid", logs: logs, sub: sub}, nil
}

// WatchRefundPaid is a free log subscription operation binding the contract event 0x856b7cb4b8bd821fadaa17b6a737dbd02cce301a5b121d1b01f3967c080ac409.
//
// Solidity: event RefundPaid(address recipient, uint256 percentage, uint256 amount)
func (_Project *ProjectFilterer) WatchRefundPaid(opts *bind.WatchOpts, sink chan<- *ProjectRefundPaid) (event.Subscription, error) {

	logs, sub, err := _Project.contract.WatchLogs(opts, "RefundPaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectRefundPaid)
				if err := _Project.contract.UnpackLog(event, "RefundPaid", log); err != nil {
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

// ParseRefundPaid is a log parse operation binding the contract event 0x856b7cb4b8bd821fadaa17b6a737dbd02cce301a5b121d1b01f3967c080ac409.
//
// Solidity: event RefundPaid(address recipient, uint256 percentage, uint256 amount)
func (_Project *ProjectFilterer) ParseRefundPaid(log types.Log) (*ProjectRefundPaid, error) {
	event := new(ProjectRefundPaid)
	if err := _Project.contract.UnpackLog(event, "RefundPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
