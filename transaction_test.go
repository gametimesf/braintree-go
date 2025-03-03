package braintree

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateTransactionRiskContextRequest(t *testing.T) {
	tmp := CreateTransactionRiskContextRequest{}
	emptyBody := `{"query":"mutation ($input: CreateTransactionRiskContextInput!) {createTransactionRiskContext(input: $input) { clientMetadataId }}","variables":{"input":{"riskContext":{"fields":[]}}}}`

	bts, err := tmp.GraphQLRequest().Buffer()
	if err != nil {
		t.Fatalf("process of build request is not success")
	}
	if emptyBody != bts.String() {
		t.Fatalf("process of build request is not compare, %s", bts.String())
	}

	tmp.SenderCreatedDate = time.Now().String()
	body := fmt.Sprintf(
		`{"query":"mutation ($input: CreateTransactionRiskContextInput!) {createTransactionRiskContext(input: $input) { clientMetadataId }}","variables":{"input":{"riskContext":{"fields":[{"name":"sender_create_date","value":"%s"}]}}}}`,
		tmp.SenderCreatedDate,
	)

	bts, err = tmp.GraphQLRequest().Buffer()
	if err != nil {
		t.Fatalf("process of build request is not success")
	}
	if body != bts.String() {
		t.Fatalf("process of build request is not compare, req=%s", bts.String())
	}
}
