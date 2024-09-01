package datastore

import "encoding/json"

func Serialize[T any](data T) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func Deserialize[T any](data string) (T, error) {
	var ret T
	err := json.Unmarshal([]byte(data), &ret)
	return ret, err
}
