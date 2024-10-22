//go:generate mockgen -source=mocks.go -destination=mocksAunt/mocksAunt.go -package=mocks

package testAunt

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Model interface {
	GetHashPassword(username string) (string, error)
}

type UserLoginMock struct {
	Username string `json:"userLogin"`
	Password string `json:"userPassword"`
}

type UserLoginModel struct {
	Model Model
}

func (m *UserLoginModel) AuthenticateUserMock(user UserLoginMock) (string, error) {
	hashPassword, err := m.Model.GetHashPassword(user.Username)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(user.Password)); err != nil {
		return "", errors.New("wrong password")
	}

	return CreateJWTTokenForMocks(), nil
}

func CreateJWTTokenForMocks() string {

	return "NewTokenForMocks"
}
