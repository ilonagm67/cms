package sqlite

import (
	"cms/internal/entity"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit"
)

func TestUserAddSuccessfull(t *testing.T) {
	ID := gofakeit.Int64()
	User := gofakeit.Name()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ID FROM users WHERE ID = ?").WithArgs(ID).WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("INSERT INTO users").WithArgs(ID, User, "1").WillReturnResult(sqlmock.NewResult(1, 1))

	repo := UserRepository{DB: db}
	err = repo.Add(ID, &entity.User{ID: ID, Name: User, Number: "1"})
	if err != nil {
		t.Errorf("Test User Add return error: %v", err)
	}
}
