package sqlite

import (
	"cms/internal/entity"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository() *OrderRepository {
	db, err := sql.Open("sqlite", "data/orders.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	sqlOrders := `CREATE TABLE IF NOT EXISTS orders (
		CustomerID INTEGER,
		Products TEXT,
		PayType TEXT,
		Address TEXT,
		Delivery TEXT
	);`

	_, err = db.Exec(sqlOrders)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) Add(customerID int64, order *entity.Order) error {
	if order == nil {
		return errors.New("cannot save a nil order")
	}

	encodedProducts, err := json.Marshal(order.Products)
	if err != nil {
		return fmt.Errorf("failed to marshal order products: %w", err)
	}
	jsonStr := string(encodedProducts)

	var existingID int64
	checkQuery := `SELECT CustomerID FROM orders WHERE CustomerID = ?`
	err = repo.DB.QueryRow(checkQuery, customerID).Scan(&existingID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			insertQuery := `
				INSERT INTO orders (CustomerID, Products, PayType, Address, Delivery)
				VALUES (?, ?, ?, ?, ?)`
			_, err := repo.DB.Exec(insertQuery, customerID, jsonStr, order.PayType, order.Address, order.Delivery)
			if err != nil {
				return fmt.Errorf("failed to insert order %d: %w", customerID, err)
			}
			return nil
		}
		return fmt.Errorf("failed to check existing order %d: %w", customerID, err)
	}

	updateQuery := `
		UPDATE orders
		SET Products = ?, PayType = ?, Address = ?, Delivery = ?
		WHERE CustomerID = ?`
	_, err = repo.DB.Exec(updateQuery, jsonStr, order.PayType, order.Address, order.Delivery, customerID)
	if err != nil {
		return fmt.Errorf("failed to update order %d: %w", customerID, err)
	}

	return nil
}

func (repo *OrderRepository) DeleteProduct(customerID int64, product string, weight int) error {
	query := `SELECT Products FROM orders WHERE CustomerID = ?`
	var jsonBytes []byte

	err := repo.DB.QueryRow(query, customerID).Scan(&jsonBytes)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("order for customer %d not found", customerID)
		}
		return fmt.Errorf("failed to fetch order %d: %w", customerID, err)
	}

	var products map[string]map[int]*entity.Product
	if len(jsonBytes) > 0 {
		if err := json.Unmarshal(jsonBytes, &products); err != nil {
			return fmt.Errorf("failed to unmarshal products for customer %d: %w", customerID, err)
		}
	}

	weights, productExists := products[product]
	if !productExists {
		return fmt.Errorf("product category %q not found for customer %d", product, customerID)
	}

	if _, weightExists := weights[weight]; !weightExists {
		return fmt.Errorf("weight %d not found under product %q", weight, product)
	}

	delete(weights, weight)

	if len(weights) == 0 {
		delete(products, product)
	}

	updatedJSON, err := json.Marshal(products)
	if err != nil {
		return fmt.Errorf("failed to marshal updated products: %w", err)
	}

	updateQuery := `UPDATE orders SET Products = ? WHERE CustomerID = ?`
	if _, err := repo.DB.Exec(updateQuery, string(updatedJSON), customerID); err != nil {
		return fmt.Errorf("failed to update order %d after deletion: %w", customerID, err)
	}

	return nil
}

func (repo *OrderRepository) Get(ID int64) (*entity.Order, error) {
	query := `SELECT CustomerID, Products, PayType, Address, Delivery FROM orders WHERE CustomerID = ?`

	var (
		order     entity.Order
		jsonBytes []byte
	)

	err := repo.DB.QueryRow(query, ID).Scan(
		&order.UserID,
		&jsonBytes,
		&order.PayType,
		&order.Address,
		&order.Delivery,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("order %d not found", ID)
		}
		return nil, fmt.Errorf("failed to query order %d: %w", ID, err)
	}

	if len(jsonBytes) > 0 {
		var products map[string]map[int]*entity.Product
		if err := json.Unmarshal(jsonBytes, &products); err != nil {
			return nil, fmt.Errorf("failed to unmarshal product list for order %d: %w", ID, err)
		}
		order.Products = products
	}

	return &order, nil
}

func (repo *OrderRepository) List() (map[int64]*entity.Order, error) {
	query := `SELECT CustomerID, Products, PayType, Address, Delivery FROM orders`

	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query orders: %w", err)
	}

	orders := make(map[int64]*entity.Order)

	for rows.Next() {
		var (
			order        entity.Order
			jsonProducts []byte
		)

		err := rows.Scan(
			&order.UserID,
			&jsonProducts,
			&order.PayType,
			&order.Address,
			&order.Delivery,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order row: %w", err)
		}

		if len(jsonProducts) > 0 {
			var products map[string]map[int]*entity.Product

			if err := json.Unmarshal(jsonProducts, &products); err != nil {
				return nil, fmt.Errorf("failed to unmarshal products for user %d: %w", order.UserID, err)
			}
			order.Products = products
		}

		orders[order.UserID] = &order
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating order rows: %w", err)
	}

	return orders, nil
}
