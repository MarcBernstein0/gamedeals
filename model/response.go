package model

import "encoding/json"

type Response struct {
	Success  bool    `json:"success"`
	Title    string  `json:"title"`
	Price    float64 `json:"price,string"`
	StoreURL string  `json:"storeURL"`
}

func MarshalResponse(response Response) ([]byte, error) {
	jsonMarshal, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return jsonMarshal, nil
}

func UnMarshalResponse(rawResponse []byte) (Response, error) {
	var response Response
	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}
