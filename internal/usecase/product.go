package usecase

import "cms/internal/entity"

type ProductRepository interface {
	Add(Name string,Product *entity.Product) error
	Delete(Name string) error
	Get(Name string) (*entity.Product,error)
	List() (map[string]*entity.Product)
} 
