package user

import (
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"player/internal/config"
	"player/pkg/security"
)

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
	cnf := config.DefaultConfig()

	connStr := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=%s",
		cnf.DBConfig.Port, cnf.DBConfig.User, cnf.DBConfig.Password, cnf.DBConfig.DBName, cnf.DBConfig.SslMode)
	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		return err
	}

	HashPassword, err := HashPassword(user.Password)

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
	cnf := config.DefaultConfig()

	connStr := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=%s",
		cnf.DBConfig.Port, cnf.DBConfig.User, cnf.DBConfig.Password, cnf.DBConfig.DBName, cnf.DBConfig.SslMode)
	db, err := sql.Open("postgres", connStr)

	err = db.Ping()

	defer db.Close()

	if err != nil {
		return "", err
	}

	sqlRequest := `SELECT hashpassword FROM users WHERE Username=$1`
	rows, err := db.Query(sqlRequest, user.Username)

	defer rows.Close()

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

func HashPassword(pass string) (string, error) {
	HashPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	return string(HashPassword), err
}
