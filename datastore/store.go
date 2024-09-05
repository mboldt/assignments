package datastore

import (
	"fmt"
	"io"
	"log"
)

func Store(w io.Writer, data any) error {
	ser, err := serialize(data)
	if err != nil {
		return fmt.Errorf("failed to serialize data: %w", err)
	}
	_, err = w.Write(ser)
	if err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}
	return nil
}

func Read[T any](r io.Reader) (T, error) {
	var data T
	log.Printf("reader: %v", r)
	b, err := io.ReadAll(r)
	if err != nil {
		return data, fmt.Errorf("failed to read data: %w", err)
	}
	data, err = deserialize[T](b)
	if err != nil {
		return data, fmt.Errorf("failed to deserialize data: %w", err)
	}
	return data, nil
}
