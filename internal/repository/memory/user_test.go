package memory_test

import (
	"cms/internal/entity"
	"cms/internal/repository/memory"
	"testing"
)

func TestUserAddSuccessfull(t *testing.T) {
	var ID int64 = 10454545
	var User = "test"
	repo := memory.NewUserRepository()
	err := repo.Add(ID, &entity.User{Name: User})
	if err != nil {
		t.Errorf("Test return error")
	}
}
