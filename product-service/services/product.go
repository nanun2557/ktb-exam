package services

import (
	"product-service/models"
	"product-service/repository"
)

type ProductService interface {
	GetProducts() ([]models.Product, error)
	CreateProduct(p *models.Product) (int64, error)
	UpdateProductById(p *models.Product) (int64, error)
	DeleteProductById(id int) error
}

type ProductServiceContext struct {
	repo *repository.Repository
}

func (s *ProductServiceContext) GetProducts() ([]models.Product, error) {
	r, err := s.repo.ProductRepo.Get(nil)
	return r, err
}

func (s *ProductServiceContext) CreateProduct(p *models.Product) (int64, error) {
	r, err := s.repo.ProductRepo.Create(nil, p)
	return r, err
}

func (s *ProductServiceContext) UpdateProductById(p *models.Product) (int64, error) {
	r, err := s.repo.ProductRepo.UpdateById(nil, p)
	return r, err
}

func (s *ProductServiceContext) DeleteProductById(id int) error {
	err := s.repo.ProductRepo.DeleteById(nil, id)
	return err
}
