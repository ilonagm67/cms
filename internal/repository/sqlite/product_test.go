package sqlite

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit"
)

func TestSetDescriptionSuccessfully(t *testing.T) {
	ProductName := gofakeit.Name()
	ProductWeight := gofakeit.Number(1, 10)
	ProductDescription := gofakeit.Extension()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT Name FROM products").WithArgs(ProductName, ProductWeight).WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("INSERT INTO products").WithArgs(ProductName, ProductWeight).WillReturnResult(sqlmock.NewResult(1, 1))
	rows := sqlmock.NewRows([]string{"name", "weight"}).AddRow(ProductName, ProductWeight)
	mock.ExpectQuery("SELECT Name FROM products").WithArgs(ProductName, ProductWeight).WillReturnRows(rows)
	mock.ExpectExec("UPDATE products").WithArgs(ProductDescription, ProductName, ProductWeight).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := ProductRepository{DB: db}
	err = repo.Add(ProductName, ProductWeight)
	if err != nil {
		t.Errorf("Test Product RepoAdd SetDescription return Error: %v", err)
	}
	err = repo.SetDescription(ProductName, ProductWeight, ProductDescription)
	if err != nil {
		t.Errorf("Test Product SetDescription return Error: %v", err)
	}
}
