package service

import (
	"cms/internal/entity"
	"cms/internal/usecase"
)

type ProductService struct {
	ProductRepo usecase.ProductRepository
}

func NewProductService(ProductRepo usecase.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: ProductRepo}
}

func (Service *ProductService) Delete(name string) error {
	err := Service.ProductRepo.Delete(name)
	if err != nil {
		return err
	}
	return nil
}

func (Service *ProductService) Get(name string, weight int) (*entity.Product, error) {
	product, err := Service.ProductRepo.Get(name, weight)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (Service *ProductService) List() (map[string]map[int]*entity.Product, error) {
	list, err := Service.ProductRepo.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}
