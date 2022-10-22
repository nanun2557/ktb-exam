package handlers

import (
	"product-service/services"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	ProductHandler ProductHandler
}

func New(s *services.Services) *Handlers {
	return &Handlers{
		ProductHandler: ProductHandler{s.ProductService},
	}
}

func SetDefault(e *echo.Echo) {

	// e.GET("/", func(c echo.Context) error {
	// 	return c.Render(http.StatusOK, "data", configs.Auth0Config)
	// })
	e.GET("/healthcheck", HealthCheckHandler)
	// e.GET("/swagger/*", echoSwagger.WrapHandler)
}
