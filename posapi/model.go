package posapi

type CheckOutput struct {
	Success  bool `json:"success"`
	Database struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"database"`
	Config struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"config"`
	Network struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"network"`
}

type InformationResponse struct {
	OperatorName  string `json:"operatorName"`
	OperatorTIN   string `json:"operatorTIN"`
	PosId         int    `json:"posId"`
	PosNo         string `json:"posNo"`
	Version       string `json:"version"`
	LastSentDate  string `json:"lastSentDate"`
	LeftLotteries int    `json:"leftLotteries"`
	AppInfo       struct {
		ApplicationDir     string   `json:"applicationDir"`
		CurrentDir         string   `json:"currentDir"`
		Database           string   `json:"database"`
		DatabaseHost       string   `json:"database-host"`
		SupportedDatabases []string `json:"supported-databases"`
		WorkDir            string   `json:"workDir"`
	} `json:"appInfo"`
	PaymentTypes []struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"paymentTypes"`
	Merchants []Merchant `json:"merchants"`
}
type Merchant struct {
	TIN       string `json:"tin"`
	Name      string `json:"name"`
	VatPayer  bool   `json:"vatPayer"`
	Customers []struct {
		TIN      string `json:"tin"`
		Name     string `json:"name"`
		VatPayer bool   `json:"vatPayer"`
	} `json:"customers"`
}
type ReceiptData struct {
	TotalAmount  string      `json:"totalAmount"`
	TotalVAT     string      `json:"totalVAT"`
	TotalCityTax string      `json:"totalCityTax"`
	DistrictCode string      `json:"districtCode"`
	MerchantTin  string      `json:"merchantTin"`
	PosNo        string      `json:"posNo"`
	CustomerTin  string      `json:"customerTin"`
	CustomerNo   string      `json:"customerNo"`
	BranchNo     string      `json:"branchNo"`
	Type         string      `json:"type"`
	InActiveID   string      `json:"inactiveId"`
	InvoiceID    string      `json:"invoiceId"`
	ReportMonth  *string     `json:"reportMonth"`
	Data         interface{} `json:"data"`
	Receipts     []Receipt   `json:"receipts"`
	Payments     []Payment   `json:"payments"`
}
type Receipt struct {
	TotalAmount   string      `json:"totalAmount"`
	TotalVAT      string      `json:"totalVAT"`
	TotalCityTax  string      `json:"totalCityTax"`
	TaxType       string      `json:"taxType"`
	MerchantTin   string      `json:"merchantTin"`
	BankAccountNo string      `json:"bankAccountNo"`
	Data          interface{} `json:"data"`
	Items         []Item      `json:"items"`
}

type Payment struct {
	Code         string      `json:"code"`
	ExchangeCode string      `json:"exchangeCode"`
	PaidAmount   string      `json:"paidAmount"`
	Data         interface{} `json:"Data"`
	Status       string      `json:"Status"`
}

type Item struct {
	Name               string      `json:"name"`
	BarCode            string      `json:"barCode"`
	BarCodeType        string      `json:"barCodeType"`
	ClassificationCode string      `json:"classificationCode"`
	TaxProductCode     string      `json:"taxProductCode"`
	MeasureUnit        string      `json:"measureUnit"`
	Qty                string      `json:"qty"`
	UnitPrice          string      `json:"unitPrice"`
	TotalVAT           string      `json:"totalVAT"`
	TotalCityTax       string      `json:"totalCityTax"`
	TotalAmount        string      `json:"totalAmount"`
	Data               interface{} `json:"data"`
}

type EBarimtReceipt struct {
	ID            string `json:"id"`
	BankAccountId string `json:"bankAccountId"`
}

type BillInput struct {
	ReturnBillID string `json:"id"`
	Date         string `json:"date"`
}

type BillOutput struct {
	Success   bool   `json:"success"`
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}
type EBarimtResponse struct {
	ID           string  `json:"id"`
	Version      string  `json:"version"`
	TotalAmount  float32 `json:"totalAmount"`
	TotalVAT     float32 `json:"totalVAT"`
	TotalCityTax float32 `json:"totalCityTax"`
	BranchNo     string  `json:"branchNo"`
	DistrictCode string  `json:"districtCode"`
	MerchantTin  string  `json:"merchantTin"`
	PosNo        string  `json:"posNo"`
	CustomerTin  string  `json:"customerTin"`
	Type         string  `json:"type"`
	InvoiceID    string  `json:"invoiceId"`
	ReportMonth  string  `json:"reportMonth"`
	Receipts     []struct {
		ID          string  `json:"id"`
		TotalAmount float32 `json:"totalAmount"`
		TaxType     string  `json:"taxType"`
		Items       []struct {
			Name               string  `json:"name"`
			BarCode            string  `json:"barCode"`
			BarCodeType        string  `json:"barCodeType"`
			ClassificationCode string  `json:"classificationCode"`
			TaxProductCode     string  `json:"taxProductCode"`
			MeasureUnit        string  `json:"measureUnit"`
			Qty                int     `json:"qty"`
			UnitPrice          float32 `json:"unitPrice"`
			TotalAmount        float32 `json:"totalAmount"`
			TotalVAT           float32 `json:"totalVAT"`
			TotalCityTax       float32 `json:"totalCityTax"`
		} `json:"items"`
		MerchantTin  string  `json:"merchantTin"`
		TotalVAT     float32 `json:"totalVAT"`
		TotalCityTax float32 `json:"totalCityTax"`
	} `json:"receipts"`
	Payments []struct {
		Code         string  `json:"code"`
		ExchangeCode string  `json:"exchangeCode"`
		PaidAmount   float32 `json:"paidAmount"`
		Status       string  `json:"status"`
	} `json:"payments"`
	PosID   int    `json:"posId"`
	Status  string `json:"status"`
	Message string `json:"message"`
	QrData  string `json:"qrData"`
	Lottery string `json:"lottery"`
	Date    string `json:"date"`
	Easy    bool   `json:"easy"`
}
