package datastore

import (
	"fmt"
	"os"
)

type Backend interface {
	Store(data *Datastore) error
	Load() (*Datastore, error)
}

type FileBackend struct {
	file string
}

var _ Backend = &FileBackend{}

func NewFileBackend(file string) *FileBackend {
	return &FileBackend{file: file}
}

func (f *FileBackend) Store(data *Datastore) error {
	bytes, err := serialize(data)
	if err != nil {
		return fmt.Errorf("error serializing data upon store: %w", err)
	}
	err = os.WriteFile(f.file, bytes, 0666)
	if err != nil {
		return fmt.Errorf("error writing file %s upon store: %w", f.file, err)
	}
	return nil
}

func (f *FileBackend) Load() (*Datastore, error) {
	bytes, err := os.ReadFile(f.file)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s upon load: %w", f.file, err)
	}
	data, err := deserialize[*Datastore](bytes)
	if err != nil {
		return nil, fmt.Errorf("error reading data upon load: %w", err)
	}
	return data, nil
}

type MemoryBackend struct{
	data *Datastore
}

var _ Backend = &MemoryBackend{}

func NewMemoryBackend() *MemoryBackend {
	return &MemoryBackend{}
}

func (m *MemoryBackend) Store(data *Datastore) error {
	m.data = data
	return nil
}

func (m *MemoryBackend) Load() (*Datastore, error) {
	return m.data, nil
}
