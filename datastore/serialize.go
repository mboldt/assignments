package datastore

import "encoding/json"

func Serialize[T any](data T) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Deserialize[T any](data []byte) (T, error) {
	var ret T
	err := json.Unmarshal(data, &ret)
	return ret, err
}
