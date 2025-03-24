package barobill

const (
	TypeIssueDirectionNormal IssueDirection = iota + 1 //정발급
	TypeIssueDirectionInvert                           //역발행
)

const (
	TypeTaxInvoice TaxInvoiceType = iota + 1 // 세금계산서
	TypeInvoice                              // 계산서
	_
	TypeConsignmentTaxInvoice // 위수탁세금계산서
	TypeConsignmentInvoice    // 위수탁계산서
)

const (
	TypeTax TaxType = iota + 1
	TypeSmallTax
	TypeTaxFree
)

const (
	TypePurposeReceipt PurposeType = iota + 1
	TypePurposeClaim
)

const (
	ModifyCodeCorrection         ModifyCode = "1" //기재사항의 착오/정정
	ModifyCodeSupplyPriceChanged ModifyCode = "2" //공급가액의 변동
	ModifyCodeRefundOfGoods      ModifyCode = "3" //재화의 환입
	ModifyCodeContractDestructed ModifyCode = "4" //계약의 해제
	ModifyCodePostOpen           ModifyCode = "5" //내국신용장 사후개설
	ModifyCodeDuplicated         ModifyCode = "6" //착오에 의한 이중발급
)

type IssueDirection int
type TaxInvoiceType int
type TaxType int
type PurposeType int
type ModifyCode string
