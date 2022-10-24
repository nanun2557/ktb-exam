package services

import (
	"context"
	"fmt"
	"product-service/models"
	"product-service/repository"
	"time"
)

type ProductService interface {
	GetProducts() ([]models.Product, error)
	CreateProduct(p *models.Product) (int64, error)
	UpdateProductById(p *models.Product) error
	DeleteProductById(id int) error
}

type ProductServiceContext struct {
	repo *repository.Repository
}

func (s *ProductServiceContext) GetProducts() ([]models.Product, error) {

	ctx := context.Background()

	// try to get products from cache first
	var products []models.Product
	err := s.repo.ProductRepo.GetCache(ctx, "products", &products)
	if err != nil {
		fmt.Println("redis get error:", err)
		return products, err
	}
	if len(products) > 0 {
		fmt.Println("get product from cache")
		return products, nil
	}

	// get product from mysql
	products, err = s.repo.ProductRepo.Get(nil)
	if err != nil {
		return products, err
	}

	// -- save to redis -------------------------
	err = s.repo.ProductRepo.SetCache(ctx, "products", products, time.Hour*1)
	if err != nil {
		fmt.Println("redis set error:", err)
		return products, err
	}

	return products, err
}

func (s *ProductServiceContext) CreateProduct(p *models.Product) (int64, error) {
	r, err := s.repo.ProductRepo.Create(nil, p)
	return r, err
}

func (s *ProductServiceContext) UpdateProductById(p *models.Product) error {
	err := s.repo.ProductRepo.UpdateById(nil, p)
	return err
}

func (s *ProductServiceContext) DeleteProductById(id int) error {
	err := s.repo.ProductRepo.DeleteById(nil, id)
	return err
}
