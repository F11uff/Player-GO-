package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	//// Пример пароля
	//password := "my_secure_password"
	//
	//// Хэшируем пароль
	//hashedPassword, err := hashPassword(password)
	//if err != nil {
	//	fmt.Println("Ошибка хэширования:", err)
	//	return
	//}
	//fmt.Println("Хэшированный пароль:", hashedPassword)
	//
	//// Проверяем пароль
	//if checkPassword(hashedPassword, password) {
	//	fmt.Println("Пароль верен!")
	//} else {
	//	fmt.Println("Неправильный пароль!")
	//}

	connStr := "host=db port=5432 user=user password=password dbname=mydb sslmode=disable"

	// Подключение к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Проверка пароля
func checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
