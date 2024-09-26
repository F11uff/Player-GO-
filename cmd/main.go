package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Пример пароля
	password := "my_secure_password"

	// Хэшируем пароль
	hashedPassword, err := hashPassword(password)
	if err != nil {
		fmt.Println("Ошибка хэширования:", err)
		return
	}
	fmt.Println("Хэшированный пароль:", hashedPassword)

	// Проверяем пароль
	if checkPassword(hashedPassword, password) {
		fmt.Println("Пароль верен!")
	} else {
		fmt.Println("Неправильный пароль!")
	}
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
