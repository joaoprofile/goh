package core

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func FromJSON(objJSON []byte) (interface{}, error) {
	var obj any
	err := json.Unmarshal([]byte(objJSON), &obj)
	if err != nil {
		return nil, err
	}
	return objJSON, nil
}

func ToJson(obj interface{}) ([]byte, error) {
	objJSON, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return objJSON, nil
}

func Encoder(obj interface{}) (bytes.Buffer, error) {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(obj); err != nil {
		return buffer, err
	}
	return buffer, nil
}

func Decoder(objBuffer []byte) (interface{}, error) {
	buffer := bytes.NewReader(objBuffer)

	var obj any
	if err := gob.NewDecoder(buffer).Decode(&obj); err != nil {
		return nil, err
	}
	return obj, nil
}
