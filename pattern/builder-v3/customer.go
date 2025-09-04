package main

import "fmt"

const (
	Id                = "id"
	AdStatus          = "ad_status"
	And               = "AND"
	ContractEndDate   = "contracts.end_date"
	ContractStartDate = "contracts.start_date"
	CustomerPics      = "customer_pics"
	ProductionStatus  = "production_status"
)

type QueryFunc func(string, HandlerFunc) string

var MakeQuery QueryFunc = func(id string, h HandlerFunc) string {
	return h(id)
}

type HandlerFunc func(string) string

var (
	CustomerHandler HandlerFunc = func(id string) string {
		return fmt.Sprintf("(%s=%s)", Id, id)
	}

	AndHandler HandlerFunc = func(id string) string {
		if id == "" {
			return ""
		}
		return fmt.Sprintf(" %s ", And)
	}

	PicHandler HandlerFunc = func(id string) string {
		return fmt.Sprintf("(%s=%s)", CustomerPics, id)
	}

	ContractStatusHandler HandlerFunc = func(s string) string {
		currentTime := "CURRENT_TIMESTAMP"

		switch s {
		case "ACTIVE":
			return fmt.Sprintf(`(%s <= "%s" %s %s >= "%s")`, ContractStartDate, currentTime, And, ContractEndDate, currentTime)
		case "EXPIRED":
			return fmt.Sprintf(`(%s < "%s")`, ContractEndDate, currentTime)
		case "NOT_ACTIVE":
			return fmt.Sprintf(`(%s > "%s")`, ContractStartDate, currentTime)
		default:
			return ""
		}
	}

	AdStatusHandler HandlerFunc = func(s string) string {
		return fmt.Sprintf("(%s=%s)", AdStatus, s)
	}

	ProductStatusHandler HandlerFunc = func(s string) string {
		return fmt.Sprintf("(%s=%s)", ProductionStatus, s)
	}
)
