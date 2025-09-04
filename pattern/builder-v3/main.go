package main

import (
	"strings"
)

type QueryBuilder struct {
	filter *strings.Builder
}

func NewQueryBuilder(queries ...string) *QueryBuilder {
	q := &QueryBuilder{&strings.Builder{}}

	for _, query := range queries {
		q.filter.WriteString(query)
	}

	return q
}

func (b *QueryBuilder) Build() string {
	return b.filter.String()
}

func main() {
	customerId := "customerId"
	picId := "pidId"

	NewQueryBuilder(
		MakeQuery(customerId, CustomerHandler),
		MakeQuery(customerId, AndHandler),
		MakeQuery(picId, PicHandler),
		MakeQuery("ACTIVE", ContractStatusHandler),
	).Build()
}
