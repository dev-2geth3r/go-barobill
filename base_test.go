package barobill

import (
	"encoding/xml"
	"testing"
)

func TestIssueTaxInvoice(t *testing.T) {
	e := NewEnvelope(IssueTaxInvoiceEx{
		CertKey: "test",
	})

	e.Body.Namespace = testGateway
	output, err := xml.MarshalIndent(e, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", output)
}
