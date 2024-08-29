package templates_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mboldt/assignments/templates"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTitle(t *testing.T) {
	r := httptest.NewRecorder()
	templates.Render(r, "index.html", nil)
	require.Equal(t, http.StatusOK, r.Code, r.Body.String())
	assert.Contains(t, r.Body.String(), "Assignments")
}
