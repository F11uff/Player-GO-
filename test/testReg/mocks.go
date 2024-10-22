//go:generate mockgen -source=mocks.go -destination=mockReg/mocksReg.go -package=mocks

package testReg

import (
	"database/sql"
	"errors"
)

type DBMocks interface {
	QueryRow(query string, args ...interface{}) (string, error) // Возвращает email и ошибку
}

type UserRegistrationMocks struct {
	Email string `json:"email"`
}

func (u *UserRegistrationMocks) FindUserForEmailMocks(db DBMocks) (bool, error) {
	sqlRequest := `SELECT Email FROM users WHERE Email=$1`
	email, err := db.QueryRow(sqlRequest, u.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return email == u.Email, nil
}
