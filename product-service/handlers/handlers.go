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

	e.GET("/healthcheck", HealthCheckHandler)

}

// unc SetApi(e *echo.Echo, h *Handlers, m echo.MiddlewareFunc) {
func SetApi(e *echo.Echo, h *Handlers) {
	g := e.Group("/api/v1")
	// g.Use(m)

	// Product
	g.GET("/product", h.ProductHandler.GetProducts)
	g.POST("/product", h.ProductHandler.CreateProduct)
	g.PUT("/product", h.ProductHandler.UpdateProductById)
	g.DELETE("/product/:id", h.ProductHandler.DeleteProductById)

}
