package memory

import (
	"cms/internal/entity"
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

func (repo *ProductRepository) Add(name string, weight int) error {
	if name == "" || weight == 0 {
		return fmt.Errorf("Name or Weight not set!")
	}
	_, exists := repo.products[name]
	if !exists {
		repo.products[name] = make(map[int]*entity.Product)
	}
	repo.products[name][weight] = &entity.Product{Name: name, Weight: weight}
	return nil
}

func (repo *ProductRepository) SetDescription(name string, weight int, description string) error {
	if name == "" || weight == 0 {
		return fmt.Errorf("Name or Weight not set!")
	}
	if description == "" {
		return fmt.Errorf("Description not found!")
	}

	product, exists := repo.products[name][weight]
	if !exists {
		return fmt.Errorf("Product %s:%d Not Found!", name, weight)
	}

	product.Description = description
	return nil
}

func (repo *ProductRepository) SetImage(name string, weight int, image string) error {
	if name == "" || weight == 0 {
		return fmt.Errorf("Name or Weight not set!")
	}
	if image == "" {
		return fmt.Errorf("Image not set!")
	}

	product, exists := repo.products[name][weight]
	if !exists {
		return fmt.Errorf("Product %s:%d Not Found!", name, weight)
	}

	product.Image = image
	return nil
}

func (repo *ProductRepository) SetPrice(name string, weight int, price int) error {
	if name == "" || weight == 0 {
		return fmt.Errorf("Name or Weight not set!")
	}
	if price == 0 {
		return fmt.Errorf("Price not set!")
	}

	product, exists := repo.products[name][weight]
	if !exists {
		return fmt.Errorf("Product %s:%d Not Found!", name, weight)
	}

	product.Price = price
	return nil
}

func (repo *ProductRepository) Delete(Name string, Weight int) error {
	if Name == "" || Weight == 0 {
		return fmt.Errorf("Name or Weight not set!")
	}

	weights, ok := repo.products[Name]
	if !ok {
		return fmt.Errorf("ProductName: %s Not Found!", Name)
	}
	_, ok = weights[Weight]
	if !ok {
		return fmt.Errorf("ProductWeight: %d Not Found!", Weight)
	}

	delete(weights, Weight)

	if len(weights) == 0 {
		delete(repo.products, Name)
	}

	return nil
}

func (repo *ProductRepository) Get(Name string, Weight int) (*entity.Product, error) {
	if Name == "" || Weight == 0 {
		return nil, fmt.Errorf("Name or Weight not set!")
	}
	weights, ok := repo.products[Name]
	if !ok {
		return nil, fmt.Errorf("ProductName: %s Not Found!", Name)
	}
	product, ok := weights[Weight]
	if !ok {
		return nil, fmt.Errorf("ProductWeight: %d Not Found!", Weight)
	}
	return product, nil
}

func (repo *ProductRepository) List() (map[string]map[int]*entity.Product, error) {
	if len(repo.products) == 0 {
		return nil, fmt.Errorf("No Items in Database Products")
	}
	return repo.products, nil
}
