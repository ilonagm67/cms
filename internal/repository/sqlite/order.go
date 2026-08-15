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
	db, err := sql.Open("sqlite", "./order.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()
	sql := `CREATE TABLE orders (
		id INTEGER PRIMARY KEY,
	);`

	db.Exec(sql)
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) Add(ID int64, Order *entity.Order) error {
	return nil
}

func (repo *OrderRepository) Get(ID int64) (*entity.Order, error) {
	return nil, nil
}

func (repo *OrderRepository) List() map[int64]*entity.Order {
	return nil
}
