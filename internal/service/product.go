package service

import (
	"cms/internal/usecase"
)

type ProductService struct {
	ProductRepo usecase.ProductRepository
}

func NewProductService(ProductRepo usecase.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: ProductRepo}
}
