// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;
pragma experimental ABIEncoderV2;

contract TraceableFarms {

    // Dirección del propietario del contrato (quien lo publica)
    address public owner;

    // Constructor
    constructor() {
        owner = msg.sender;
    }

    // Estructura de consorcio de productores agrícolas
    struct Consortium {
        string name;
    }

    // Estructura de una empresa que forma parte de un consorcio de productores agrícolas
    struct Company {
        string nif;
        string bussinessName;
        string description;
        string location;
        string locationCoordinates;
        string informationalResourceUrl;
        Consortium consortium;
    }

    // Estructura de tipo de acreditación
    struct AccreditationType {
        string name;
    }

    // Estructura de acreditación
    struct Accreditation {
        string name;
        string description;
        string informationalResourceUrl;
        AccreditationType accreditationType;
    }

    // Estructura de acreditación de una empresa
    struct CompanyAccreditation {
        Accreditation accreditation;
        string checkerHash;
        string checkerUrl;
    }

    // Estructura de tipo de huella
    struct FootprintType {
        string name;
        string description;
        string unitMeasurementName;
        string unitMeasurementSymbol;
    }

    // Estructura de tipo de reporte de huella (reportabilidad de la huella)
    struct FootprintReportType {
        string name;
    }

    // Estructura de reportabilidad de huella de una empresa
    struct CompanyFootprintReportability {
        Company company;
        FootprintType footprintType;
        FootprintReportType footprintReportType;
        string description;
    }

    // Estructura de lote de un producto agrícola
    struct Batch {
        string number;
        string date;
        string productName;
        string productVariety;
        string productAppearance;
        string productSize;
        string productDescription;
        string productPhotoUrl;
        string productStatisticsImageUrl;
        Company company;
    }

    // Estructura de origen de un lote (de un producto agrícola)
    struct BatchOrigin {
        string description;
        string location;
        string locationCoordinates;
    }

    // Estructura de proceso de un lote (de un producto agrícola)
    struct BatchProcess {
        string name;
        string description;
    }

    // Estructura de huella de un proceso (de un lote) (considera valor total y su verificador)
    struct BatchFootprint {
        FootprintType footprintType;
        uint totalValue;
        string checkerHash;
        string checkerUrl;
    }

    // Estructura de valores y verificadores de huella de un proceso (de un lote)
    struct BatchFootprintValue {
        string description;
        uint value;
        string checkerHash;
        string checkerUrl;
    }
    

    // Mapping para persistencia de consorcios
    mapping (bytes32 => Consortium) private consortiums;

    // Mapping para persistencia de empresas
    mapping (bytes32 => Company) private companies;

    // Mapping para persistencia de tipos de acreditación
    mapping (bytes32 => AccreditationType) private accreditationTypes;

    // Mapping para persistencia de acreditaciones
    mapping (bytes32 => Accreditation) private accreditations;

    // Mapping para persistencia de acreditaciones de una empresa
    mapping (bytes32 => CompanyAccreditation[]) private companiesAccreditation;

    // Mapping para persistencia de tipos de huella
    mapping (bytes32 => FootprintType) private footprintTypes;
    
    // Mapping para persistencia de tipos de reporte de huella
    mapping (bytes32 => FootprintReportType) private footprintReportTypes;

    // Mapping para persistencia de reportabilidad de huella de una empresa
    mapping (bytes32 => CompanyFootprintReportability[]) private companiesFootprintReportability;

    // Mapping para persistencia de lotes de producto
    mapping (bytes32 => Batch) private batches;

    // Mapping para persistencia de origenes de un lote
    mapping (bytes32 => BatchOrigin[]) private batchOrigins;

    // Mapping para persistencia de procesos de un lote
    mapping (bytes32 => BatchProcess[]) private batchProcesses;

    // Mapping para persistencia de huellas de un proceso (de un lote)
    mapping (bytes32 => BatchFootprint[]) private batchFootprints;

    // Mapping para persistencia de valores y verificadores de huellas de un proceso (de un lote)
    mapping (bytes32 => BatchFootprintValue[]) private batchFootprintValues;


    function setConsortium(string memory _name) public {
        bytes32 consortiumHash = keccak256(abi.encodePacked(_name));
        consortiums[consortiumHash] = Consortium(_name);
    }

    function getConsortium(string memory _name) public view returns (Consortium memory) {
        bytes32 consortiumHash = keccak256(abi.encodePacked(_name));
        return consortiums[consortiumHash];
    }

    function setCompany(string memory _nif, string memory _bussinessName, string memory _description, string memory _location, string memory _locationCoordinates, string memory _informationalResourceUrl, string memory _consortiumName) public {
        bytes32 companyHash = keccak256(abi.encodePacked(_nif));
        companies[companyHash] = Company(_nif, _bussinessName, _description, _location, _locationCoordinates, _informationalResourceUrl, getConsortium(_consortiumName));
    }

    function getCompany(string memory _nif) public view returns (Company memory) {
        bytes32 companyHash = keccak256(abi.encodePacked(_nif));
        return companies[companyHash];
    }

    function setAccreditationType(string memory _name) public {
        bytes32 accreditationTypeHash = keccak256(abi.encodePacked(_name));
        accreditationTypes[accreditationTypeHash] = AccreditationType(_name);
    }

    function getAccreditationType(string memory _name) public view returns (AccreditationType memory) {
        bytes32 accreditationTypeHash = keccak256(abi.encodePacked(_name));
        return accreditationTypes[accreditationTypeHash];
    }

    function setAccreditation(string memory _name, string memory _description, string memory _informationalResourceUrl, string memory _accreditationTypeName) public {
        bytes32 accreditationHash = keccak256(abi.encodePacked(_name));
        accreditations[accreditationHash] = Accreditation(_name, _description, _informationalResourceUrl, getAccreditationType(_accreditationTypeName));
    }

    function getAccreditation(string memory _name) public view returns (Accreditation memory) {
        bytes32 accreditationHash = keccak256(abi.encodePacked(_name));
        return accreditations[accreditationHash];
    }    

    function setCompanyAccreditation(string memory _nif, string memory _accreditationName, string memory _checkerHash, string memory _checkerUrl) public {
        bytes32 companyHash = keccak256(abi.encodePacked(_nif));
        
        companiesAccreditation[companyHash].push(
            CompanyAccreditation(getAccreditation(_accreditationName), _checkerHash, _checkerUrl)
        );
    }

    function getCompanyAccreditation(string memory _nif) public view returns (CompanyAccreditation[] memory) {
        bytes32 companyHash = keccak256(abi.encodePacked(_nif));
        return companiesAccreditation[companyHash];
    }

    function setFootprintType(string memory _name, string memory _description, string memory _unitMeasurementName, string memory _unitMeasurementSymbol) public {
        bytes32 footprintTypeHash = keccak256(abi.encodePacked(_name));
        footprintTypes[footprintTypeHash] = FootprintType(_name, _description, _unitMeasurementName, _unitMeasurementSymbol);
    }

    function getFootprintType(string memory _name) public view returns (FootprintType memory) {
        bytes32 footprintTypeHash = keccak256(abi.encodePacked(_name));
        return footprintTypes[footprintTypeHash];
    }
    
    function setFootprintReportType(string memory _name) public {
        bytes32 footprintReportTypeHash = keccak256(abi.encodePacked(_name));
        footprintReportTypes[footprintReportTypeHash] = FootprintReportType(_name);
    }

    function getFootprintReportType(string memory _name) public view returns (FootprintReportType memory) {
        bytes32 footprintReportTypeHash = keccak256(abi.encodePacked(_name));
        return footprintReportTypes[footprintReportTypeHash];
    }

    function setCompanyFootprintReportability(string memory _nif, string memory _footprintTypeName, string memory _footprintReportTypeName, string memory _description) public {
        bytes32 companyFootprintReportabilityHash = keccak256(abi.encodePacked(_nif, _footprintTypeName));

        companiesFootprintReportability[companyFootprintReportabilityHash].push(
            CompanyFootprintReportability(getCompany(_nif), getFootprintType(_footprintTypeName), getFootprintReportType(_footprintReportTypeName), _description)
        );
    }

    function getCompanyFootprintReportability(string memory _nif, string memory _footprintTypeName) public view returns (CompanyFootprintReportability[] memory) {
        bytes32 companyFootprintReportabilityHash = keccak256(abi.encodePacked(_nif, _footprintTypeName));

        return companiesFootprintReportability[companyFootprintReportabilityHash];
    }

    function setBatch(string memory _number, string memory _date, string memory _productName, string memory _productVariety, string memory _productAppearance, string memory _productSize, string memory _productDescription, string memory _productPhotoUrl, string memory _productStatisticsImageUrl, string memory _companyNif) public {
        bytes32 batchHash = keccak256(abi.encodePacked(_companyNif, _number));
        batches[batchHash] = Batch(_number, _date, _productName, _productVariety, _productAppearance, _productSize, _productDescription, _productPhotoUrl, _productStatisticsImageUrl, getCompany(_companyNif));
    }

    function getBatch(string memory _companyNif, string memory _number) public view returns (Batch memory) {
        bytes32 batchHash = keccak256(abi.encodePacked(_companyNif, _number));
        return batches[batchHash];
    }

    function setBatchOrigin(string memory _companyNif, string memory _batchNumber, string memory _description, string memory _location, string memory _locationCoordinates) public {
        bytes32 batchHash = keccak256(abi.encodePacked(_companyNif, _batchNumber));

        batchOrigins[batchHash].push(
            BatchOrigin(_description, _location, _locationCoordinates)
        );
    }

    function getBatchOrigin(string memory _companyNif, string memory _batchNumber) public view returns (BatchOrigin[] memory) {
        bytes32 batchHash = keccak256(abi.encodePacked(_companyNif, _batchNumber));
        return batchOrigins[batchHash];
    }

    function setBatchProcess(string memory _companyNif, string memory _batchNumber, string memory _name, string memory _description) public {
        bytes32 batchHash = keccak256(abi.encodePacked(_companyNif, _batchNumber));

        batchProcesses[batchHash].push(
            BatchProcess(_name, _description)
        );
    }

    function getBatchProcess(string memory _companyNif, string memory _batchNumber) public view returns (BatchProcess[] memory) {
        bytes32 batchHash = keccak256(abi.encodePacked(_companyNif, _batchNumber));
        return batchProcesses[batchHash];
    }

    function setBatchFootprint(string memory _companyNif, string memory _batchNumber, string memory _processName, string memory _footprintTypeName, uint _totalValue, string memory _checkerHash, string memory _checkerUrl) public {
        bytes32 batchFootprintHash = keccak256(abi.encodePacked(_companyNif, _batchNumber, _processName));

        batchFootprints[batchFootprintHash].push(
            BatchFootprint(getFootprintType(_footprintTypeName), _totalValue, _checkerHash, _checkerUrl)
        );
    }

    function getBatchFootprint(string memory _companyNif, string memory _batchNumber, string memory _processName) public view returns (BatchFootprint[] memory) {
        bytes32 batchFootprintHash = keccak256(abi.encodePacked(_companyNif, _batchNumber, _processName));
        return batchFootprints[batchFootprintHash];
    }

    function setBatchFootprintValue(string memory _companyNif, string memory _batchNumber, string memory _processName, string memory _footprintTypeName, string memory _description, uint _value, string memory _checkerHash, string memory _checkerUrl) public {
        bytes32 batchFootprintValueHash = keccak256(abi.encodePacked(_companyNif, _batchNumber, _processName, getFootprintType(_footprintTypeName).name));

        batchFootprintValues[batchFootprintValueHash].push(
            BatchFootprintValue(_description, _value, _checkerHash, _checkerUrl)
        );
    }

    function getBatchFootprintValue(string memory _companyNif, string memory _batchNumber, string memory _processName, string memory _footprintTypeName) public view returns (BatchFootprintValue[] memory) {
        bytes32 batchFootprintValueHash = keccak256(abi.encodePacked(_companyNif, _batchNumber, _processName, getFootprintType(_footprintTypeName).name));
        return batchFootprintValues[batchFootprintValueHash];
    }

}