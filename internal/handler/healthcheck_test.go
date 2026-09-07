package handler_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	w := sendRequest(t, http.MethodGet, "/healthcheck", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `"ok"`, w.Body.String())
}
