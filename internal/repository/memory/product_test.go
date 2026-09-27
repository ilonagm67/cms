package memory_test

import (
	"cms/internal/entity"
	"cms/internal/repository/memory"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestProductAddSuccessfully(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(0, 10)
	product := &entity.Product{Name: ProductName, Weight: ProductWeight}
	repo := memory.NewProductRepository()
	err := repo.Add(product)
	if err != nil {
		t.Errorf("Test Product Add Return Error")
	}
}

func TestProductAddNameError(t *testing.T) {
	ProductName := gofakeit.Name()
	repo := memory.NewProductRepository()
	err := repo.Add(&entity.Product{Name: ProductName})
	if err == nil {
		t.Errorf("Test Product Add Return nil")
	}
}

func TestProductAddWeightError(t *testing.T) {
	ProductWeight := gofakeit.Number(1, 10)
	repo := memory.NewProductRepository()
	err := repo.Add(&entity.Product{Weight: ProductWeight})
	if err == nil {
		t.Errorf("Test Product Add Return nil")
	}
}

func TestProductDeleteSuccessfully(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	repo := memory.NewProductRepository()
	err := repo.Add(&entity.Product{Name: ProductName, Weight: ProductWeight})
	if err != nil {
		t.Errorf("Test Product RepoAdd Delete return Error")
	}
	err = repo.Delete(ProductName, ProductWeight)
	if err != nil {
		t.Errorf("Test Product Delete return Error")
	}
}

func TestProductDeleteError(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	repo := memory.NewProductRepository()
	err := repo.Delete(ProductName, ProductWeight)
	if err == nil {
		t.Errorf("Test Product Delete return nil")
	}
}

func TestProductGetSuccessfully(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	product := &entity.Product{Name: ProductName, Weight: ProductWeight}
	repo := memory.NewProductRepository()
	err := repo.Add(product)
	if err != nil {
		t.Errorf("Test Product RepoAdd Get return Error")
	}
	testproduct, err := repo.Get(ProductName, ProductWeight)
	if err != nil {
		t.Errorf("Test Product Get return Error")
	}
	if testproduct != product {
		t.Errorf("Test Product Get is not equal product")
	}
}

func TestProductGetError(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	product := &entity.Product{Name: ProductName, Weight: ProductWeight}
	repo := memory.NewProductRepository()
	testproduct, err := repo.Get(ProductName, ProductWeight)
	if err == nil {
		t.Errorf("Test Product Get return nil")
	}
	if testproduct == product {
		t.Errorf("Test Product Get is equal product")
	}
}

func TestProductListSuccessfully(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	product := &entity.Product{Name: ProductName, Weight: ProductWeight}
	repo := memory.NewProductRepository()
	err := repo.Add(product)
	if err != nil {
		t.Errorf("Test Product RepoAdd List return Error")
	}
	_, err = repo.List()
	if err != nil {
		t.Errorf("Test Product List return Error")
	}
}

func TestProductListError(t *testing.T) {
	repo := memory.NewProductRepository()
	_, err := repo.List()
	if err == nil {
		t.Errorf("Test Product List return nil")
	}
}
