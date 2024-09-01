package datastore_test

import (
	"strings"
	"testing"

	"github.com/mboldt/assignments/datastore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStoreRead(t *testing.T) {
	type Str struct {
		Data string `json:"data"`
	}
	var w strings.Builder
	src := Str{Data: "foobar"}
	err := datastore.Store(&w, src)
	require.NoError(t, err)
	dst, err := datastore.Read[Str](strings.NewReader(w.String()))
	require.NoError(t, err)
	assert.Equal(t, src, dst)
}
