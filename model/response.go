package model

import "encoding/json"

// A Response represents the data being sent back the to the clietn
// Success is if call was successful
// Title is the title of the game being requested
// Price is the price of the game
// StoreURL is the url for the game to be purchased
type Response struct {
	Success  bool    `json:"success"`
	Title    string  `json:"title"`
	Price    float64 `json:"price,string"`
	StoreURL string  `json:"storeURL"`
}

// MarshalResponse takes a Response struct and returns back a marshalled
// json response
func MarshalResponse(response Response) ([]byte, error) {
	jsonMarshal, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return jsonMarshal, nil
}

// UnMarshalResponse takes a raw json data and marshals the data into a
// Response struct
func UnMarshalResponse(rawResponse []byte) (Response, error) {
	var response Response
	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}
