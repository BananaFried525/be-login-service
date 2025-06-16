package utils

import "encoding/json"

func JsonMarshal(data interface{}) ([]byte, error) {
	data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return data.([]byte), nil
}
