package user

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"player/internal/services"
	"player/internal/storage/postgresql"
	"player/pkg/security"
)

var db *sql.DB

type UserRegistration struct {
	Username string `json:"userLogin"`
	Password string `json:"userPassword"`
	Email    string `json:"userEmail"`
	Remember bool   `json:"userRemember"`
}

type UserLogin struct {
	Username string `json:"userLogin"`
	Password string `json:"userPassword"`
}

func (u *UserRegistration) AddUser(user UserRegistration) error {
	var err error
	db, err = postgresql.OpenDB()

	if err != nil {
		return err
	}

	defer db.Close()

	exists, err := user.FindUserForEmail(user, db)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("A user with this email already exists")
	}

	HashPassword, err := services.HashPassword(user.Password)

	if err != nil {
		return err
	}

	sqlRequest := `INSERT INTO users (Username, Email, HashPassword) VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlRequest, user.Username, user.Email, HashPassword)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserLogin) AuthenticateUser(user UserLogin) (string, error) {
	var err error
	db, err = postgresql.OpenDB()

	if err != nil {
		return "", err
	}

	defer db.Close()

	sqlRequest := `SELECT hashpassword FROM users WHERE Username=$1`
	rows, err := db.Query(sqlRequest, user.Username)

	defer rows.Close()

	if err != nil {
		return "", err
	}

	var hashPassword string

	for rows.Next() {
		if err = rows.Scan(&hashPassword); err != nil {
			return "", err
		}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(user.Password)); err != nil {
		return "", errors.New("Invalid username or password")
	}

	return security.CreateJWTToken(user.Password, user.Username), nil
}

func (u *UserRegistration) FindUserForEmail(userStructReg UserRegistration, db2 *sql.DB) (bool, error) {
	var err error
	var email string

	sqlRequest := `SELECT Email FROM users WHERE Email=$1`
	err = db2.QueryRow(sqlRequest, userStructReg.Email).Scan(&email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
