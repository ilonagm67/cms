package service_test

import (
	"errors"
	"testing"

	"cms/internal/entity"
	"cms/internal/usecase/mocks"

	"go.uber.org/mock/gomock"
)

func TestUserDeleteSuccessfully(t *testing.T) {
	var UserID int64
	UserID = 123

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().Delete(UserID).Return(nil)
	err := mockRepo.Delete(123)
	if err != nil {
		t.Errorf("error unexpected: %v", err)
	}
}

func TestUserDeleteError(t *testing.T) {
	var UserID int64
	UserID = 0

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().Delete(UserID).Return(errors.New("User not Found"))
	err := mockRepo.Delete(0)
	if err == nil {
		t.Errorf("error unexpected: %v", err)
	}
}

func TestUserGetSuccessfully(t *testing.T) {
	var UserID int64
	UserID = 123

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().Get(UserID).Return(&entity.User{ID: UserID}, nil)
	user, err := mockRepo.Get(123)
	if err != nil {
		t.Errorf("error unexpected: %v", err)
	}
	if user.ID != UserID {
		t.Errorf("error unexpected: %v", err)
	}
}

func TestUserGetError(t *testing.T) {
	var UserID int64
	UserID = 0

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().Get(UserID).Return(nil, errors.New("User not Found"))
	user, err := mockRepo.Get(0)
	if user != nil {
		t.Errorf("error unexpected: %v", err)
	}
	if err == nil {
		t.Errorf("error unexpected: %v", err)
	}
}
