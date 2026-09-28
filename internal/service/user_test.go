package service_test

import (
	"testing"

	"cms/internal/entity"
	"cms/internal/usecase/mocks"

	"go.uber.org/mock/gomock"
)

func TestAddUserSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockUserRepository(ctrl)
	var UserID int64
	UserID = 123
	mockRepo.EXPECT().Add(UserID, &entity.User{ID: 111}).Return(nil)
	err := mockRepo.Add(123, &entity.User{ID: 111})
	if err != nil {
		t.Errorf("error unexpected: %v", err)
	}
}
