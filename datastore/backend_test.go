package datastore_test

import (
	"os"
	"testing"

	"github.com/mboldt/assignments/datastore"
	"github.com/mboldt/assignments/student"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileBackend(t *testing.T) {
	file, err := os.CreateTemp("", "assignments")
	t.Logf("temp file: %s", file.Name())
	defer os.Remove(file.Name())
	require.NoError(t, err)
	backend := datastore.NewFileBackend(file.Name())
	data := datastore.New()
	student1 := student.New("John Doe")
	student2 := student.New("Jane Doe")
	data.AddStudent(student1)
	data.AddStudent(student2)
	err = backend.Store(data)
	require.NoError(t, err)
	actual, err := backend.Load()
	require.NoError(t, err)
	assert.Contains(t, actual.ListStudents(), student1)
	assert.Contains(t, actual.ListStudents(), student2)
}

func TestMemoryBackend(t *testing.T) {
	backend := datastore.NewMemoryBackend()
	data := datastore.New()
	student1 := student.New("John Doe")
	student2 := student.New("Jane Doe")
	data.AddStudent(student1)
	data.AddStudent(student2)
	err := backend.Store(data)
	require.NoError(t, err)
	actual, err := backend.Load()
	require.NoError(t, err)
	assert.Contains(t, actual.ListStudents(), student1)
	assert.Contains(t, actual.ListStudents(), student2)
}
