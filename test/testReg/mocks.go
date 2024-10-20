//go:generate mockgen -source=mocks.go -destination=mocks/mocksReg.go -package=mocks

package testReg

import (
	"database/sql"
	"errors"
)

type DBMocks interface {
	QueryRow(query string, args ...interface{}) RowMocks
}

type RowMocks interface {
	Scan(dest ...interface{}) error
}

type UserRegistrationMocks struct {
	Email string `json:"email"`
}

func (u *UserRegistrationMocks) FindUserForEmailMocks(db DBMocks) (bool, error) {
	var email string
	sqlRequest := `SELECT Email FROM users WHERE Email=$1`
	err := db.QueryRow(sqlRequest, u.Email).Scan(&email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
