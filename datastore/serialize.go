package datastore

import "encoding/json"

func serialize[T any](data T) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func deserialize[T any](data []byte) (T, error) {
	var ret T
	err := json.Unmarshal(data, &ret)
	return ret, err
}
