package barobill

import (
	"encoding/xml"
	"errors"
	"github.com/secr3t/req/v3"
	"time"
)

const (
	issueTaxInvoiceEx        = "IssueTaxInvoiceEx"
	registAndIssueTaxInvoice = "RegistAndIssueTaxInvoice"

	taxInvoiceEndpoint = "TI.asmx"
)

type RegistAndIssueTaxInvoice struct {
	XMLName    xml.Name `xml:"http://ws.baroservice.com/ RegistAndIssueTaxInvoice"`
	CertKey    string   `xml:"CERTKEY"`
	CorpNum    string   `xml:"CorpNum"`
	Invoice    Invoice
	SendSMS    bool   `xml:"SendSMS"`
	ForceIssue bool   `xml:"ForceIssue"`
	MailTitle  string `xml:"MailTitle"`
}

// Invoice struct
type Invoice struct {
	InvoicerParty            Party                    `xml:"InvoicerParty"`
	InvoiceeParty            Party                    `xml:"InvoiceeParty"`
	BrokerParty              Party                    `xml:"BrokerParty"`
	IssueDirection           IssueDirection           `xml:"IssueDirection"`
	TaxInvoiceType           TaxInvoiceType           `xml:"TaxInvoiceType"`
	TaxType                  TaxType                  `xml:"TaxType"`
	PurposeType              PurposeType              `xml:"PurposeType"`
	ModifyCode               ModifyCode               `xml:"ModifyCode"`
	WriteDate                string                   `xml:"WriteDate"`
	AmountTotal              int                      `xml:"AmountTotal"`
	TaxTotal                 int                      `xml:"TaxTotal"`
	TotalAmount              int                      `xml:"TotalAmount"`
	Cash                     int                      `xml:"Cash"`
	ChkBill                  string                   `xml:"ChkBill"`
	Note                     string                   `xml:"Note"`
	Credit                   string                   `xml:"Credit"`
	Remark1                  string                   `xml:"Remark1"`
	Remark2                  string                   `xml:"Remark2"`
	Remark3                  string                   `xml:"Remark3"`
	Kwon                     string                   `xml:"Kwon"`
	Ho                       string                   `xml:"Ho"`
	SerialNum                string                   `xml:"SerialNum"`
	TaxInvoiceTradeLineItems TaxInvoiceTradeLineItems `xml:"TaxInvoiceTradeLineItems"`
}

func NewInvoice(invoicer Party, invoicee Party, taxInvoiceTradeLineItems ...TaxInvoiceTradeLineItem) Invoice {
	var amount, tax int
	today := time.Now().Format("20060102")
	for _, item := range taxInvoiceTradeLineItems {
		amount += item.Amount
		tax += item.Tax
		item.PurchaseExpiry = today
	}
	return Invoice{
		InvoicerParty:  invoicer,
		InvoiceeParty:  invoicee,
		IssueDirection: TypeIssueDirectionNormal,
		TaxInvoiceType: TypeTaxInvoice,
		TaxType:        TypeTax,
		PurposeType:    TypePurposeReceipt,
		WriteDate:      today,
		AmountTotal:    amount,
		TaxTotal:       tax,
		TotalAmount:    amount + tax,
		Cash:           amount + tax,
		TaxInvoiceTradeLineItems: TaxInvoiceTradeLineItems{
			TaxInvoiceTradeLineItem: taxInvoiceTradeLineItems,
		},
	}
}

// Party struct for InvoicerParty, InvoiceeParty, BrokerParty
type Party struct {
	ContactID   string `xml:"ContactID"` // 바로빌 회원아이디
	CorpNum     string `xml:"CorpNum"`
	MgtNum      string `xml:"MgtNum"`
	CorpName    string `xml:"CorpName"`
	TaxRegID    string `xml:"TaxRegID"`
	CEOName     string `xml:"CEOName"`
	Addr        string `xml:"Addr"`
	BizClass    string `xml:"BizClass"`
	BizType     string `xml:"BizType"`
	ContactName string `xml:"ContactName"`
	TEL         string `xml:"TEL"`
	HP          string `xml:"HP"`
	Email       string `xml:"Email"`
}

// TaxInvoiceTradeLineItems struct
type TaxInvoiceTradeLineItems struct {
	TaxInvoiceTradeLineItem []TaxInvoiceTradeLineItem `xml:"TaxInvoiceTradeLineItem"`
}

// TaxInvoiceTradeLineItem struct
type TaxInvoiceTradeLineItem struct {
	PurchaseExpiry string `xml:"PurchaseExpiry"`
	Name           string `xml:"Name"`
	Information    string `xml:"Information"`
	ChargeableUnit string `xml:"ChargeableUnit"`
	UnitPrice      string `xml:"UnitPrice"`
	Amount         int    `xml:"Amount"`
	Tax            int    `xml:"Tax"`
	Description    string `xml:"Description"`
}

// IssueTaxInvoiceEx struct
type IssueTaxInvoiceEx struct {
	XMLName           xml.Name `xml:"http://ws.baroservice.com/ IssueTaxInvoiceEx"`
	CertKey           string   `xml:"CERTKEY"`
	CorpNum           string   `xml:"CorpNum"`
	MgtKey            string   `xml:"MgtKey"`
	SendSMS           bool     `xml:"SendSMS"`
	SMSMessage        string   `xml:"SMSMessage,omitempty"`
	ForceIssue        bool     `xml:"ForceIssue"`
	MailTitle         string   `xml:"MailTitle,omitempty"`
	BusinessLicenseYN bool     `xml:"BusinessLicenseYN"`
	BankBookYN        bool     `xml:"BankBookYN"`
}

type IssueTaxInvoiceExResponse struct {
	XMLName                 xml.Name `xml:"IssueTaxInvoiceExResponse"`
	IssueTaxInvoiceExResult int      `xml:"IssueTaxInvoiceExResult"`
}

type RegistAndIssueTaxInvoiceResponse struct {
	XMLName                        xml.Name `xml:"RegistAndIssueTaxInvoiceResponse"`
	RegistAndIssueTaxInvoiceResult int      `xml:"RegistAndIssueTaxInvoiceResult"`
}

func (r RegistAndIssueTaxInvoiceResponse) ErrorString() string {
	return errorCodes[r.RegistAndIssueTaxInvoiceResult]
}

func (c *Client) IssueTaxInvoiceEx(iTi IssueTaxInvoiceEx) (res IssueTaxInvoiceExResponse, err error) {
	iTi.CertKey = c.certKey
	iTi.CorpNum = c.corpNum
	var (
		resp        *req.Response
		resEnvelope Envelope[IssueTaxInvoiceExResponse]
		envelope    = NewEnvelope(iTi)
	)
	envelope.Body.Namespace = c.namespace

	resp, err = c.rc.R().
		SetBodyXmlMarshal(envelope).
		SetHeader("Content-Type", contentTypeSoapXml).
		SetSuccessResult(&resEnvelope).Post(taxInvoiceEndpoint)
	if resp != nil && resp.IsErrorState() && err == nil {
		err = errors.New(resp.String())
	}
	res = resEnvelope.Body.Body
	return
}

func (c *Client) RegistAndIssueTaxInvoice(invoice RegistAndIssueTaxInvoice) (res RegistAndIssueTaxInvoiceResponse, err error) {
	invoice.CertKey = c.certKey
	invoice.CorpNum = c.corpNum
	var (
		resp        *req.Response
		resEnvelope ResEnvelope[RegistAndIssueTaxInvoiceResponse]
		envelope    = NewEnvelope(invoice)
	)
	envelope.Body.Namespace = c.namespace

	resp, err = c.rc.R().
		SetBodyXmlMarshal(envelope).
		SetHeader("Content-Type", contentTypeSoapXml).
		SetSuccessResult(&resEnvelope).Post(taxInvoiceEndpoint)
	if resp != nil && resp.IsErrorState() && err == nil {
		err = errors.New(resp.String())
	}
	res = resEnvelope.Body.Body
	return
}
