package barobill

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	xj "github.com/basgys/goxml2json"
	"github.com/secr3t/req/v3"
	"github.com/tidwall/gjson"
)

const (
	bankAccountEndpoint        = "BANKACCOUNT.asmx"
	getPeriodBankAccountLogEx2 = "GetPeriodBankAccountLogEx2"
)

type GetPeriodBankAccountLogEx2 struct {
	XMLName        xml.Name `xml:"http://ws.baroservice.com/ GetPeriodBankAccountLogEx2"`
	CertKey        string   `xml:"CERTKEY"`
	CorpNum        string   `xml:"CorpNum"`
	ID             string   `xml:"ID"`
	BankAccountNum string   `xml:"BankAccountNum"`
	StartDate      string   `xml:"StartDate"`
	EndDate        string   `xml:"EndDate"`
	TransDirection int      `xml:"TransDirection"` //1: all, 2: income, 3: outcome
	CountPerPage   int      `xml:"CountPerPage"`
	CurrentPage    int      `xml:"CurrentPage"`
	OrderDirection int      `xml:"OrderDirection"` //1: ASC, 2: DESC
}

type GetPeriodBankAccountLogEx2Response struct {
	GetPeriodBankAccountLogEx2Result GetPeriodBankAccountLogEx2Result `xml:"GetPeriodBankAccountLogEx2Result"`
}
type GetPeriodBankAccountLogEx2Result struct {
	CurrentPage        string `xml:"CurrentPage"`
	CountPerPage       string `xml:"CountPerPage"`
	MaxPageNum         string `xml:"MaxPageNum"`
	MaxIndex           string `xml:"MaxIndex"`
	BankAccountLogList struct {
		BankAccountLogEx2 []BankAccountLogEx2 `xml:"BankAccountLogEx2"`
	} `xml:"BankAccountLogList"`
}

type BankAccountLogEx2 struct {
	CorpNum        string `xml:"CorpNum"`
	BankAccountNum string `xml:"BankAccountNum"`
	TransDirection string `xml:"TransDirection"`
	Deposit        string `xml:"Deposit"`
	Withdrawal     string `xml:"Withdrawal"`
	Balance        string `xml:"Balance"`
	TransDT        string `xml:"TransDT"`
	TransType      string `xml:"TransType"`
	TransOffice    string `xml:"TransOffice"`
	TransRemark    string `xml:"TransRemark"`
	MgRemark1      string `xml:"MgRemark1"`
	MgRemark2      string `xml:"MgRemark2"`
	TransRefKey    string `xml:"TransRefKey"`
	Memo           string `xml:"Memo"`
}

func NewGetPeriodBankAccountLogEx2(bankAccountNum, startDate, endDate string, countPerPage, currentPage, transDirection, orderDirection int) GetPeriodBankAccountLogEx2 {
	return GetPeriodBankAccountLogEx2{
		BankAccountNum: bankAccountNum,
		StartDate:      startDate,
		EndDate:        endDate,
		TransDirection: transDirection,
		CountPerPage:   countPerPage,
		CurrentPage:    currentPage,
		OrderDirection: orderDirection,
	}
}

func (c *Client) GetPeriodBankAccountLogEx2(bal GetPeriodBankAccountLogEx2) (res GetPeriodBankAccountLogEx2Result, err error) {
	bal.CertKey = c.certKey
	bal.CorpNum = c.corpNum
	bal.ID = c.barobillUsername
	var (
		resp     *req.Response
		envelope = NewEnvelope(bal)
	)
	envelope.Body.Namespace = c.namespace

	resp, err = c.rc.R().
		SetBodyXmlMarshal(envelope).
		SetHeader("Content-Type", contentTypeSoapXml).
		Post(bankAccountEndpoint)
	if resp != nil && resp.IsErrorState() && err == nil {
		err = errors.New(resp.String())
		return
	}

	if resp != nil && resp.Body != nil {
		jsonBuffer, _ := xj.Convert(resp.Body)
		json.Unmarshal([]byte(gjson.Parse(jsonBuffer.String()).Get("Envelope.Body.GetPeriodBankAccountLogEx2Response.GetPeriodBankAccountLogEx2Result").Raw), &res)
	} else {
		err = errors.New("internal server error from barobill")
	}
	return
}
