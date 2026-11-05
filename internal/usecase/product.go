package usecase

import "cms/internal/entity"

type ProductRepository interface {
	Add(Name string, Weight int) error
	SetDescription(Name string, Weight int, Description string) error
	SetImage(Name string, Weight int, Image string) error
	SetPrice(Name string, Weight int, Price int) error
	Delete(Name string, Weight int) error
	Get(Name string, Weight int) (*entity.Product, error)
	List() (map[string]map[int]*entity.Product, error)
}
