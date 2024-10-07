package user

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"userLogin"`
	Password string `json:"userPassword"`
	Email    string `json:"userEmail"`
	Remember bool   `json:"userRemember"`
}

//	UserLogin    string `json:"userLogin"`    // Логин пользователя
//	UserPassword string `json:"userPassword"` // Пароль пользователя
//	UserEmail    string `json:"userEmail"`
//	UserRemember bool   `json:"userRemember"`

func (u *User) AddUser(user User) error {
	connStr := "host=localhost port=5432 user=test password=12345 dbname=hw6 sslmode=disable"
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

func HashPassword(pass string) (string, error) {
	HashPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	return string(HashPassword), err
}
