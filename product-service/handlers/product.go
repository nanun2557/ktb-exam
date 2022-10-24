package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"product-service/models"
	"product-service/services"
	"product-service/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	services.ProductService
}

func (h *ProductHandler) GetProducts(c echo.Context) error {
	r, err := h.ProductService.GetProducts()

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var p *models.Product

	if err := c.Bind(&p); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}

	r, err := h.ProductService.CreateProduct(p)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

func (h *ProductHandler) UpdateProductById(c echo.Context) error {
	var p *models.Product
	// get data from request payload
	if err := c.Bind(&p); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, utils.Error{Message: "args is invalid"})
	}

	err := h.ProductService.UpdateProductById(p)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, utils.Error{Message: "not found"})
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, "update successfully")
}

func (h *ProductHandler) DeleteProductById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, utils.Error{Message: "ID is Invalid"})
	}

	err = h.ProductService.DeleteProductById(id)

	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, utils.Error{Message: "not found"})
	} else if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, "OK")
}
