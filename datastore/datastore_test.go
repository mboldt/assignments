package datastore_test

import (
	"testing"

	"github.com/mboldt/assignments/datastore"
	"github.com/mboldt/assignments/student"

	"github.com/stretchr/testify/assert"
)

func TestAddStudent(t *testing.T) {
	t.Run("when I add a student", func(t *testing.T) {
		d := datastore.New()
		s := student.New("John Doe")
		d.AddStudent(s)
		t.Run("it includes the student in the list", func(t *testing.T) {
			ss := d.ListStudents()
			assert.Len(t, ss, 1)
			assert.Contains(t, ss, s)
		})
	})
}
