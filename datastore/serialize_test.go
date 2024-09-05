package datastore

import (
	"testing"

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
			ser, err := serialize(tt)
			assert.NoError(t, err)
			t.Logf("%v serilaized to %s", tt, ser)
			des, err := deserialize[Str](ser)
			assert.NoError(t, err)
			assert.Equal(t, tt, des)
		}
	})
}
