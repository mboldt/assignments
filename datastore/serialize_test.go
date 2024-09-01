package datastore_test

import (
	"testing"

	"github.com/mboldt/assignments/datastore"
	"github.com/stretchr/testify/assert"
)

func TestSeailalize(t *testing.T) {
	t.Run("serialize and deserialize are inverses", func(t *testing.T) {
		type Str struct {Data string `json:"data"`}
		tests := []Str{
			{Data: "foo"},
			{Data: "bar"},
		}
		for _, tt := range tests {
			ser, err := datastore.Serialize(tt)
			assert.NoError(t, err)
			t.Logf("%v serilaized to %s", tt, ser)
			des, err := datastore.Deserialize[Str](ser)
			assert.NoError(t, err)
			assert.Equal(t, tt, des)
		}
	})
}
