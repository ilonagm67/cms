package sqlite

import (
	"cms/internal/entity"
	"database/sql"
	"errors"
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

	mock.ExpectQuery("SELECT ID FROM users").WithArgs(ID).WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("INSERT INTO users").WithArgs(ID, User, "1").WillReturnResult(sqlmock.NewResult(1, 1))

	repo := UserRepository{DB: db}
	err = repo.Add(ID, &entity.User{ID: ID, Name: User, Number: "1"})
	if err != nil {
		t.Errorf("Test User Add return error: %v", err)
	}
}

func TestUserAddError(t *testing.T) {
	ID := gofakeit.Int64()
	User := gofakeit.Name()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ID FROM users").WithArgs(ID).WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("INSERT INTO users").WithArgs(ID, User).WillReturnError(errors.New("error"))

	repo := UserRepository{DB: db}
	err = repo.Add(ID, &entity.User{ID: ID, Name: User})
	if err == nil {
		t.Errorf("Test User Add return nil: %v", err)
	}
}

func TestUserSetNameSuccessfully(t *testing.T) {
	ID := gofakeit.Int64()
	Name := gofakeit.Name()
	Phone := "1"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ID FROM users").WithArgs(ID).WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("INSERT INTO users").WithArgs(ID, Name, Phone).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE users").WithArgs(Name, ID).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := UserRepository{DB: db}
	err = repo.Add(ID, &entity.User{ID: ID, Name: Name, Number: Phone})
	if err != nil {
		t.Errorf("Test Update Name User Add return error: %v", err)
	}
	err = repo.SetName(ID, Name)
	if err != nil {
		t.Errorf("Test Update Name return error: %v", err)
	}
}

func TestUserSetNameError(t *testing.T) {
	ID := gofakeit.Int64()
	Name := gofakeit.Name()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ID FROM users").WithArgs(ID).WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("UPDATE users").WithArgs(Name, ID).WillReturnError(errors.New("error"))

	repo := UserRepository{DB: db}
	err = repo.SetName(ID, Name)
	if err == nil {
		t.Errorf("Test Update Name return nil: %v", err)
	}
}

func TestUserSetPhoneSuccessfully(t *testing.T) {
	ID := gofakeit.Int64()
	Name := gofakeit.Name()
	Phone := "1"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ID FROM users").WithArgs(ID).WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("INSERT INTO users").WithArgs(ID, Name, Phone).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE users").WithArgs(Phone, ID).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := UserRepository{DB: db}
	err = repo.Add(ID, &entity.User{ID: ID, Name: Name, Number: Phone})
	if err != nil {
		t.Errorf("Test Update Name User Add return error: %v", err)
	}
	err = repo.SetPhone(ID, Phone)
	if err != nil {
		t.Errorf("Test Update Name return error: %v", err)
	}
}

func TestUserSetPhoneError(t *testing.T) {
	ID := gofakeit.Int64()
	Phone := "1"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ID FROM users").WithArgs(ID).WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("UPDATE users").WithArgs(Phone, ID).WillReturnError(errors.New("error"))

	repo := UserRepository{DB: db}
	err = repo.SetPhone(ID, Phone)
	if err == nil {
		t.Errorf("Test Update Phone return nil: %v", err)
	}
}

func TestUserDeleteSuccessfully(t *testing.T) {
	ID := gofakeit.Int64()
	Name := gofakeit.Name()
	Phone := "1"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ID FROM users").WithArgs(ID).WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("INSERT INTO users").WithArgs(ID, Name, Phone).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("DELETE FROM users").WithArgs(ID).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := UserRepository{DB: db}
	err = repo.Add(ID, &entity.User{ID: ID, Name: Name, Number: Phone})
	if err != nil {
		t.Errorf("Test Delete User Add return error: %v", err)
	}
	err = repo.Delete(ID)
	if err != nil {
		t.Errorf("Test Delete return error: %v", err)
	}
}

func TestUserDeleteError(t *testing.T) {
	ID := gofakeit.Int64()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ID FROM users").WithArgs(ID).WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("DELETE FROM users").WithArgs(ID).WillReturnError(errors.New("error"))

	repo := UserRepository{DB: db}
	err = repo.Delete(ID)
	if err == nil {
		t.Errorf("Test Delete return nil: %v", err)
	}
}

func TestUserGetSuccessfully(t *testing.T) {
	ID := gofakeit.Int64()
	Name := gofakeit.Name()
	Phone := "1"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ID FROM users").WithArgs(ID).WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("INSERT INTO users").WithArgs(ID, Name, Phone).WillReturnResult(sqlmock.NewResult(1, 1))
	rows := sqlmock.NewRows([]string{"id", "name", "number"}).AddRow(ID, Name, Phone)
	mock.ExpectQuery("SELECT ID, Name, Number FROM users WHERE ID = ?").WithArgs(ID).WillReturnRows(rows)

	repo := UserRepository{DB: db}
	err = repo.Add(ID, &entity.User{ID: ID, Name: Name, Number: Phone})
	if err != nil {
		t.Errorf("Test Get User Add return error: %v", err)
	}
	_, err = repo.Get(ID)
	if err != nil {
		t.Errorf("Test Get return error: %v", err)
	}
}

func TestUserGetError(t *testing.T) {
	ID := gofakeit.Int64()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ID, Name, Number FROM users WHERE ID = ?").WithArgs(ID).WillReturnError(sql.ErrNoRows)

	repo := UserRepository{DB: db}
	_, err = repo.Get(ID)
	if err == nil {
		t.Errorf("Test Get return nil: %v", err)
	}
}

func TestUserListSuccessfully(t *testing.T) {
	ID := gofakeit.Int64()
	Name := gofakeit.Name()
	Phone := "1"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ID FROM users").WithArgs(ID).WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("INSERT INTO users").WithArgs(ID, Name, Phone).WillReturnResult(sqlmock.NewResult(1, 1))
	rows := sqlmock.NewRows([]string{"id", "name", "number"}).AddRow(ID, Name, Phone)
	mock.ExpectQuery("SELECT ID, Name, Number FROM users").WillReturnRows(rows)

	repo := UserRepository{DB: db}
	err = repo.Add(ID, &entity.User{ID: ID, Name: Name, Number: Phone})
	if err != nil {
		t.Errorf("Test List User Add return error: %v", err)
	}
	_, err = repo.List()
	if err != nil {
		t.Errorf("Test List return error: %v", err)
	}
}

func TestUserListError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ID, Name, Number FROM users").WillReturnError(errors.New("error"))

	repo := UserRepository{DB: db}
	_, err = repo.List()
	if err == nil {
		t.Errorf("Test List return nil: %v", err)
	}
}
