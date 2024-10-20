package testAunt

import (
	"github.com/golang/mock/gomock"
	"player/internal/services"
	"player/test/testAunt/mocks"
	"testing"
)

func TestAunteticateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	//defer ctrl.Finish()
	mock := mocks.NewMockModel(ctrl)

	username := "test"
	password := "testpassword"

	hashedpassword, _ := services.HashPassword("testpassword")

	mock.EXPECT().GetHashPassword(username).Return(hashedpassword, nil)

	userLoginMock := UserLoginMock{
		Username: username,
		Password: password,
	}

	auth := UserLoginModel{Model: mock}

	token, err := auth.AuthenticateUserMock(userLoginMock)
	if err != nil {
		t.Fatalf("Error - %v", err)
	}

	if token == "" {
		t.Fatalf("Error - token is empty")
	}

}

func TestAunteticateUser_FailPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := mocks.NewMockModel(ctrl)

	username := "test"
	password := "testpasswordfail"
	hashedpassword, err := services.HashPassword("testpassword")

	mock.EXPECT().GetHashPassword(username).Return(hashedpassword, nil)

	userLoginMock := UserLoginMock{
		Username: username,
		Password: password,
	}

	auth := UserLoginModel{Model: mock}
	token, err := auth.AuthenticateUserMock(userLoginMock)

	if err == nil {
		t.Fatalf("Error - %v", err)
	}

	if err.Error() != "wrong password" {
		t.Fatalf("Expected error 'wrong password', got '%v'", err)
	}

	if token != "" {
		t.Fatalf("Error - %v", err)
	}

}
