package sqlite

import (
	"cms/internal/entity"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository() *UserRepository {
	db, err := sql.Open("sqlite", "./users.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()
	sqlUsers := `CREATE TABLE IF NOT EXISTS users (
		CustomerID INTEGER,
		Name VARCHAR(255),
		State VARCHAR(255),
		Number VARCHAR(255)
	);`

	_, err = db.Exec(sqlUsers)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &UserRepository{DB: db}
}

func (repo *UserRepository) Add(ID int64, User *entity.User) error {
	return nil
}

func (repo *UserRepository) Delete(ID int64) error {
	return nil
}

func (repo *UserRepository) Get(ID int64) (*entity.User, error) {
	return nil, nil
}

func (repo *UserRepository) List() (map[int64]*entity.User, error) {
	return nil, nil
}
