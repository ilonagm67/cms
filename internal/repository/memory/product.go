package memory

import (
	"cms/internal/entity"
	"errors"
)

type ProductRepository struct {
	products map[string]map[int]*entity.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: make(map[string]map[int]*entity.Product, 0),
	}
}

func (repo *ProductRepository) Add(Name string, Weight int, Product *entity.Product) error {
	repo.products[Name] = make(map[int]*entity.Product)
	repo.products[Name][Weight] = Product
	return nil
}

func (repo *ProductRepository) Delete(Name string) error {
	_, ok := repo.products[Name]
	if ok {
		delete(repo.products, Name)
		return nil
	} else {
		return errors.New("Product Not Found")
	}
}

func (repo *ProductRepository) Get(Name string, Weight int) (*entity.Product, error) {
	i, ok := repo.products[Name][Weight]
	if ok {
		return i, nil
	} else {
		return nil, errors.New("Product Not Found")
	}
}

func (repo *ProductRepository) List() map[string]map[int]*entity.Product {
	return repo.products
}
