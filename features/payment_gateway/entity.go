package payment_gateway

type InvoiceObj struct {
	Id          string
	Amount      float64
	Name        string
	Email       string
	Description string
	Currency    string
}

type InvoiceData struct {
	Id         string
	InvoiceUrl string
}

type Data interface {
	CreateInvoice(InvoiceObj) (InvoiceData, error)
	GetInvoice(string) (InvoiceData, error)
}
