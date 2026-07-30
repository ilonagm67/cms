package service

import (
	"cms/internal/usecase"
	"cms/internal/entity"
)

type ProductService struct {
	ProductRepo usecase.ProductRepository
}

func NewProductService(ProductRepo usecase.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: ProductRepo}
}

func(service *ProductService) Add(name string) {
	service.ProductRepo.Create(name,&entity.Product{Name: name})
}
