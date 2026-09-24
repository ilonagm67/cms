package sqlite

import (
	"cms/internal/entity"
	"database/sql"
	"errors"
	"fmt"
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
	sqlUsers := `CREATE TABLE IF NOT EXISTS users (
		ID INTEGER,
		Name VARCHAR(255),
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
	var existingID int64
	query := `SELECT ID FROM users WHERE ID = ?`
	err := repo.DB.QueryRow(query, ID).Scan(&existingID)

	if errors.Is(err, sql.ErrNoRows) {
		insertQuery := `INSERT INTO users (ID,Name,Number) VALUES (?,?,?)`
		_, err := repo.DB.Exec(insertQuery, ID, User.Name, User.Number)
		if err != nil {
			return fmt.Errorf("failed to insert user: %w", err)
		}
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to check existing user: %w", err)
	}

	updateQuery := `UPDATE users SET Name = ?, Number = ? WHERE ID = ?`
	_, err = repo.DB.Exec(updateQuery, User.Name, User.Number, ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (repo *UserRepository) Delete(ID int64) error {
	query := `DELETE FROM users WHERE ID = ?`
	result, err := repo.DB.Exec(query, ID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to fetch rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user: %d, not found", ID)
	}

	return nil
}

func (repo *UserRepository) Get(ID int64) (*entity.User, error) {
	query := `SELECT ID, Name, Number FROM users WHERE ID = ?`
	var user entity.User

	err := repo.DB.QueryRow(query, ID).Scan(&user.ID, &user.Name, &user.Number)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user: %d, not found", ID)
	} else if err != nil {
		return nil, fmt.Errorf("failed to query user: %w", err)
	}
	return &user, nil
}

func (repo *UserRepository) List() (map[int64]*entity.User, error) {
	rows, _ := repo.DB.Query("SELECT ID,Name,Number FROM users")
	result := make(map[int64]*entity.User)
	for rows.Next() {
		var user entity.User
		rows.Scan(&user.ID, &user.Name, &user.Number)
		result[user.ID] = &user
	}
	return result, nil
}
