package sqlite

import (
	"cms/internal/entity"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository() *ProductRepository {
	db, err := sql.Open("sqlite", "./products.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()
	sqlProducts := `CREATE TABLE IF NOT EXISTS products (
		Name VARCHAR(255),
		Description VARCHAR(255),
		Image VARCHAR(255),
		Weight INTEGER,
		Price INTEGER
	);`

	_, err = db.Exec(sqlProducts)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) Add(Product *entity.Product) error {
	return nil
}

func (repo *ProductRepository) Delete(Name string) error {
	return nil
}

func (repo *ProductRepository) Get(Name string, Weight int) (*entity.Product, error) {
	return nil, nil
}

func (repo *ProductRepository) List() (map[string]map[int]*entity.Product, error) {
	return nil, nil
}
