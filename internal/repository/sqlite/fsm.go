package sqlite

import (
	"cms/internal/entity"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type FSMRepository struct {
	DB *sql.DB
}

func NewFSMRepository() *FSMRepository {
	db, err := sql.Open("sqlite", "data/fsm.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	sqlFSM := `CREATE TABLE IF NOT EXISTS FSM (
		UserID INTEGER,
		State TEXT,
		DataName TEXT,
		DataWeight INTEGER
	);`

	_, err = db.Exec(sqlFSM)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &FSMRepository{DB: db}
}

func (repo *FSMRepository) Set(ID int64, State entity.State) error {
	if ID == 0 || State == "" {
		return fmt.Errorf("ID or State not set!")
	}

	var existingID int64
	checkQuery := `SELECT UserID FROM FSM WHERE UserID = ?`
	err := repo.DB.QueryRow(checkQuery, ID).Scan(&existingID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			insertQuery := `
				INSERT INTO FSM (UserID, State)
				VALUES (?, ?)`
			_, err := repo.DB.Exec(insertQuery, ID, State)
			if err != nil {
				return fmt.Errorf("failed to insert order %d: %w", ID, err)
			}
			return nil
		}
		return fmt.Errorf("failed to check existing order %d: %w", ID, err)
	}
	updateQuery := `
		UPDATE FSM
		SET State = ?
		WHERE UserID = ?`
	_, err = repo.DB.Exec(updateQuery, State, ID)
	if err != nil {
		return fmt.Errorf("failed to update order %d: %w", ID, err)
	}
	return nil
}

func (repo *FSMRepository) SetName(ID int64, Name string) error {
	if ID == 0 || Name == "" {
		return fmt.Errorf("ID or Name not set!")
	}

	var existingName *string
	checkQuery := `SELECT DataName FROM FSM WHERE UserID = ?`
	err := repo.DB.QueryRow(checkQuery, ID).Scan(&existingName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			insertQuery := `
				INSERT INTO FSM (UserID, DataName)
				VALUES (?, ?)`
			_, err := repo.DB.Exec(insertQuery, ID, Name)
			if err != nil {
				return fmt.Errorf("failed to insert fsm dataname %d: %w", ID, err)
			}
			return nil
		}
		return fmt.Errorf("failed to check existing fsm %d: %w", ID, err)
	}
	updateQuery := `
		UPDATE FSM
		SET DataName = ?
		WHERE UserID = ?`
	_, err = repo.DB.Exec(updateQuery, Name, ID)
	if err != nil {
		return fmt.Errorf("failed to update dataname %d: %w", ID, err)
	}
	return nil
}

func (repo *FSMRepository) GetName(ID int64) (string, error) {
	if ID == 0 {
		return "", fmt.Errorf("ID not set!")
	}
	query := `
		SELECT DataName
		FROM FSM
		WHERE UserID = ?`

	var DataName string
	err := repo.DB.QueryRow(query, ID).Scan(&DataName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("DataName FSM ID: %d not found", ID)
		}
		return "", fmt.Errorf("failed to query order %d: %w", ID, err)
	}
	return DataName, nil
}

func (repo *FSMRepository) SetWeight(ID int64, Weight int) error {
	if ID == 0 || Weight == 0 {
		return fmt.Errorf("ID or Weight not set!")
	}

	var existingWeight *int
	checkQuery := `SELECT DataWeight FROM FSM WHERE UserID = ?`
	err := repo.DB.QueryRow(checkQuery, ID).Scan(&existingWeight)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			insertQuery := `
				INSERT INTO FSM (UserID, DataWeight)
				VALUES (?, ?)`
			_, err := repo.DB.Exec(insertQuery, ID, Weight)
			if err != nil {
				return fmt.Errorf("failed to insert order %d: %w", ID, err)
			}
			return nil
		}
		return fmt.Errorf("failed to check existing order %d: %w", ID, err)
	}
	updateQuery := `
		UPDATE FSM
		SET DataWeight = ?
		WHERE UserID = ?`
	_, err = repo.DB.Exec(updateQuery, Weight, ID)
	if err != nil {
		return fmt.Errorf("failed to update order %d: %w", ID, err)
	}
	return nil
}

func (repo *FSMRepository) GetWeight(ID int64) (int, error) {
	if ID == 0 {
		return 0, fmt.Errorf("ID not set!")
	}
	query := `
		SELECT DataWeight
		FROM FSM
		WHERE UserID = ?`

	var DataWeight int
	err := repo.DB.QueryRow(query, ID).Scan(&DataWeight)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("DataWeight FSM ID: %d not found", ID)
		}
		return 0, fmt.Errorf("failed to query FSM ID: %d, %w", ID, err)
	}
	return DataWeight, nil
}

func (repo *FSMRepository) Delete(ID int64) error {
	if ID == 0 {
		return fmt.Errorf("ID not set!")
	}
	query := `DELETE FROM FSM WHERE UserID = ?`
	result, err := repo.DB.Exec(query, ID)
	if err != nil {
		return fmt.Errorf("failed to delete FSM ID: %d, %w", ID, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to fetch rows affected FSM ID: %d, %w", ID, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("FSM ID: %d, not found", ID)
	}
	return nil
}

func (repo *FSMRepository) Get(ID int64) (entity.State, error) {
	if ID == 0 {
		return "", fmt.Errorf("ID not set!")
	}
	query := `
		SELECT State
		FROM FSM
		WHERE UserID = ?`

	var State entity.State
	err := repo.DB.QueryRow(query, ID).Scan(&State)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("State FSM ID: %d not found", ID)
		}
		return "", fmt.Errorf("failed to query FSM ID: %d, %w", ID, err)
	}
	return State, nil
}
