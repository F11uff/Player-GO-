package testAunt

import (
	"github.com/golang/mock/gomock"
	"player/test/testAunt/mocks"
	"testing"
)

func TestAunteticateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	//defer ctrl.Finish()
	mock := mocks.NewMockModel(ctrl)

	username := "test"
	password := "testpassword"
	hashedpassword := "$2a$10$e0MYW8dO7h7Juj.7yC1wV.fzG5NQRC8W8yGnD6bgAvfJfw0m8DQSy"

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
