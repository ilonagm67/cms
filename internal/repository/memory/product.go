package memory

import (
	"cms/internal/entity"
	"log"
)

type ProductRepository struct {
	products map[string]*entity.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: make(map[string]*entity.Product, 0),
	}
}

func (repo *ProductRepository) Create(name string, product *entity.Product) error {
	repo.products[name] = product
	log.Println(repo.products)
	return nil
}
