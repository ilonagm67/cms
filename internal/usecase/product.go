package usecase

import "cms/internal/entity"

type ProductRepository interface {
	Create(name string,product *entity.Product) error
} 
