package barobill

import "github.com/secr3t/req/v3"

const (
	testGateway       = "https://testws.baroservice.com"
	productionGateway = "https://ws.baroservice.com"

	contentTypeSoapXml = "application/soap+xml; charset=utf-8"

	corporateStatusEndpoint = "CORPSTATE.asmx"
)

var (
	defaultClient = req.C().
		EnableDumpAll().
		SetUserAgent("").
		SetBaseURL(testGateway)
)

type Client struct {
	rc               *req.Client
	namespace        string
	certKey          string
	corpNum          string
	barobillUsername string
}

// NewClient default baseURL is testGateway
// @param certKey 바로빌 연동인증키
// @param corpNum 바로빌 연동인증키에 맞는 사업자 번호
func NewClient(certKey, corpNum string, barobillUsernameArgs ...string) *Client {
	var barobillUsername string
	if len(barobillUsernameArgs) == 1 {
		barobillUsername = barobillUsernameArgs[0]
	}
	return &Client{
		namespace:        testGateway,
		rc:               defaultClient.Clone(),
		certKey:          certKey,
		corpNum:          corpNum,
		barobillUsername: barobillUsername,
	}
}

func (c *Client) TestGateway() *Client {
	c.namespace = testGateway
	c.rc.SetBaseURL(testGateway)
	return c
}

func (c *Client) ProductionGateway() *Client {
	c.namespace = productionGateway
	c.rc.SetBaseURL(productionGateway).DisableDumpAll()
	return c
}
