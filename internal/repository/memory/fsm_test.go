package memory_test

import (
	"cms/internal/repository/memory"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestFSMSetSuccessfully(t *testing.T) {
	UserID := gofakeit.Int64()
	repo := memory.NewFSMRepository()
	err := repo.Set(UserID, "FSMData")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFSMSetError(t *testing.T) {
	UserID := gofakeit.Int64()
	repo := memory.NewFSMRepository()
	err := repo.Set(UserID, "")
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFSMSetNameSuccessfully(t *testing.T) {
	UserID := gofakeit.Int64()
	repo := memory.NewFSMRepository()
	err := repo.Set(UserID, "FSMData")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	err = repo.SetName(UserID, "FSMData")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFSMSetNameError(t *testing.T) {
	UserID := gofakeit.Int64()
	repo := memory.NewFSMRepository()
	err := repo.SetName(UserID, "FSMData")
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFSMGetNameSuccessfully(t *testing.T) {
	UserID := gofakeit.Int64()
	repo := memory.NewFSMRepository()
	err := repo.Set(UserID, "FSMData")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	err = repo.SetName(UserID, "FSMData")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	_, err = repo.GetName(UserID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFSMGetNameError(t *testing.T) {
	UserID := gofakeit.Int64()
	repo := memory.NewFSMRepository()
	_, err := repo.GetName(UserID)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFSMSetWeightSuccessfully(t *testing.T) {
	UserID := gofakeit.Int64()
	Weight := gofakeit.Number(1, 10)
	repo := memory.NewFSMRepository()
	err := repo.Set(UserID, "FSMData")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	err = repo.SetWeight(UserID, Weight)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFSMSetWeightError(t *testing.T) {
	UserID := gofakeit.Int64()
	Weight := gofakeit.Number(1, 10)
	repo := memory.NewFSMRepository()
	err := repo.SetWeight(UserID, Weight)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFSMGetWeightSuccessfully(t *testing.T) {
	UserID := gofakeit.Int64()
	Weight := gofakeit.Number(1, 10)
	repo := memory.NewFSMRepository()
	err := repo.Set(UserID, "FSMData")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	err = repo.SetWeight(UserID, Weight)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	_, err = repo.GetWeight(UserID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFSMGetWeightError(t *testing.T) {
	UserID := gofakeit.Int64()
	repo := memory.NewFSMRepository()
	_, err := repo.GetWeight(UserID)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFSMDeleteSuccessfully(t *testing.T) {
	UserID := gofakeit.Int64()
	repo := memory.NewFSMRepository()
	err := repo.Set(UserID, "FSMData")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	err = repo.Delete(UserID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFSMDeleteError(t *testing.T) {
	UserID := gofakeit.Int64()
	repo := memory.NewFSMRepository()
	err := repo.Delete(UserID)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFSMGetSuccessfully(t *testing.T) {
	UserID := gofakeit.Int64()
	repo := memory.NewFSMRepository()
	err := repo.Set(UserID, "FSMData")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	_, err = repo.Get(UserID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFSMGetError(t *testing.T) {
	UserID := gofakeit.Int64()
	repo := memory.NewFSMRepository()
	_, err := repo.Get(UserID)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}
