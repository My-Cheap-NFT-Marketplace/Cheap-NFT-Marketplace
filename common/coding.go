package common

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func EncodeMessage(message interface{}) ([]byte, error) {
	data := bytes.Buffer{}
	err := gob.NewEncoder(&data).Encode(message)
	if err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}

func DecodeMessage(message []byte, body interface{}) error {
	return json.Unmarshal(message, body)
}
