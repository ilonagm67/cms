package memory

import (
	"cms/internal/entity"
	"errors"
	"fmt"
)

type ProductRepository struct {
	products map[string]map[int]*entity.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: make(map[string]map[int]*entity.Product, 0),
	}
}

func (repo *ProductRepository) Add(Product *entity.Product) error {
	_, exists := repo.products[Product.Name]
	if !exists {
		repo.products[Product.Name] = make(map[int]*entity.Product)
	}
	repo.products[Product.Name][Product.Weight] = Product
	return nil
}

func (repo *ProductRepository) Delete(Name string) error {
	_, ok := repo.products[Name]
	if ok {
		delete(repo.products, Name)
		return nil
	}
	return fmt.Errorf("Product Not Found: %s", Name)
}

func (repo *ProductRepository) Get(Name string, Weight int) (*entity.Product, error) {
	i, ok := repo.products[Name][Weight]
	if ok {
		return i, nil
	}
	return nil, fmt.Errorf("Product Not Found, Name: %s, Weight: %d", Name, Weight)
}

func (repo *ProductRepository) List() (map[string]map[int]*entity.Product, error) {
	if len(repo.products) > 0 {
		return repo.products, nil
	}
	return nil, errors.New("No Items in Database Products")
}
