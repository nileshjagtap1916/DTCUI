package main

type Contract struct {
	ContractId              string `json:"CONTRACT_ID"`
	OrderId                 string `json:"ORDER_INFO"`
	PaymentCommitment       bool   `json:"PAYMENT_COMMITMENT"`
	PaymentConfirmation     bool   `json:"PAYMENT_CONFIRMATION"`
	InformationCounterparty bool   `json:"INFORMATION_COUNTERPARTY"`
	ForfeitingInvoice       bool   `json:"FORFEITING_INVOICE"`
	ContractCreateDate      string `json:"CONTRACT_CREATE_DATE"`
	PaymentDueDate          string `json:"PAYMENT_DUE_DATE"`
	InvoiceStatus           string `json:"INVOICE_STATUS"`
	PaymentStatus           string `json:"PAYMENT_STATUS"`
	ContractStatus          string `json:"CONTRACT_STATUS"`
	DeliveryStatus          string `json:"DELIVERY_STATUS"`
}

type Order struct {
	OrderId     string    `json:"ORDER_ID"`
	ArticleList []Article `json:"ARTICALE_LIST"`
	BuyerId     string    `json:"BUYER_INFO"`
	SellerId    string    `json:"SELLER_INFO"`
	ShipmentId  string    `json:"SHIPMENT_INFO"`
	TotalAmount string    `json:"TOTAL_AMOUNT"`
}

type OutputContract struct {
	ContractId              string      `json:"CONTRACT_ID"`
	OrderDetails            OutputOrder `json:"ORDER_INFO"`
	PaymentCommitment       bool        `json:"PAYMENT_COMMITMENT"`
	PaymentConfirmation     bool        `json:"PAYMENT_CONFIRMATION"`
	InformationCounterparty bool        `json:"INFORMATION_COUNTERPARTY"`
	ForfeitingInvoice       bool        `json:"FORFEITING_INVOICE"`
	ContractCreateDate      string      `json:"CONTRACT_CREATE_DATE"`
	PaymentDueDate          string      `json:"PAYMENT_DUE_DATE"`
	InvoiceStatus           string      `json:"INVOICE_STATUS"`
	PaymentStatus           string      `json:"PAYMENT_STATUS"`
	ContractStatus          string      `json:"CONTRACT_STATUS"`
	DeliveryStatus          string      `json:"DELIVERY_STATUS"`
}

type OutputOrder struct {
	OrderId         string    `json:"ORDER_ID"`
	ArticleList     []Article `json:"ARTICALE_LIST"`
	BuyerDetails    User      `json:"BUYER_INFO"`
	SellerDetails   User      `json:"SELLER_INFO"`
	ShipmentDetails Shipment  `json:"SHIPMENT_INFO"`
	TotalAmount     string    `json:"TOTAL_AMOUNT"`
}

type User struct {
	UserId      string `json:"USER_ID"`
	UserName    string `json:"USER_NAME"`
	UserBank    string `json:"USER_BANK"`
	UserAddress string `json:"USER_ADDRESS"`
	UserType    string `json:"USER_TYPE"`
}

type Shipment struct {
	ShipmentId      string `json:"SHIPMENT_ID"`
	ShipmentCompany string `json:"SHIPMENT_COMPANY"`
	ShipmentStatus  string `json:"SHIPMENT_STATUS"`
}

type Article struct {
	ArticleDescription string `json:"ARTICLE_DESC"`
	ArticleQty         string `json:"ARTICLE_QTY"`
	ArticlePrice       string `json:"ARTICLE_PRICE"`
}

/*type Company struct {
	CompanyName string `json:"name"`
}

type Address struct {
	Line    string `json:"line"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode string `json:"pincode"`
}
type Bank struct {
	BankName   string `json:"name"`
	BranchName string `json:"branch"`
	Country    string `json:"country"`
	Currency   string `json:"currency"`
}
type Amount struct {
	Currency string `json:"currency"`
	Value    uint64 `json:"value"`
}*/
