package service

import (
	"assignment_8/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRegisterUserExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	mockRepo.EXPECT().
		GetByEmail("test@mail.com").
		Return(&repository.User{}, nil)

	err := service.RegisterUser(&repository.User{}, "test@mail.com")

	assert.Error(t, err)
}

func TestRegisterUserSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	mockRepo.EXPECT().
		GetByEmail("new@mail.com").
		Return(nil, nil)

	mockRepo.EXPECT().
		CreateUser(gomock.Any()).
		Return(nil)

	err := service.RegisterUser(&repository.User{}, "new@mail.com")

	assert.NoError(t, err)
}

func TestUpdateUserName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	user := &repository.User{ID: 2, Name: "Old"}

	mockRepo.EXPECT().
		GetUserByID(2).
		Return(user, nil)

	mockRepo.EXPECT().
		UpdateUser(user).
		Return(nil)

	err := service.UpdateUserName(2, "New")

	assert.NoError(t, err)
	assert.Equal(t, "New", user.Name)
}

func TestDeleteAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	err := service.DeleteUser(1)

	assert.Error(t, err)
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	mockRepo.EXPECT().
		DeleteUser(5).
		Return(nil)

	err := service.DeleteUser(5)

	assert.NoError(t, err)
}
