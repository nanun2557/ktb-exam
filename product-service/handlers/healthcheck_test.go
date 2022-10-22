package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	e := echo.New()
	t.Log("test healthcheck handler")
	{
		req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		if assert.NoError(t, HealthCheckHandler(c)) {
			assert.Equal(t, rec.Code, http.StatusOK)
			assert.Equal(t, rec.Body.String(), "OK")
		}
	}
}
