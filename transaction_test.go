package braintree

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestCreateTransactionRiskContextRequest(t *testing.T) {
	tmp := CreateTransactionRiskContextRequest{}
	emptyBody := `{"riskContext":{"fields":[]}}`

	bts, err := json.Marshal(tmp.Request())
	if err != nil {
		t.Fatalf("process of build request is not success")
	}
	if emptyBody != string(bts) {
		t.Fatalf("process of build request is not compare")
	}

	tmp.SenderCreatedDate = time.Now().String()
	body := fmt.Sprintf(
		"{\"riskContext\":{\"fields\":[{\"name\":\"sender_create_date\",\"value\":\"%s\"}]}}",
		tmp.SenderCreatedDate)
	bts, err = json.Marshal(tmp.Request())
	if err != nil {
		t.Fatalf("process of build request is not success")
	}
	if body != string(bts) {
		t.Fatalf("process of build request is not compare, req=%s", string(bts))
	}
}
