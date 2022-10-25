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
		if err.Error() != "not found key" {
			return products, err
		}
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
	if err != nil {
		return 0, err
	}
	err = s.updateProductCache()
	if err != nil {
		fmt.Println("update product cache fail:", err)
		return 0, err
	}
	return r, err
}

func (s *ProductServiceContext) UpdateProductById(p *models.Product) error {
	err := s.repo.ProductRepo.UpdateById(nil, p)
	if err != nil {
		return err
	}
	err = s.updateProductCache()
	if err != nil {
		fmt.Println("update product cache fail:", err)
		return err
	}

	return err
}

func (s *ProductServiceContext) DeleteProductById(id int) error {
	err := s.repo.ProductRepo.DeleteById(nil, id)
	if err != nil {
		return err
	}
	err = s.updateProductCache()
	if err != nil {
		fmt.Println("update product cache fail:", err)
		return err
	}
	return err
}

func (s *ProductServiceContext) updateProductCache() error {

	// get product from mysql
	var products []models.Product
	var err error

	ctx := context.Background()

	products, err = s.repo.ProductRepo.Get(nil)
	if err != nil {
		return err
	}

	// -- save to redis -------------------------
	err = s.repo.ProductRepo.SetCache(ctx, "products", products, time.Hour*1)
	if err != nil {
		fmt.Println("redis set error:", err)
		return err
	}
	return nil
}
