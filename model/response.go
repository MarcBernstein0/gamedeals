package model

import "encoding/json"

type Response struct {
	Success  bool    `json:"success"`
	Title    string  `json:"title"`
	Price    float64 `json:"price,string"`
	StoreURL string  `json:"storeURL"`
}

func (r Response) MarshalResponse() ([]byte, error) {
	jsonMarshal, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return jsonMarshal, nil
}
