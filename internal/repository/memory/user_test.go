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
		t.Errorf("Test return error")
	}
}

func TestUserAddError(t *testing.T) {
	ID := gofakeit.Int64()
	repo := memory.NewUserRepository()
	err := repo.Add(ID, &entity.User{})
	if err == nil {
		t.Errorf("Test return nil")
	}
}
