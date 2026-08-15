package usecase

import "cms/internal/entity"

type ProductRepository interface {
	Add(Name string, weight int, Product *entity.Product) error
	Delete(Name string) error
	Get(Name string, Weight int) (*entity.Product, error)
	List() map[string]map[int]*entity.Product
}
