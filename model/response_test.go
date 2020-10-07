package model

import (
	"reflect"
	"testing"
)

// Test that the response is marshalled correctly
func TestMarshal(t *testing.T) {
	testResponse := Response{
		Success:  true,
		Title:    "string",
		Price:    59.99,
		StoreURL: "github.com",
	}
	expectedResult := "{\"success\":true,\"title\":\"string\",\"price\":\"59.99\",\"storeURL\":\"github.com\"}"

	testResult, err := MarshalResponse(testResponse)
	if err != nil {
		t.Errorf("Error when calling json.Marshal()\n%v\n", err)
	}

	if string(testResult) != expectedResult {
		t.Errorf("Marshal result did not match expected result\ntestResult: %v\nexpectedResult: %v\n", string(testResult), expectedResult)
	}

	if err != nil {
		t.Errorf("Error when calling json.Marshal()\n%v\n", err)
	}

}

func TestUnMarshal(t *testing.T) {
	testResponse := []byte("{\"success\":true,\"title\":\"string\",\"price\":\"59.99\",\"storeURL\":\"github.com\"}")
	expectedResponse := Response{
		Success:  true,
		Title:    "string",
		Price:    59.99,
		StoreURL: "github.com",
	}

	testResult, err := UnMarshalResponse(testResponse)
	if err != nil {
		t.Errorf("Error when calling json.UnMarshal()\n%v\n", err)
	}

	if !reflect.DeepEqual(testResult, expectedResponse) {
		t.Errorf("UnMarshaled results did not match\ntestResult: %+v\nexpectedResult: %+v\n", testResult, expectedResponse)
	}

}
