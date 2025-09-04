package queries

type Query string

const (
	CustomerId Query = "customerId"
	AdsId      Query = "adsId"
)

type MakeQuery func() Query

func (q *MakeQuery) MakeQueryCustomer(customerId string) string {
	return customerId
}

func (q *MakeQuery) MakeQueryAds(adsId string) string {
	return adsId
}

var MakeCustomerQuery = func() Query { return CustomerId }

var MakeAdsQuery = func() Query { return AdsId }
