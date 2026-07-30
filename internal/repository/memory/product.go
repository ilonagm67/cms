package memory

import "cms/internal/entity"

type ProductRepository struct {
	products map[string]*entity.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: make(map[string]*entity.Product, 0),
	}
}
