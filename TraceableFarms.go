// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TraceableFarms

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

// TraceableFarmsAccreditation is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsAccreditation struct {
	Name                     string
	Description              string
	InformationalResourceUrl string
	AccreditationType        TraceableFarmsAccreditationType
}

// TraceableFarmsAccreditationType is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsAccreditationType struct {
	Name string
}

// TraceableFarmsBatch is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsBatch struct {
	Number                    string
	Date                      string
	ProductName               string
	ProductVariety            string
	ProductAppearance         string
	ProductSize               string
	ProductDescription        string
	ProductPhotoUrl           string
	ProductStatisticsImageUrl string
	Company                   TraceableFarmsCompany
}

// TraceableFarmsBatchFootprint is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsBatchFootprint struct {
	FootprintType TraceableFarmsFootprintType
	TotalValue    *big.Int
	CheckerHash   string
	CheckerUrl    string
}

// TraceableFarmsBatchFootprintValue is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsBatchFootprintValue struct {
	Description string
	Value       *big.Int
	CheckerHash string
	CheckerUrl  string
}

// TraceableFarmsBatchOrigin is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsBatchOrigin struct {
	Description         string
	Location            string
	LocationCoordinates string
}

// TraceableFarmsBatchProcess is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsBatchProcess struct {
	Name        string
	Description string
}

// TraceableFarmsCompany is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsCompany struct {
	Nif                      string
	BussinessName            string
	Description              string
	Location                 string
	LocationCoordinates      string
	InformationalResourceUrl string
	Consortium               TraceableFarmsConsortium
}

// TraceableFarmsCompanyAccreditation is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsCompanyAccreditation struct {
	Accreditation TraceableFarmsAccreditation
	CheckerHash   string
	CheckerUrl    string
}

// TraceableFarmsCompanyFootprintReportability is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsCompanyFootprintReportability struct {
	Company             TraceableFarmsCompany
	FootprintType       TraceableFarmsFootprintType
	FootprintReportType TraceableFarmsFootprintReportType
	Description         string
}

// TraceableFarmsConsortium is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsConsortium struct {
	Name string
}

// TraceableFarmsFootprintReportType is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsFootprintReportType struct {
	Name string
}

// TraceableFarmsFootprintType is an auto generated low-level Go binding around an user-defined struct.
type TraceableFarmsFootprintType struct {
	Name                  string
	Description           string
	UnitMeasurementName   string
	UnitMeasurementSymbol string
}

// TraceableFarmsMetaData contains all meta data concerning the TraceableFarms contract.
var TraceableFarmsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getAccreditation\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"informationalResourceUrl\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.AccreditationType\",\"name\":\"accreditationType\",\"type\":\"tuple\"}],\"internalType\":\"structTraceableFarms.Accreditation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getAccreditationType\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.AccreditationType\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_companyNif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_number\",\"type\":\"string\"}],\"name\":\"getBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"number\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productVariety\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productAppearance\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productSize\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productDescription\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productPhotoUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"productStatisticsImageUrl\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"nif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bussinessName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"locationCoordinates\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"informationalResourceUrl\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.Consortium\",\"name\":\"consortium\",\"type\":\"tuple\"}],\"internalType\":\"structTraceableFarms.Company\",\"name\":\"company\",\"type\":\"tuple\"}],\"internalType\":\"structTraceableFarms.Batch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_companyNif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_batchNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_processName\",\"type\":\"string\"}],\"name\":\"getBatchFootprint\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unitMeasurementName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unitMeasurementSymbol\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.FootprintType\",\"name\":\"footprintType\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"checkerHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"checkerUrl\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.BatchFootprint[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_companyNif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_batchNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_processName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_footprintTypeName\",\"type\":\"string\"}],\"name\":\"getBatchFootprintValue\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"checkerHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"checkerUrl\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.BatchFootprintValue[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_companyNif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_batchNumber\",\"type\":\"string\"}],\"name\":\"getBatchOrigin\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"locationCoordinates\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.BatchOrigin[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_companyNif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_batchNumber\",\"type\":\"string\"}],\"name\":\"getBatchProcess\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.BatchProcess[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nif\",\"type\":\"string\"}],\"name\":\"getCompany\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bussinessName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"locationCoordinates\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"informationalResourceUrl\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.Consortium\",\"name\":\"consortium\",\"type\":\"tuple\"}],\"internalType\":\"structTraceableFarms.Company\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nif\",\"type\":\"string\"}],\"name\":\"getCompanyAccreditation\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"informationalResourceUrl\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.AccreditationType\",\"name\":\"accreditationType\",\"type\":\"tuple\"}],\"internalType\":\"structTraceableFarms.Accreditation\",\"name\":\"accreditation\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"checkerHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"checkerUrl\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.CompanyAccreditation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_footprintTypeName\",\"type\":\"string\"}],\"name\":\"getCompanyFootprintReportability\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bussinessName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"locationCoordinates\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"informationalResourceUrl\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.Consortium\",\"name\":\"consortium\",\"type\":\"tuple\"}],\"internalType\":\"structTraceableFarms.Company\",\"name\":\"company\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unitMeasurementName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unitMeasurementSymbol\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.FootprintType\",\"name\":\"footprintType\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.FootprintReportType\",\"name\":\"footprintReportType\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.CompanyFootprintReportability[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getConsortium\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.Consortium\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getFootprintReportType\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.FootprintReportType\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getFootprintType\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unitMeasurementName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unitMeasurementSymbol\",\"type\":\"string\"}],\"internalType\":\"structTraceableFarms.FootprintType\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_informationalResourceUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_accreditationTypeName\",\"type\":\"string\"}],\"name\":\"setAccreditation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setAccreditationType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_number\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_date\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_productName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_productVariety\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_productAppearance\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_productSize\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_productDescription\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_productPhotoUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_productStatisticsImageUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_companyNif\",\"type\":\"string\"}],\"name\":\"setBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_companyNif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_batchNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_processName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_footprintTypeName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_totalValue\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_checkerHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_checkerUrl\",\"type\":\"string\"}],\"name\":\"setBatchFootprint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_companyNif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_batchNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_processName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_footprintTypeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_checkerHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_checkerUrl\",\"type\":\"string\"}],\"name\":\"setBatchFootprintValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_companyNif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_batchNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_locationCoordinates\",\"type\":\"string\"}],\"name\":\"setBatchOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_companyNif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_batchNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"setBatchProcess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_bussinessName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_locationCoordinates\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_informationalResourceUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_consortiumName\",\"type\":\"string\"}],\"name\":\"setCompany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_accreditationName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_checkerHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_checkerUrl\",\"type\":\"string\"}],\"name\":\"setCompanyAccreditation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_footprintTypeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_footprintReportTypeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"setCompanyFootprintReportability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setConsortium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setFootprintReportType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_unitMeasurementName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_unitMeasurementSymbol\",\"type\":\"string\"}],\"name\":\"setFootprintType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TraceableFarmsABI is the input ABI used to generate the binding from.
// Deprecated: Use TraceableFarmsMetaData.ABI instead.
var TraceableFarmsABI = TraceableFarmsMetaData.ABI

// TraceableFarms is an auto generated Go binding around an Ethereum contract.
type TraceableFarms struct {
	TraceableFarmsCaller     // Read-only binding to the contract
	TraceableFarmsTransactor // Write-only binding to the contract
	TraceableFarmsFilterer   // Log filterer for contract events
}

// TraceableFarmsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraceableFarmsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceableFarmsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraceableFarmsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceableFarmsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraceableFarmsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceableFarmsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraceableFarmsSession struct {
	Contract     *TraceableFarms   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceableFarmsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraceableFarmsCallerSession struct {
	Contract *TraceableFarmsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TraceableFarmsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraceableFarmsTransactorSession struct {
	Contract     *TraceableFarmsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TraceableFarmsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraceableFarmsRaw struct {
	Contract *TraceableFarms // Generic contract binding to access the raw methods on
}

// TraceableFarmsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraceableFarmsCallerRaw struct {
	Contract *TraceableFarmsCaller // Generic read-only contract binding to access the raw methods on
}

// TraceableFarmsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraceableFarmsTransactorRaw struct {
	Contract *TraceableFarmsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTraceableFarms creates a new instance of TraceableFarms, bound to a specific deployed contract.
func NewTraceableFarms(address common.Address, backend bind.ContractBackend) (*TraceableFarms, error) {
	contract, err := bindTraceableFarms(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TraceableFarms{TraceableFarmsCaller: TraceableFarmsCaller{contract: contract}, TraceableFarmsTransactor: TraceableFarmsTransactor{contract: contract}, TraceableFarmsFilterer: TraceableFarmsFilterer{contract: contract}}, nil
}

// NewTraceableFarmsCaller creates a new read-only instance of TraceableFarms, bound to a specific deployed contract.
func NewTraceableFarmsCaller(address common.Address, caller bind.ContractCaller) (*TraceableFarmsCaller, error) {
	contract, err := bindTraceableFarms(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraceableFarmsCaller{contract: contract}, nil
}

// NewTraceableFarmsTransactor creates a new write-only instance of TraceableFarms, bound to a specific deployed contract.
func NewTraceableFarmsTransactor(address common.Address, transactor bind.ContractTransactor) (*TraceableFarmsTransactor, error) {
	contract, err := bindTraceableFarms(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraceableFarmsTransactor{contract: contract}, nil
}

// NewTraceableFarmsFilterer creates a new log filterer instance of TraceableFarms, bound to a specific deployed contract.
func NewTraceableFarmsFilterer(address common.Address, filterer bind.ContractFilterer) (*TraceableFarmsFilterer, error) {
	contract, err := bindTraceableFarms(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraceableFarmsFilterer{contract: contract}, nil
}

// bindTraceableFarms binds a generic wrapper to an already deployed contract.
func bindTraceableFarms(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TraceableFarmsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraceableFarms *TraceableFarmsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraceableFarms.Contract.TraceableFarmsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraceableFarms *TraceableFarmsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraceableFarms.Contract.TraceableFarmsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraceableFarms *TraceableFarmsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraceableFarms.Contract.TraceableFarmsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraceableFarms *TraceableFarmsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraceableFarms.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraceableFarms *TraceableFarmsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraceableFarms.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraceableFarms *TraceableFarmsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraceableFarms.Contract.contract.Transact(opts, method, params...)
}

// GetAccreditation is a free data retrieval call binding the contract method 0x01384d8c.
//
// Solidity: function getAccreditation(string _name) view returns((string,string,string,(string)))
func (_TraceableFarms *TraceableFarmsCaller) GetAccreditation(opts *bind.CallOpts, _name string) (TraceableFarmsAccreditation, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getAccreditation", _name)

	if err != nil {
		return *new(TraceableFarmsAccreditation), err
	}

	out0 := *abi.ConvertType(out[0], new(TraceableFarmsAccreditation)).(*TraceableFarmsAccreditation)

	return out0, err

}

// GetAccreditation is a free data retrieval call binding the contract method 0x01384d8c.
//
// Solidity: function getAccreditation(string _name) view returns((string,string,string,(string)))
func (_TraceableFarms *TraceableFarmsSession) GetAccreditation(_name string) (TraceableFarmsAccreditation, error) {
	return _TraceableFarms.Contract.GetAccreditation(&_TraceableFarms.CallOpts, _name)
}

// GetAccreditation is a free data retrieval call binding the contract method 0x01384d8c.
//
// Solidity: function getAccreditation(string _name) view returns((string,string,string,(string)))
func (_TraceableFarms *TraceableFarmsCallerSession) GetAccreditation(_name string) (TraceableFarmsAccreditation, error) {
	return _TraceableFarms.Contract.GetAccreditation(&_TraceableFarms.CallOpts, _name)
}

// GetAccreditationType is a free data retrieval call binding the contract method 0x10c74bd2.
//
// Solidity: function getAccreditationType(string _name) view returns((string))
func (_TraceableFarms *TraceableFarmsCaller) GetAccreditationType(opts *bind.CallOpts, _name string) (TraceableFarmsAccreditationType, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getAccreditationType", _name)

	if err != nil {
		return *new(TraceableFarmsAccreditationType), err
	}

	out0 := *abi.ConvertType(out[0], new(TraceableFarmsAccreditationType)).(*TraceableFarmsAccreditationType)

	return out0, err

}

// GetAccreditationType is a free data retrieval call binding the contract method 0x10c74bd2.
//
// Solidity: function getAccreditationType(string _name) view returns((string))
func (_TraceableFarms *TraceableFarmsSession) GetAccreditationType(_name string) (TraceableFarmsAccreditationType, error) {
	return _TraceableFarms.Contract.GetAccreditationType(&_TraceableFarms.CallOpts, _name)
}

// GetAccreditationType is a free data retrieval call binding the contract method 0x10c74bd2.
//
// Solidity: function getAccreditationType(string _name) view returns((string))
func (_TraceableFarms *TraceableFarmsCallerSession) GetAccreditationType(_name string) (TraceableFarmsAccreditationType, error) {
	return _TraceableFarms.Contract.GetAccreditationType(&_TraceableFarms.CallOpts, _name)
}

// GetBatch is a free data retrieval call binding the contract method 0x810bc2ea.
//
// Solidity: function getBatch(string _companyNif, string _number) view returns((string,string,string,string,string,string,string,string,string,(string,string,string,string,string,string,(string))))
func (_TraceableFarms *TraceableFarmsCaller) GetBatch(opts *bind.CallOpts, _companyNif string, _number string) (TraceableFarmsBatch, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getBatch", _companyNif, _number)

	if err != nil {
		return *new(TraceableFarmsBatch), err
	}

	out0 := *abi.ConvertType(out[0], new(TraceableFarmsBatch)).(*TraceableFarmsBatch)

	return out0, err

}

// GetBatch is a free data retrieval call binding the contract method 0x810bc2ea.
//
// Solidity: function getBatch(string _companyNif, string _number) view returns((string,string,string,string,string,string,string,string,string,(string,string,string,string,string,string,(string))))
func (_TraceableFarms *TraceableFarmsSession) GetBatch(_companyNif string, _number string) (TraceableFarmsBatch, error) {
	return _TraceableFarms.Contract.GetBatch(&_TraceableFarms.CallOpts, _companyNif, _number)
}

// GetBatch is a free data retrieval call binding the contract method 0x810bc2ea.
//
// Solidity: function getBatch(string _companyNif, string _number) view returns((string,string,string,string,string,string,string,string,string,(string,string,string,string,string,string,(string))))
func (_TraceableFarms *TraceableFarmsCallerSession) GetBatch(_companyNif string, _number string) (TraceableFarmsBatch, error) {
	return _TraceableFarms.Contract.GetBatch(&_TraceableFarms.CallOpts, _companyNif, _number)
}

// GetBatchFootprint is a free data retrieval call binding the contract method 0xd1c173c2.
//
// Solidity: function getBatchFootprint(string _companyNif, string _batchNumber, string _processName) view returns(((string,string,string,string),uint256,string,string)[])
func (_TraceableFarms *TraceableFarmsCaller) GetBatchFootprint(opts *bind.CallOpts, _companyNif string, _batchNumber string, _processName string) ([]TraceableFarmsBatchFootprint, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getBatchFootprint", _companyNif, _batchNumber, _processName)

	if err != nil {
		return *new([]TraceableFarmsBatchFootprint), err
	}

	out0 := *abi.ConvertType(out[0], new([]TraceableFarmsBatchFootprint)).(*[]TraceableFarmsBatchFootprint)

	return out0, err

}

// GetBatchFootprint is a free data retrieval call binding the contract method 0xd1c173c2.
//
// Solidity: function getBatchFootprint(string _companyNif, string _batchNumber, string _processName) view returns(((string,string,string,string),uint256,string,string)[])
func (_TraceableFarms *TraceableFarmsSession) GetBatchFootprint(_companyNif string, _batchNumber string, _processName string) ([]TraceableFarmsBatchFootprint, error) {
	return _TraceableFarms.Contract.GetBatchFootprint(&_TraceableFarms.CallOpts, _companyNif, _batchNumber, _processName)
}

// GetBatchFootprint is a free data retrieval call binding the contract method 0xd1c173c2.
//
// Solidity: function getBatchFootprint(string _companyNif, string _batchNumber, string _processName) view returns(((string,string,string,string),uint256,string,string)[])
func (_TraceableFarms *TraceableFarmsCallerSession) GetBatchFootprint(_companyNif string, _batchNumber string, _processName string) ([]TraceableFarmsBatchFootprint, error) {
	return _TraceableFarms.Contract.GetBatchFootprint(&_TraceableFarms.CallOpts, _companyNif, _batchNumber, _processName)
}

// GetBatchFootprintValue is a free data retrieval call binding the contract method 0x465af2ae.
//
// Solidity: function getBatchFootprintValue(string _companyNif, string _batchNumber, string _processName, string _footprintTypeName) view returns((string,uint256,string,string)[])
func (_TraceableFarms *TraceableFarmsCaller) GetBatchFootprintValue(opts *bind.CallOpts, _companyNif string, _batchNumber string, _processName string, _footprintTypeName string) ([]TraceableFarmsBatchFootprintValue, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getBatchFootprintValue", _companyNif, _batchNumber, _processName, _footprintTypeName)

	if err != nil {
		return *new([]TraceableFarmsBatchFootprintValue), err
	}

	out0 := *abi.ConvertType(out[0], new([]TraceableFarmsBatchFootprintValue)).(*[]TraceableFarmsBatchFootprintValue)

	return out0, err

}

// GetBatchFootprintValue is a free data retrieval call binding the contract method 0x465af2ae.
//
// Solidity: function getBatchFootprintValue(string _companyNif, string _batchNumber, string _processName, string _footprintTypeName) view returns((string,uint256,string,string)[])
func (_TraceableFarms *TraceableFarmsSession) GetBatchFootprintValue(_companyNif string, _batchNumber string, _processName string, _footprintTypeName string) ([]TraceableFarmsBatchFootprintValue, error) {
	return _TraceableFarms.Contract.GetBatchFootprintValue(&_TraceableFarms.CallOpts, _companyNif, _batchNumber, _processName, _footprintTypeName)
}

// GetBatchFootprintValue is a free data retrieval call binding the contract method 0x465af2ae.
//
// Solidity: function getBatchFootprintValue(string _companyNif, string _batchNumber, string _processName, string _footprintTypeName) view returns((string,uint256,string,string)[])
func (_TraceableFarms *TraceableFarmsCallerSession) GetBatchFootprintValue(_companyNif string, _batchNumber string, _processName string, _footprintTypeName string) ([]TraceableFarmsBatchFootprintValue, error) {
	return _TraceableFarms.Contract.GetBatchFootprintValue(&_TraceableFarms.CallOpts, _companyNif, _batchNumber, _processName, _footprintTypeName)
}

// GetBatchOrigin is a free data retrieval call binding the contract method 0x6f8e5aa1.
//
// Solidity: function getBatchOrigin(string _companyNif, string _batchNumber) view returns((string,string,string)[])
func (_TraceableFarms *TraceableFarmsCaller) GetBatchOrigin(opts *bind.CallOpts, _companyNif string, _batchNumber string) ([]TraceableFarmsBatchOrigin, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getBatchOrigin", _companyNif, _batchNumber)

	if err != nil {
		return *new([]TraceableFarmsBatchOrigin), err
	}

	out0 := *abi.ConvertType(out[0], new([]TraceableFarmsBatchOrigin)).(*[]TraceableFarmsBatchOrigin)

	return out0, err

}

// GetBatchOrigin is a free data retrieval call binding the contract method 0x6f8e5aa1.
//
// Solidity: function getBatchOrigin(string _companyNif, string _batchNumber) view returns((string,string,string)[])
func (_TraceableFarms *TraceableFarmsSession) GetBatchOrigin(_companyNif string, _batchNumber string) ([]TraceableFarmsBatchOrigin, error) {
	return _TraceableFarms.Contract.GetBatchOrigin(&_TraceableFarms.CallOpts, _companyNif, _batchNumber)
}

// GetBatchOrigin is a free data retrieval call binding the contract method 0x6f8e5aa1.
//
// Solidity: function getBatchOrigin(string _companyNif, string _batchNumber) view returns((string,string,string)[])
func (_TraceableFarms *TraceableFarmsCallerSession) GetBatchOrigin(_companyNif string, _batchNumber string) ([]TraceableFarmsBatchOrigin, error) {
	return _TraceableFarms.Contract.GetBatchOrigin(&_TraceableFarms.CallOpts, _companyNif, _batchNumber)
}

// GetBatchProcess is a free data retrieval call binding the contract method 0xa6fc05ee.
//
// Solidity: function getBatchProcess(string _companyNif, string _batchNumber) view returns((string,string)[])
func (_TraceableFarms *TraceableFarmsCaller) GetBatchProcess(opts *bind.CallOpts, _companyNif string, _batchNumber string) ([]TraceableFarmsBatchProcess, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getBatchProcess", _companyNif, _batchNumber)

	if err != nil {
		return *new([]TraceableFarmsBatchProcess), err
	}

	out0 := *abi.ConvertType(out[0], new([]TraceableFarmsBatchProcess)).(*[]TraceableFarmsBatchProcess)

	return out0, err

}

// GetBatchProcess is a free data retrieval call binding the contract method 0xa6fc05ee.
//
// Solidity: function getBatchProcess(string _companyNif, string _batchNumber) view returns((string,string)[])
func (_TraceableFarms *TraceableFarmsSession) GetBatchProcess(_companyNif string, _batchNumber string) ([]TraceableFarmsBatchProcess, error) {
	return _TraceableFarms.Contract.GetBatchProcess(&_TraceableFarms.CallOpts, _companyNif, _batchNumber)
}

// GetBatchProcess is a free data retrieval call binding the contract method 0xa6fc05ee.
//
// Solidity: function getBatchProcess(string _companyNif, string _batchNumber) view returns((string,string)[])
func (_TraceableFarms *TraceableFarmsCallerSession) GetBatchProcess(_companyNif string, _batchNumber string) ([]TraceableFarmsBatchProcess, error) {
	return _TraceableFarms.Contract.GetBatchProcess(&_TraceableFarms.CallOpts, _companyNif, _batchNumber)
}

// GetCompany is a free data retrieval call binding the contract method 0xb24a4e52.
//
// Solidity: function getCompany(string _nif) view returns((string,string,string,string,string,string,(string)))
func (_TraceableFarms *TraceableFarmsCaller) GetCompany(opts *bind.CallOpts, _nif string) (TraceableFarmsCompany, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getCompany", _nif)

	if err != nil {
		return *new(TraceableFarmsCompany), err
	}

	out0 := *abi.ConvertType(out[0], new(TraceableFarmsCompany)).(*TraceableFarmsCompany)

	return out0, err

}

// GetCompany is a free data retrieval call binding the contract method 0xb24a4e52.
//
// Solidity: function getCompany(string _nif) view returns((string,string,string,string,string,string,(string)))
func (_TraceableFarms *TraceableFarmsSession) GetCompany(_nif string) (TraceableFarmsCompany, error) {
	return _TraceableFarms.Contract.GetCompany(&_TraceableFarms.CallOpts, _nif)
}

// GetCompany is a free data retrieval call binding the contract method 0xb24a4e52.
//
// Solidity: function getCompany(string _nif) view returns((string,string,string,string,string,string,(string)))
func (_TraceableFarms *TraceableFarmsCallerSession) GetCompany(_nif string) (TraceableFarmsCompany, error) {
	return _TraceableFarms.Contract.GetCompany(&_TraceableFarms.CallOpts, _nif)
}

// GetCompanyAccreditation is a free data retrieval call binding the contract method 0x78f2ea6f.
//
// Solidity: function getCompanyAccreditation(string _nif) view returns(((string,string,string,(string)),string,string)[])
func (_TraceableFarms *TraceableFarmsCaller) GetCompanyAccreditation(opts *bind.CallOpts, _nif string) ([]TraceableFarmsCompanyAccreditation, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getCompanyAccreditation", _nif)

	if err != nil {
		return *new([]TraceableFarmsCompanyAccreditation), err
	}

	out0 := *abi.ConvertType(out[0], new([]TraceableFarmsCompanyAccreditation)).(*[]TraceableFarmsCompanyAccreditation)

	return out0, err

}

// GetCompanyAccreditation is a free data retrieval call binding the contract method 0x78f2ea6f.
//
// Solidity: function getCompanyAccreditation(string _nif) view returns(((string,string,string,(string)),string,string)[])
func (_TraceableFarms *TraceableFarmsSession) GetCompanyAccreditation(_nif string) ([]TraceableFarmsCompanyAccreditation, error) {
	return _TraceableFarms.Contract.GetCompanyAccreditation(&_TraceableFarms.CallOpts, _nif)
}

// GetCompanyAccreditation is a free data retrieval call binding the contract method 0x78f2ea6f.
//
// Solidity: function getCompanyAccreditation(string _nif) view returns(((string,string,string,(string)),string,string)[])
func (_TraceableFarms *TraceableFarmsCallerSession) GetCompanyAccreditation(_nif string) ([]TraceableFarmsCompanyAccreditation, error) {
	return _TraceableFarms.Contract.GetCompanyAccreditation(&_TraceableFarms.CallOpts, _nif)
}

// GetCompanyFootprintReportability is a free data retrieval call binding the contract method 0xb5b14f23.
//
// Solidity: function getCompanyFootprintReportability(string _nif, string _footprintTypeName) view returns(((string,string,string,string,string,string,(string)),(string,string,string,string),(string),string)[])
func (_TraceableFarms *TraceableFarmsCaller) GetCompanyFootprintReportability(opts *bind.CallOpts, _nif string, _footprintTypeName string) ([]TraceableFarmsCompanyFootprintReportability, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getCompanyFootprintReportability", _nif, _footprintTypeName)

	if err != nil {
		return *new([]TraceableFarmsCompanyFootprintReportability), err
	}

	out0 := *abi.ConvertType(out[0], new([]TraceableFarmsCompanyFootprintReportability)).(*[]TraceableFarmsCompanyFootprintReportability)

	return out0, err

}

// GetCompanyFootprintReportability is a free data retrieval call binding the contract method 0xb5b14f23.
//
// Solidity: function getCompanyFootprintReportability(string _nif, string _footprintTypeName) view returns(((string,string,string,string,string,string,(string)),(string,string,string,string),(string),string)[])
func (_TraceableFarms *TraceableFarmsSession) GetCompanyFootprintReportability(_nif string, _footprintTypeName string) ([]TraceableFarmsCompanyFootprintReportability, error) {
	return _TraceableFarms.Contract.GetCompanyFootprintReportability(&_TraceableFarms.CallOpts, _nif, _footprintTypeName)
}

// GetCompanyFootprintReportability is a free data retrieval call binding the contract method 0xb5b14f23.
//
// Solidity: function getCompanyFootprintReportability(string _nif, string _footprintTypeName) view returns(((string,string,string,string,string,string,(string)),(string,string,string,string),(string),string)[])
func (_TraceableFarms *TraceableFarmsCallerSession) GetCompanyFootprintReportability(_nif string, _footprintTypeName string) ([]TraceableFarmsCompanyFootprintReportability, error) {
	return _TraceableFarms.Contract.GetCompanyFootprintReportability(&_TraceableFarms.CallOpts, _nif, _footprintTypeName)
}

// GetConsortium is a free data retrieval call binding the contract method 0x86dae8ee.
//
// Solidity: function getConsortium(string _name) view returns((string))
func (_TraceableFarms *TraceableFarmsCaller) GetConsortium(opts *bind.CallOpts, _name string) (TraceableFarmsConsortium, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getConsortium", _name)

	if err != nil {
		return *new(TraceableFarmsConsortium), err
	}

	out0 := *abi.ConvertType(out[0], new(TraceableFarmsConsortium)).(*TraceableFarmsConsortium)

	return out0, err

}

// GetConsortium is a free data retrieval call binding the contract method 0x86dae8ee.
//
// Solidity: function getConsortium(string _name) view returns((string))
func (_TraceableFarms *TraceableFarmsSession) GetConsortium(_name string) (TraceableFarmsConsortium, error) {
	return _TraceableFarms.Contract.GetConsortium(&_TraceableFarms.CallOpts, _name)
}

// GetConsortium is a free data retrieval call binding the contract method 0x86dae8ee.
//
// Solidity: function getConsortium(string _name) view returns((string))
func (_TraceableFarms *TraceableFarmsCallerSession) GetConsortium(_name string) (TraceableFarmsConsortium, error) {
	return _TraceableFarms.Contract.GetConsortium(&_TraceableFarms.CallOpts, _name)
}

// GetFootprintReportType is a free data retrieval call binding the contract method 0x26f0d89b.
//
// Solidity: function getFootprintReportType(string _name) view returns((string))
func (_TraceableFarms *TraceableFarmsCaller) GetFootprintReportType(opts *bind.CallOpts, _name string) (TraceableFarmsFootprintReportType, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getFootprintReportType", _name)

	if err != nil {
		return *new(TraceableFarmsFootprintReportType), err
	}

	out0 := *abi.ConvertType(out[0], new(TraceableFarmsFootprintReportType)).(*TraceableFarmsFootprintReportType)

	return out0, err

}

// GetFootprintReportType is a free data retrieval call binding the contract method 0x26f0d89b.
//
// Solidity: function getFootprintReportType(string _name) view returns((string))
func (_TraceableFarms *TraceableFarmsSession) GetFootprintReportType(_name string) (TraceableFarmsFootprintReportType, error) {
	return _TraceableFarms.Contract.GetFootprintReportType(&_TraceableFarms.CallOpts, _name)
}

// GetFootprintReportType is a free data retrieval call binding the contract method 0x26f0d89b.
//
// Solidity: function getFootprintReportType(string _name) view returns((string))
func (_TraceableFarms *TraceableFarmsCallerSession) GetFootprintReportType(_name string) (TraceableFarmsFootprintReportType, error) {
	return _TraceableFarms.Contract.GetFootprintReportType(&_TraceableFarms.CallOpts, _name)
}

// GetFootprintType is a free data retrieval call binding the contract method 0x0ba33f01.
//
// Solidity: function getFootprintType(string _name) view returns((string,string,string,string))
func (_TraceableFarms *TraceableFarmsCaller) GetFootprintType(opts *bind.CallOpts, _name string) (TraceableFarmsFootprintType, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "getFootprintType", _name)

	if err != nil {
		return *new(TraceableFarmsFootprintType), err
	}

	out0 := *abi.ConvertType(out[0], new(TraceableFarmsFootprintType)).(*TraceableFarmsFootprintType)

	return out0, err

}

// GetFootprintType is a free data retrieval call binding the contract method 0x0ba33f01.
//
// Solidity: function getFootprintType(string _name) view returns((string,string,string,string))
func (_TraceableFarms *TraceableFarmsSession) GetFootprintType(_name string) (TraceableFarmsFootprintType, error) {
	return _TraceableFarms.Contract.GetFootprintType(&_TraceableFarms.CallOpts, _name)
}

// GetFootprintType is a free data retrieval call binding the contract method 0x0ba33f01.
//
// Solidity: function getFootprintType(string _name) view returns((string,string,string,string))
func (_TraceableFarms *TraceableFarmsCallerSession) GetFootprintType(_name string) (TraceableFarmsFootprintType, error) {
	return _TraceableFarms.Contract.GetFootprintType(&_TraceableFarms.CallOpts, _name)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraceableFarms *TraceableFarmsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraceableFarms.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraceableFarms *TraceableFarmsSession) Owner() (common.Address, error) {
	return _TraceableFarms.Contract.Owner(&_TraceableFarms.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraceableFarms *TraceableFarmsCallerSession) Owner() (common.Address, error) {
	return _TraceableFarms.Contract.Owner(&_TraceableFarms.CallOpts)
}

// SetAccreditation is a paid mutator transaction binding the contract method 0x87ef27d9.
//
// Solidity: function setAccreditation(string _name, string _description, string _informationalResourceUrl, string _accreditationTypeName) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetAccreditation(opts *bind.TransactOpts, _name string, _description string, _informationalResourceUrl string, _accreditationTypeName string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setAccreditation", _name, _description, _informationalResourceUrl, _accreditationTypeName)
}

// SetAccreditation is a paid mutator transaction binding the contract method 0x87ef27d9.
//
// Solidity: function setAccreditation(string _name, string _description, string _informationalResourceUrl, string _accreditationTypeName) returns()
func (_TraceableFarms *TraceableFarmsSession) SetAccreditation(_name string, _description string, _informationalResourceUrl string, _accreditationTypeName string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetAccreditation(&_TraceableFarms.TransactOpts, _name, _description, _informationalResourceUrl, _accreditationTypeName)
}

// SetAccreditation is a paid mutator transaction binding the contract method 0x87ef27d9.
//
// Solidity: function setAccreditation(string _name, string _description, string _informationalResourceUrl, string _accreditationTypeName) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetAccreditation(_name string, _description string, _informationalResourceUrl string, _accreditationTypeName string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetAccreditation(&_TraceableFarms.TransactOpts, _name, _description, _informationalResourceUrl, _accreditationTypeName)
}

// SetAccreditationType is a paid mutator transaction binding the contract method 0xaa2c57f3.
//
// Solidity: function setAccreditationType(string _name) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetAccreditationType(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setAccreditationType", _name)
}

// SetAccreditationType is a paid mutator transaction binding the contract method 0xaa2c57f3.
//
// Solidity: function setAccreditationType(string _name) returns()
func (_TraceableFarms *TraceableFarmsSession) SetAccreditationType(_name string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetAccreditationType(&_TraceableFarms.TransactOpts, _name)
}

// SetAccreditationType is a paid mutator transaction binding the contract method 0xaa2c57f3.
//
// Solidity: function setAccreditationType(string _name) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetAccreditationType(_name string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetAccreditationType(&_TraceableFarms.TransactOpts, _name)
}

// SetBatch is a paid mutator transaction binding the contract method 0xfe0bb4e0.
//
// Solidity: function setBatch(string _number, string _date, string _productName, string _productVariety, string _productAppearance, string _productSize, string _productDescription, string _productPhotoUrl, string _productStatisticsImageUrl, string _companyNif) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetBatch(opts *bind.TransactOpts, _number string, _date string, _productName string, _productVariety string, _productAppearance string, _productSize string, _productDescription string, _productPhotoUrl string, _productStatisticsImageUrl string, _companyNif string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setBatch", _number, _date, _productName, _productVariety, _productAppearance, _productSize, _productDescription, _productPhotoUrl, _productStatisticsImageUrl, _companyNif)
}

// SetBatch is a paid mutator transaction binding the contract method 0xfe0bb4e0.
//
// Solidity: function setBatch(string _number, string _date, string _productName, string _productVariety, string _productAppearance, string _productSize, string _productDescription, string _productPhotoUrl, string _productStatisticsImageUrl, string _companyNif) returns()
func (_TraceableFarms *TraceableFarmsSession) SetBatch(_number string, _date string, _productName string, _productVariety string, _productAppearance string, _productSize string, _productDescription string, _productPhotoUrl string, _productStatisticsImageUrl string, _companyNif string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatch(&_TraceableFarms.TransactOpts, _number, _date, _productName, _productVariety, _productAppearance, _productSize, _productDescription, _productPhotoUrl, _productStatisticsImageUrl, _companyNif)
}

// SetBatch is a paid mutator transaction binding the contract method 0xfe0bb4e0.
//
// Solidity: function setBatch(string _number, string _date, string _productName, string _productVariety, string _productAppearance, string _productSize, string _productDescription, string _productPhotoUrl, string _productStatisticsImageUrl, string _companyNif) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetBatch(_number string, _date string, _productName string, _productVariety string, _productAppearance string, _productSize string, _productDescription string, _productPhotoUrl string, _productStatisticsImageUrl string, _companyNif string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatch(&_TraceableFarms.TransactOpts, _number, _date, _productName, _productVariety, _productAppearance, _productSize, _productDescription, _productPhotoUrl, _productStatisticsImageUrl, _companyNif)
}

// SetBatchFootprint is a paid mutator transaction binding the contract method 0x7abff3e7.
//
// Solidity: function setBatchFootprint(string _companyNif, string _batchNumber, string _processName, string _footprintTypeName, uint256 _totalValue, string _checkerHash, string _checkerUrl) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetBatchFootprint(opts *bind.TransactOpts, _companyNif string, _batchNumber string, _processName string, _footprintTypeName string, _totalValue *big.Int, _checkerHash string, _checkerUrl string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setBatchFootprint", _companyNif, _batchNumber, _processName, _footprintTypeName, _totalValue, _checkerHash, _checkerUrl)
}

// SetBatchFootprint is a paid mutator transaction binding the contract method 0x7abff3e7.
//
// Solidity: function setBatchFootprint(string _companyNif, string _batchNumber, string _processName, string _footprintTypeName, uint256 _totalValue, string _checkerHash, string _checkerUrl) returns()
func (_TraceableFarms *TraceableFarmsSession) SetBatchFootprint(_companyNif string, _batchNumber string, _processName string, _footprintTypeName string, _totalValue *big.Int, _checkerHash string, _checkerUrl string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatchFootprint(&_TraceableFarms.TransactOpts, _companyNif, _batchNumber, _processName, _footprintTypeName, _totalValue, _checkerHash, _checkerUrl)
}

// SetBatchFootprint is a paid mutator transaction binding the contract method 0x7abff3e7.
//
// Solidity: function setBatchFootprint(string _companyNif, string _batchNumber, string _processName, string _footprintTypeName, uint256 _totalValue, string _checkerHash, string _checkerUrl) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetBatchFootprint(_companyNif string, _batchNumber string, _processName string, _footprintTypeName string, _totalValue *big.Int, _checkerHash string, _checkerUrl string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatchFootprint(&_TraceableFarms.TransactOpts, _companyNif, _batchNumber, _processName, _footprintTypeName, _totalValue, _checkerHash, _checkerUrl)
}

// SetBatchFootprintValue is a paid mutator transaction binding the contract method 0x82339705.
//
// Solidity: function setBatchFootprintValue(string _companyNif, string _batchNumber, string _processName, string _footprintTypeName, string _description, uint256 _value, string _checkerHash, string _checkerUrl) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetBatchFootprintValue(opts *bind.TransactOpts, _companyNif string, _batchNumber string, _processName string, _footprintTypeName string, _description string, _value *big.Int, _checkerHash string, _checkerUrl string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setBatchFootprintValue", _companyNif, _batchNumber, _processName, _footprintTypeName, _description, _value, _checkerHash, _checkerUrl)
}

// SetBatchFootprintValue is a paid mutator transaction binding the contract method 0x82339705.
//
// Solidity: function setBatchFootprintValue(string _companyNif, string _batchNumber, string _processName, string _footprintTypeName, string _description, uint256 _value, string _checkerHash, string _checkerUrl) returns()
func (_TraceableFarms *TraceableFarmsSession) SetBatchFootprintValue(_companyNif string, _batchNumber string, _processName string, _footprintTypeName string, _description string, _value *big.Int, _checkerHash string, _checkerUrl string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatchFootprintValue(&_TraceableFarms.TransactOpts, _companyNif, _batchNumber, _processName, _footprintTypeName, _description, _value, _checkerHash, _checkerUrl)
}

// SetBatchFootprintValue is a paid mutator transaction binding the contract method 0x82339705.
//
// Solidity: function setBatchFootprintValue(string _companyNif, string _batchNumber, string _processName, string _footprintTypeName, string _description, uint256 _value, string _checkerHash, string _checkerUrl) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetBatchFootprintValue(_companyNif string, _batchNumber string, _processName string, _footprintTypeName string, _description string, _value *big.Int, _checkerHash string, _checkerUrl string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatchFootprintValue(&_TraceableFarms.TransactOpts, _companyNif, _batchNumber, _processName, _footprintTypeName, _description, _value, _checkerHash, _checkerUrl)
}

// SetBatchOrigin is a paid mutator transaction binding the contract method 0x1c479fc5.
//
// Solidity: function setBatchOrigin(string _companyNif, string _batchNumber, string _description, string _location, string _locationCoordinates) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetBatchOrigin(opts *bind.TransactOpts, _companyNif string, _batchNumber string, _description string, _location string, _locationCoordinates string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setBatchOrigin", _companyNif, _batchNumber, _description, _location, _locationCoordinates)
}

// SetBatchOrigin is a paid mutator transaction binding the contract method 0x1c479fc5.
//
// Solidity: function setBatchOrigin(string _companyNif, string _batchNumber, string _description, string _location, string _locationCoordinates) returns()
func (_TraceableFarms *TraceableFarmsSession) SetBatchOrigin(_companyNif string, _batchNumber string, _description string, _location string, _locationCoordinates string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatchOrigin(&_TraceableFarms.TransactOpts, _companyNif, _batchNumber, _description, _location, _locationCoordinates)
}

// SetBatchOrigin is a paid mutator transaction binding the contract method 0x1c479fc5.
//
// Solidity: function setBatchOrigin(string _companyNif, string _batchNumber, string _description, string _location, string _locationCoordinates) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetBatchOrigin(_companyNif string, _batchNumber string, _description string, _location string, _locationCoordinates string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatchOrigin(&_TraceableFarms.TransactOpts, _companyNif, _batchNumber, _description, _location, _locationCoordinates)
}

// SetBatchProcess is a paid mutator transaction binding the contract method 0x6fefd4fe.
//
// Solidity: function setBatchProcess(string _companyNif, string _batchNumber, string _name, string _description) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetBatchProcess(opts *bind.TransactOpts, _companyNif string, _batchNumber string, _name string, _description string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setBatchProcess", _companyNif, _batchNumber, _name, _description)
}

// SetBatchProcess is a paid mutator transaction binding the contract method 0x6fefd4fe.
//
// Solidity: function setBatchProcess(string _companyNif, string _batchNumber, string _name, string _description) returns()
func (_TraceableFarms *TraceableFarmsSession) SetBatchProcess(_companyNif string, _batchNumber string, _name string, _description string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatchProcess(&_TraceableFarms.TransactOpts, _companyNif, _batchNumber, _name, _description)
}

// SetBatchProcess is a paid mutator transaction binding the contract method 0x6fefd4fe.
//
// Solidity: function setBatchProcess(string _companyNif, string _batchNumber, string _name, string _description) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetBatchProcess(_companyNif string, _batchNumber string, _name string, _description string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetBatchProcess(&_TraceableFarms.TransactOpts, _companyNif, _batchNumber, _name, _description)
}

// SetCompany is a paid mutator transaction binding the contract method 0xaa36ec5f.
//
// Solidity: function setCompany(string _nif, string _bussinessName, string _description, string _location, string _locationCoordinates, string _informationalResourceUrl, string _consortiumName) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetCompany(opts *bind.TransactOpts, _nif string, _bussinessName string, _description string, _location string, _locationCoordinates string, _informationalResourceUrl string, _consortiumName string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setCompany", _nif, _bussinessName, _description, _location, _locationCoordinates, _informationalResourceUrl, _consortiumName)
}

// SetCompany is a paid mutator transaction binding the contract method 0xaa36ec5f.
//
// Solidity: function setCompany(string _nif, string _bussinessName, string _description, string _location, string _locationCoordinates, string _informationalResourceUrl, string _consortiumName) returns()
func (_TraceableFarms *TraceableFarmsSession) SetCompany(_nif string, _bussinessName string, _description string, _location string, _locationCoordinates string, _informationalResourceUrl string, _consortiumName string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetCompany(&_TraceableFarms.TransactOpts, _nif, _bussinessName, _description, _location, _locationCoordinates, _informationalResourceUrl, _consortiumName)
}

// SetCompany is a paid mutator transaction binding the contract method 0xaa36ec5f.
//
// Solidity: function setCompany(string _nif, string _bussinessName, string _description, string _location, string _locationCoordinates, string _informationalResourceUrl, string _consortiumName) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetCompany(_nif string, _bussinessName string, _description string, _location string, _locationCoordinates string, _informationalResourceUrl string, _consortiumName string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetCompany(&_TraceableFarms.TransactOpts, _nif, _bussinessName, _description, _location, _locationCoordinates, _informationalResourceUrl, _consortiumName)
}

// SetCompanyAccreditation is a paid mutator transaction binding the contract method 0x7eb270a4.
//
// Solidity: function setCompanyAccreditation(string _nif, string _accreditationName, string _checkerHash, string _checkerUrl) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetCompanyAccreditation(opts *bind.TransactOpts, _nif string, _accreditationName string, _checkerHash string, _checkerUrl string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setCompanyAccreditation", _nif, _accreditationName, _checkerHash, _checkerUrl)
}

// SetCompanyAccreditation is a paid mutator transaction binding the contract method 0x7eb270a4.
//
// Solidity: function setCompanyAccreditation(string _nif, string _accreditationName, string _checkerHash, string _checkerUrl) returns()
func (_TraceableFarms *TraceableFarmsSession) SetCompanyAccreditation(_nif string, _accreditationName string, _checkerHash string, _checkerUrl string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetCompanyAccreditation(&_TraceableFarms.TransactOpts, _nif, _accreditationName, _checkerHash, _checkerUrl)
}

// SetCompanyAccreditation is a paid mutator transaction binding the contract method 0x7eb270a4.
//
// Solidity: function setCompanyAccreditation(string _nif, string _accreditationName, string _checkerHash, string _checkerUrl) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetCompanyAccreditation(_nif string, _accreditationName string, _checkerHash string, _checkerUrl string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetCompanyAccreditation(&_TraceableFarms.TransactOpts, _nif, _accreditationName, _checkerHash, _checkerUrl)
}

// SetCompanyFootprintReportability is a paid mutator transaction binding the contract method 0x8e46b1f8.
//
// Solidity: function setCompanyFootprintReportability(string _nif, string _footprintTypeName, string _footprintReportTypeName, string _description) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetCompanyFootprintReportability(opts *bind.TransactOpts, _nif string, _footprintTypeName string, _footprintReportTypeName string, _description string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setCompanyFootprintReportability", _nif, _footprintTypeName, _footprintReportTypeName, _description)
}

// SetCompanyFootprintReportability is a paid mutator transaction binding the contract method 0x8e46b1f8.
//
// Solidity: function setCompanyFootprintReportability(string _nif, string _footprintTypeName, string _footprintReportTypeName, string _description) returns()
func (_TraceableFarms *TraceableFarmsSession) SetCompanyFootprintReportability(_nif string, _footprintTypeName string, _footprintReportTypeName string, _description string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetCompanyFootprintReportability(&_TraceableFarms.TransactOpts, _nif, _footprintTypeName, _footprintReportTypeName, _description)
}

// SetCompanyFootprintReportability is a paid mutator transaction binding the contract method 0x8e46b1f8.
//
// Solidity: function setCompanyFootprintReportability(string _nif, string _footprintTypeName, string _footprintReportTypeName, string _description) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetCompanyFootprintReportability(_nif string, _footprintTypeName string, _footprintReportTypeName string, _description string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetCompanyFootprintReportability(&_TraceableFarms.TransactOpts, _nif, _footprintTypeName, _footprintReportTypeName, _description)
}

// SetConsortium is a paid mutator transaction binding the contract method 0x2fd47f2b.
//
// Solidity: function setConsortium(string _name) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetConsortium(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setConsortium", _name)
}

// SetConsortium is a paid mutator transaction binding the contract method 0x2fd47f2b.
//
// Solidity: function setConsortium(string _name) returns()
func (_TraceableFarms *TraceableFarmsSession) SetConsortium(_name string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetConsortium(&_TraceableFarms.TransactOpts, _name)
}

// SetConsortium is a paid mutator transaction binding the contract method 0x2fd47f2b.
//
// Solidity: function setConsortium(string _name) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetConsortium(_name string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetConsortium(&_TraceableFarms.TransactOpts, _name)
}

// SetFootprintReportType is a paid mutator transaction binding the contract method 0x4e902e8c.
//
// Solidity: function setFootprintReportType(string _name) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetFootprintReportType(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setFootprintReportType", _name)
}

// SetFootprintReportType is a paid mutator transaction binding the contract method 0x4e902e8c.
//
// Solidity: function setFootprintReportType(string _name) returns()
func (_TraceableFarms *TraceableFarmsSession) SetFootprintReportType(_name string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetFootprintReportType(&_TraceableFarms.TransactOpts, _name)
}

// SetFootprintReportType is a paid mutator transaction binding the contract method 0x4e902e8c.
//
// Solidity: function setFootprintReportType(string _name) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetFootprintReportType(_name string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetFootprintReportType(&_TraceableFarms.TransactOpts, _name)
}

// SetFootprintType is a paid mutator transaction binding the contract method 0xfd2aeace.
//
// Solidity: function setFootprintType(string _name, string _description, string _unitMeasurementName, string _unitMeasurementSymbol) returns()
func (_TraceableFarms *TraceableFarmsTransactor) SetFootprintType(opts *bind.TransactOpts, _name string, _description string, _unitMeasurementName string, _unitMeasurementSymbol string) (*types.Transaction, error) {
	return _TraceableFarms.contract.Transact(opts, "setFootprintType", _name, _description, _unitMeasurementName, _unitMeasurementSymbol)
}

// SetFootprintType is a paid mutator transaction binding the contract method 0xfd2aeace.
//
// Solidity: function setFootprintType(string _name, string _description, string _unitMeasurementName, string _unitMeasurementSymbol) returns()
func (_TraceableFarms *TraceableFarmsSession) SetFootprintType(_name string, _description string, _unitMeasurementName string, _unitMeasurementSymbol string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetFootprintType(&_TraceableFarms.TransactOpts, _name, _description, _unitMeasurementName, _unitMeasurementSymbol)
}

// SetFootprintType is a paid mutator transaction binding the contract method 0xfd2aeace.
//
// Solidity: function setFootprintType(string _name, string _description, string _unitMeasurementName, string _unitMeasurementSymbol) returns()
func (_TraceableFarms *TraceableFarmsTransactorSession) SetFootprintType(_name string, _description string, _unitMeasurementName string, _unitMeasurementSymbol string) (*types.Transaction, error) {
	return _TraceableFarms.Contract.SetFootprintType(&_TraceableFarms.TransactOpts, _name, _description, _unitMeasurementName, _unitMeasurementSymbol)
}
