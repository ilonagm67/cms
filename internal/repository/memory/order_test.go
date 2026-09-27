package memory_test

import (
	"cms/internal/entity"
	"cms/internal/repository/memory"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestOrderAddSuccessfully(t *testing.T) {
	UserID := gofakeit.Int64()
	repo := memory.NewOrderRepository()
	Order := &entity.Order{UserID: UserID}
	err := repo.Add(UserID, Order)
	if err != nil {
		t.Errorf("Test Order Add Return error")
	}
}

func TestOrderAddError(t *testing.T) {
	UserID := gofakeit.Int64()
	repo := memory.NewOrderRepository()
	Order := &entity.Order{}
	err := repo.Add(UserID, Order)
	if err == nil {
		t.Errorf("Test Order Add Return nil")
	}
}

func TestOrderAddProductSuccessfully(t *testing.T) {
	UserID := gofakeit.Int64()
	order := &entity.Order{UserID: UserID}

	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	product := &entity.Product{Name: ProductName, Weight: ProductWeight}

	repo := memory.NewOrderRepository()
	err := repo.Add(UserID, order)
	if err != nil {
		t.Errorf("Test Order RepoAdd AddProduct return error")
	}

	err = repo.AddProduct(UserID, product)
	if err != nil {
		t.Errorf("Test Order AddProduct return error")
	}
}

func TestOrderAddProductNameError(t *testing.T) {
	UserID := gofakeit.Int64()
	ProductName := gofakeit.Name()
	product := &entity.Product{Name: ProductName}

	repo := memory.NewOrderRepository()
	err := repo.AddProduct(UserID, product)
	if err == nil {
		t.Errorf("Test Order AddProduct return nil")
	}
}

func TestOrderAddProductWeightError(t *testing.T) {
	UserID := gofakeit.Int64()
	ProductWeight := gofakeit.Number(1, 10)
	product := &entity.Product{Weight: ProductWeight}

	repo := memory.NewOrderRepository()
	err := repo.AddProduct(UserID, product)
	if err == nil {
		t.Errorf("Test Order AddProduct return nil")
	}
}

func TestOrderDeleteProductSuccessfully(t *testing.T) {
	UserID := gofakeit.Int64()
	order := &entity.Order{UserID: UserID}

	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	product := &entity.Product{Name: ProductName, Weight: ProductWeight}

	repo := memory.NewOrderRepository()
	err := repo.Add(UserID, order)
	if err != nil {
		t.Errorf("Test Order RepoAdd AddProduct return error")
	}

	err = repo.AddProduct(UserID, product)
	if err != nil {
		t.Errorf("Test Order RepoAddProduct Delete return error")
	}

	err = repo.DeleteProduct(UserID, ProductName, ProductWeight)
	if err != nil {
		t.Errorf("Test Order DeleteProduct return error")
	}
}

func TestOrderDeleteProductError(t *testing.T) {
	UserID := gofakeit.Int64()
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	repo := memory.NewOrderRepository()
	err := repo.DeleteProduct(UserID, ProductName, ProductWeight)
	if err == nil {
		t.Errorf("Test Order DeleteProduct return nil")
	}
}

func TestOrderGetSuccessfully(t *testing.T) {
	UserID := gofakeit.Int64()
	order := &entity.Order{UserID: UserID}

	repo := memory.NewOrderRepository()
	err := repo.Add(UserID, order)
	if err != nil {
		t.Errorf("Test Order RepoAdd Get return error")
	}
	testorder, err := repo.Get(UserID)
	if err != nil {
		t.Errorf("Test Order Get return error")
	}
	if testorder != order {
		t.Errorf("Test Order Not equal to order!")
	}
}

func TestOrderGetError(t *testing.T) {
	UserID := gofakeit.Int64()
	repo := memory.NewOrderRepository()
	_, err := repo.Get(UserID)
	if err == nil {
		t.Errorf("Test Order Get return nil")
	}
}

func TestOrderListSuccessfully(t *testing.T) {
	UserID := gofakeit.Int64()
	order := &entity.Order{UserID: UserID}
	repo := memory.NewOrderRepository()
	err := repo.Add(UserID, order)
	if err != nil {
		t.Errorf("Test Order RepoAdd Get return error")
	}
	_, err = repo.List()
	if err != nil {
		t.Errorf("Test Order List return error")
	}
}

func TestOrderListError(t *testing.T) {
	repo := memory.NewOrderRepository()
	_, err := repo.List()
	if err == nil {
		t.Errorf("Test Order List return nil")
	}
}
