package model

import "testing"

// Test that the response is marshalled correctly
func TestMarshal(t *testing.T) {
	testResponse := Response{
		Success:  true,
		Title:    "string",
		Price:    59.99,
		StoreURL: "github.com",
	}
	expectedResult := `
		{
			"success": true,
			"title": "string",
			"price": "59.99",
			"storeURL": "github.com"
		}
	`

	testResult, err := testResponse.MarshalResponse()
	if err != nil {
		t.Errorf("Error when calling json.Marshal()\n%v\n", err)
	}

	if string(testResult) != expectedResult {
		t.Errorf("Marshal result did not match expected result\ntestResult: %v\nexpectedResult: %v\n", string(testResult), expectedResult)
	}
}
