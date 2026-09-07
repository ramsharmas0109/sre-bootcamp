package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"srebootcamp/internal/config"
	"srebootcamp/internal/db"
	"srebootcamp/internal/handler"
	"srebootcamp/internal/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	c := config.LoadConfig()
	dbConn := db.InitDB(c)
	defer dbConn.Close()

	h := &handler.Handler{DB: dbConn}
	r := router.SetupRouter(h)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/healthcheck", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "\"ok\"", w.Body.String())
}
