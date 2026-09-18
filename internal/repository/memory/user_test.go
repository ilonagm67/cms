package memory_test

import (
	"cms/internal/entity"
	"cms/internal/repository/memory"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestUserAddSuccessfull(t *testing.T) {
	ID := gofakeit.Int64()
	User := gofakeit.Name()
	repo := memory.NewUserRepository()
	err := repo.Add(ID, &entity.User{ID: ID, Name: User})
	if err != nil {
		t.Errorf("Test Repo Add return error")
	}
}

func TestUserAddError(t *testing.T) {
	ID := gofakeit.Int64()
	repo := memory.NewUserRepository()
	err := repo.Add(ID, &entity.User{})
	if err == nil {
		t.Errorf("Test Repo Add return nil")
	}
}

func TestUserDeleteSuccessfull(t *testing.T) {
	ID := gofakeit.Int64()
	User := &entity.User{ID: ID}
	repo := memory.NewUserRepository()
	err := repo.Add(ID, User)
	if err != nil {
		t.Errorf("Test Repo Delete useradd return error")
	}
	err = repo.Delete(ID)
	if err != nil {
		t.Errorf("Test Repo Delete return error")
	}
}

func TestUserDeleteError(t *testing.T) {
	ID := gofakeit.Int64()
	repo := memory.NewUserRepository()
	err := repo.Delete(ID)
	if err == nil {
		t.Errorf("Test Repo Delete return nil")
	}
}

func TestUserGetSuccessfull(t *testing.T) {
	ID := gofakeit.Int64()
	User := &entity.User{ID: ID}
	repo := memory.NewUserRepository()
	err := repo.Add(ID, User)
	if err != nil {
		t.Errorf("Test Repo Get useradd return error")
	}
	_, err = repo.Get(ID)
	if err != nil {
		t.Errorf("Test Repo Get return error")
	}
}

func TestUserGetError(t *testing.T) {
	ID := gofakeit.Int64()
	repo := memory.NewUserRepository()
	_, err := repo.Get(ID)
	if err == nil {
		t.Errorf("Test Repo Get return nil")
	}
}

func TestUserListSuccessfull(t *testing.T) {
	ID := gofakeit.Int64()
	User := &entity.User{ID: ID}
	repo := memory.NewUserRepository()
	err := repo.Add(ID, User)
	if err != nil {
		t.Errorf("Test Repo List useradd return error")
	}
	_, err = repo.List()
	if err != nil {
		t.Errorf("Test Repo List return error")
	}
}

func TestUserListError(t *testing.T) {
	repo := memory.NewUserRepository()
	_, err := repo.List()
	if err == nil {
		t.Errorf("Test Repo List return nil")
	}
}
