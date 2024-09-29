package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"player/internal/storage/postgresql/user"
)

func main() {
	connStr := "host=localhost port=5432 user=test password=12345 dbname=hw6 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Пример пользователя
	user := user.User{
		Username:     "Amir",
		Email:        "Amir.doe@example.com",
		HashPassword: "123456789",
	}

	// Добавление пользователя
	if err := user.AddUser(db, user); err != nil {
		log.Fatal("Error adding user:", err)
	}
}
