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

type InformationOutput struct {
	RegisterNo string `json:"registerNo"`
	BranchNo   string `json:"branchNo"`
	PosID      string `json:"posId"`
	DBDirPath  string `json:"dbDirPath"`
	ExtraInfo  struct {
		CountBill    string `json:"countBill"`
		CountLottery int    `json:"countLottery"`
		LastSentDate string `json:"lastSentDate"`
		PosVersion   string `json:"posVersion"`
	} `json:"extraInfo"`
}

type PutOutput struct {
	Amount           string             `json:"amount"`
	Vat              string             `json:"vat"`
	CashAmount       string             `json:"cashAmount"`
	NonCashAmount    string             `json:"nonCashAmount"`
	CityTax          string             `json:"cityTax"`
	DistrictCode     string             `json:"districtCode"`
	PosNo            string             `json:"dosNo"`
	CustomerNo       string             `json:"customerNo"`
	BillType         string             `json:"billType"`
	BillIdSuffix     string             `json:"billIdSuffix"`
	ReturnBillId     string             `json:"returnBillId"`
	TaxType          string             `json:"taxType"`
	InvoiceID        string             `json:"invoiceId"`
	ReportMonth      string             `json:"reportMonth"`
	BranchNo         string             `json:"branchNo"`
	Stocks           []Stocks           `json:"stocks"`
	BankTransactions []bankTransactions `json:"bankTransactions"`
}

type Barimt struct {
	Amount           string            `json:"amount"`
	Vat              string            `json:"vat"`
	CashAmount       string            `json:"cashAmount"`
	NonCashAmount    string            `json:"nonCashAmount"`
	CityTax          string            `json:"cityTax"`
	DistrictCode     string            `json:"districtCode"`
	PosNo            string            `json:"posNo"`
	CustomerNo       string            `json:"customerNo"`
	BillType         string            `json:"billType"`
	BillIDSuffix     string            `json:"billIdSuffix"`
	ReturnBillID     string            `json:"returnBillId"`
	TaxType          string            `json:"taxType"`
	InvoiceID        string            `json:"invoiceId"`
	ReportMonth      string            `json:"reportMonth"`
	BranchNo         string            `json:"branchNo"`
	Stocks           []Stock           `json:"stocks"`
	BankTransactions []BankTransaction `json:"bankTransactions"`
}

type Stock struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	MeasureUnit string `json:"measureUnit"`
	Qty         string `json:"qty"`
	UnitPrice   string `json:"unitPrice"`
	TotalAmount string `json:"totalAmount"`
	CityTax     string `json:"cityTax"`
	Vat         string `json:"vat"`
	BarCode     string `json:"barCode"`
}

type BankTransaction struct {
	Rrn          string `json:"rrn"`
	BankID       string `json:"bankId"`
	TerminalID   string `json:"terminalId"`
	ApprovalCode string `json:"approvalCode"`
	Amount       string `json:"amount"`
}

type Stocks struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	MeasureUnit string `json:"measureUnit"`
	Qty         string `json:"qty"`
	UnitPrice   string `json:"unitPrice"`
	TotalAmount string `json:"totalAmount"`
	CityTax     string `json:"cityTax"`
	Vat         string `json:"vat"`
	BarCode     string `json:"barCode"`
}

type bankTransactions struct {
	Rrn          string `json:"rrn"`
	BankID       string `json:"bankId"`
	TerminalID   string `json:"terminalId"`
	ApprovalCode string `json:"approvalCode"`
	Amount       string `json:"amount"`
}

type BillInput struct {
	ReturnBillID string `json:"returnBillId"`
	Date         string `json:"date"`
}

type BillOutput struct {
	Success   bool   `json:"success"`
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}

type DataOutput struct {
	Success   bool   `json:"success"`
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}
