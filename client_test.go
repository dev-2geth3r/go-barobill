package barobill

import (
	"os"
	"testing"
)

func TestGetBankAccountLogs(t *testing.T) {
	c := NewClient(os.Getenv(envCertKey), os.Getenv(envCorpNum), "secr3t").ProductionGateway()
	res, err := c.GetPeriodBankAccountLogEx2(NewGetPeriodBankAccountLogEx2("54690101226236", "20250514", "20250514", 100, 1, 2, 2))
	t.Log(res, err)
}

func TestIssueTaxInvoiceEx(t *testing.T) {
	c := NewClient(os.Getenv(envCertKey), os.Getenv(envCorpNum))
	res, err := c.IssueTaxInvoiceEx(IssueTaxInvoiceEx{
		MgtKey: "teest",
	})

	t.Log(res, err)
}

func TestRegistAndIssueTaxInvoice(t *testing.T) {
	c := NewClient(os.Getenv(envCertKey), os.Getenv(envCorpNum))
	invoice := NewInvoice(
		Party{
			MgtNum:      NewTSID(),
			CorpNum:     os.Getenv(envCorpNum),
			CorpName:    "정발행사업자명",
			CEOName:     "대표자",
			Addr:        "서울시",
			ContactID:   "바로빌Username",
			ContactName: "담당자명",
			HP:          "010-0000-0000",
			Email:       "example@contact.com",
		}, Party{
			MgtNum:      NewTSID(),
			CorpName:    "(주)케이넷",
			CorpNum:     "4168138772",
			CEOName:     "이천호",
			Addr:        "광주광역시 북구",
			ContactName: "당당자명",
			HP:          "010-0000-0000",
		}, TaxInvoiceTradeLineItem{
			Amount: 1000,
			Tax:    100,
		},
	)
	res, _ := c.RegistAndIssueTaxInvoice(RegistAndIssueTaxInvoice{
		Invoice:    invoice,
		SendSMS:    true,
		ForceIssue: true,
	})

	t.Log(res.ErrorString())
}
