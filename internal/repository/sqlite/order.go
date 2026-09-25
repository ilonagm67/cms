package sqlite

import (
	"cms/internal/entity"
	"database/sql"
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
	defer db.Close()
	sqlOrders := `CREATE TABLE IF NOT EXISTS orders (
		OrderID INTEGER,
		CustomerID INTEGER,
		PayType VARCHAR(255),
		Address VARCHAR(255),
		Delivery VARCHAR(255)
	);`

	_, err = db.Exec(sqlOrders)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) Add(ID int64, Order *entity.Order) error {
	return nil
}

func (repo *OrderRepository) Get(ID int64) (*entity.Order, error) {
	return nil, nil
}

func (repo *OrderRepository) List() (map[int64]*entity.Order, error) {
	return nil, nil
}
