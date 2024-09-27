package user

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           uint   `gorm:"primary_key"`
	Username     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	HashPassword string `gorm:"unique;not null"`
}

func (u *User) AddUser(db *sql.DB, user User) error {
	HashPassword, err := HashPassword(user.HashPassword)

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
