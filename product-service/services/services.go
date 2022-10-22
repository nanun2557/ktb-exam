package services

import "product-service/repository"

type Services struct {
	ProductService
}

func New(r *repository.Repository) *Services {
	return &Services{
		ProductService: &ProductServiceContext{repo: r},
	}
}
