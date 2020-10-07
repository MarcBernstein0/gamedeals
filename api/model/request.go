package model

import "encoding/json"

type Request struct {
	Title  string `json:"title"`
	System string `json:"system"`
}

func MarshalRequest(request Request) ([]byte, error) {
	jsonMarshal, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	return jsonMarshal, err
}

func UnMarshalRequest(rawRequest []byte) (Request, error) {
	var request Request
	err := json.Unmarshal(rawRequest, &request)
	if err != nil {
		return Request{}, err
	}
	return request, nil
}
