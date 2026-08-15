package service

import (
	"cms/internal/entity"
	"cms/internal/usecase"
	"encoding/json"
)

type ProductService struct {
	ProductRepo usecase.ProductRepository
}

func NewProductService(ProductRepo usecase.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: ProductRepo}
}

func (service *ProductService) Add(name string, weight int) {
	service.ProductRepo.Add(&entity.Product{Name: name, Weight: weight})
}

func (Service *ProductService) Delete(name string) bool {
	err := Service.ProductRepo.Delete(name)
	if err != nil {
		return false
	}
	return true
}

func (Service *ProductService) Get(name string, weight int) ([]byte, error) {
	user, err := Service.ProductRepo.Get(name, weight)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (Service *ProductService) List() ([]byte, error) {
	list := Service.ProductRepo.List()
	json, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	return json, nil
}
