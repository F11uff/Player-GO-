package repository

import (
	"database/sql"
	"fmt"
	"player/internal/config"
	"player/internal/storage/postgresql/user"
)

// НАДО СДЕЛАТЬ ТАК, ЧТОБЫ USERID НЕПРИХОДИЛ В ДАННУЮ ФУНКЦИЮ
// НАДО ЧТОБЫ ЗАПРОС ПРОИЗВОДИЛСЯ ЗДЕСЬ,
// И ОТДАВАЛСЯ В ПАКЕТ PKG В МЕТОД CreateJWTPayload

type POSTUserRep struct {
	DataBase *sql.DB
}

// Подключаться к бд надо здесь
func (rep *POSTUserRep) GetUserId(userID string) (user.UserRegistration, error) {
	var user user.UserRegistration
	cnf := config.DefaultConfig()

	connStr := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=%s",
		cnf.DBConfig.Port, cnf.DBConfig.User, cnf.DBConfig.Password, cnf.DBConfig.DBName, cnf.DBConfig.SslMode)

	rep.DataBase, err = sql.Open("postgres", connStr)

	sqlReq := `SELECT id, username FROM users WHERE user_id = ?`

}
