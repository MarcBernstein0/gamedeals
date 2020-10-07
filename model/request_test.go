package model

import (
	"testing"
)

func TestRequestMarshal(t *testing.T) {
	testRequest := Request{
		Title:  "string",
		System: "PC",
	}

	expextedRequestResult := "{\"title\":\"string\",\"system\":\"PC\"}"

	testResult, err := MarshalRequest(testRequest)
	if err != nil {
		t.Errorf("Error when calling json.Marshal()\n%v\n", err)
	}
	if string(testResult) != expextedRequestResult {
		t.Errorf("Marshal result did not match expected result\ntestResult: %v\nexpectedRequestResult: %v\n", string(testResult), expextedRequestResult)
	}
}
