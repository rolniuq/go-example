package main

import (
	"fmt"
	"sony/props"
	"sony/queries"
	"strings"
)

type Sony struct {
	filter *strings.Builder
}

func NewSony(props *props.SonyProps, makeQueries ...queries.MakeQuery) *Sony {
	sony := &Sony{filter: &strings.Builder{}}

	for _, makeQuery := range makeQueries {
		switch makeQuery() {
		case queries.CustomerId:
			sony.filter.WriteString(makeQuery.MakeQueryCustomer(props.GetCustomerId()))
		case queries.AdsId:
			sony.filter.WriteString(makeQuery.MakeQueryAds(props.GetAdsId()))
		default:
			fmt.Println("unknown query")
		}
	}

	return sony
}

func (s *Sony) Build() string {
	return s.filter.String()
}

func main() {
	sony := NewSony(
		props.NewSonyProps(
			props.WithCustomerId("customerId"),
			props.WithAdsId("adsId"),
		),

		queries.MakeCustomerQuery,
		queries.MakeAdsQuery,
	).Build()

	fmt.Println(sony)
}
