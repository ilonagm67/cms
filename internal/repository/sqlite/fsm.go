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
	if State == "" {
		return fmt.Errorf("cannot save a empty state")
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
				return fmt.Errorf("failed to insert FSM ID: %d, %w", ID, err)
			}
			return nil
		}
		return fmt.Errorf("failed to check existing FSM ID: %d, %w", ID, err)
	}
	updateQuery := `
		UPDATE FSM
		SET State = ?
		WHERE UserID = ?`
	_, err = repo.DB.Exec(updateQuery, State, ID)
	if err != nil {
		return fmt.Errorf("failed to update FSM ID: %d, %w", ID, err)
	}
	return nil
}

func (repo *FSMRepository) SetName(ID int64, Name string) error {
	if Name == "" {
		return fmt.Errorf("cannot save a empty DataName")
	}

	var existingID int64
	checkQuery := `SELECT UserID FROM FSM WHERE UserID = ?`
	err := repo.DB.QueryRow(checkQuery, ID).Scan(&existingID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			insertQuery := `
				INSERT INTO FSM (UserID, DataName)
				VALUES (?, ?)`
			_, err := repo.DB.Exec(insertQuery, ID, Name)
			if err != nil {
				return fmt.Errorf("failed to insert FSM DataName %d: %w", ID, err)
			}
			return nil
		}
		return fmt.Errorf("failed to check existing FSM ID %d: %w", ID, err)
	}
	updateQuery := `
		UPDATE FSM
		SET DataName = ?
		WHERE UserID = ?`
	_, err = repo.DB.Exec(updateQuery, Name, ID)
	if err != nil {
		return fmt.Errorf("failed to update FSM DataName ID %d: %w", ID, err)
	}
	return nil
}

func (repo *FSMRepository) GetName(ID int64) (string, error) {
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
	if Weight == 0 {
		return fmt.Errorf("cannot save a empty DataWeight")
	}

	var existingID int64
	checkQuery := `SELECT UserID FROM FSM WHERE UserID = ?`
	err := repo.DB.QueryRow(checkQuery, ID).Scan(&existingID)

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
