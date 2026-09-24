package sqlite

import (
	"cms/internal/entity"
	"database/sql"
	"errors"
	"fmt"
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
	sqlProducts := `CREATE TABLE IF NOT EXISTS products (
		Name TEXT,
		Description TEXT,
		Image TEXT,
		Weight INTEGER,
		Count INTEGER,
		Price INTEGER
	);`

	_, err = db.Exec(sqlProducts)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) Add(product *entity.Product) error {
	var existingName string
	query := `SELECT Name FROM products WHERE Name = ?`
	err := repo.DB.QueryRow(query, product.Name).Scan(&existingName)

	if errors.Is(err, sql.ErrNoRows) {
		query := `
		INSERT INTO products (Name, Description, Image, Weight, Count, Price)
		VALUES (?, ?, ?, ?, ?, ?)`
		_, err := repo.DB.Exec(query,
			product.Name,
			product.Description,
			product.Image,
			product.Weight,
			product.Count,
			product.Price,
		)
		if err != nil {
			return fmt.Errorf("failed to insert product %s: %w", product.Name, err)
		}
		return nil
	}
	updateQuery := `UPDATE products SET Name = ?, Description = ?, Image = ?, Weight = ?, Count = ?,Price = ? WHERE Name = ?`
	_, err = repo.DB.Exec(updateQuery, product.Name, product.Description, product.Image, product.Weight, product.Count, product.Price, product.Name)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}
	return nil
}

func (repo *ProductRepository) Delete(name string) error {
	query := `DELETE FROM products WHERE Name = ?`

	result, err := repo.DB.Exec(query, name)
	if err != nil {
		return fmt.Errorf("failed to delete product %s: %w", name, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product %s not found", name)
	}

	return nil
}

func (repo *ProductRepository) Get(name string, weight int) (*entity.Product, error) {
	query := `SELECT Name, Description, Image, Weight, Count, Price FROM products WHERE Name = ? AND Weight = ?`

	var product entity.Product

	err := repo.DB.QueryRow(query, name, weight).Scan(&product.Name, &product.Description, &product.Image, &product.Weight, &product.Count, &product.Price)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("product %s with weight %d not found", name, weight)
	} else if err != nil {
		return nil, fmt.Errorf("failed to query product: %w", err)
	}

	return &product, nil
}

func (repo *ProductRepository) List() (map[string]map[int]*entity.Product, error) {
	query := `SELECT Name, Description, Image, Weight, Count, Price FROM products`

	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query products list: %w", err)
	}
	defer rows.Close()

	result := make(map[string]map[int]*entity.Product)

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(
			&product.Name,
			&product.Description,
			&product.Image,
			&product.Weight,
			&product.Count,
			&product.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %w", err)
		}

		if _, exists := result[product.Name]; !exists {
			result[product.Name] = make(map[int]*entity.Product)
		}

		p := product
		result[product.Name][product.Weight] = &p
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return result, nil
}
