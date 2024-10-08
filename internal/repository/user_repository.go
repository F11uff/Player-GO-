package repository

import (
	"database/sql"
	"player/internal/storage/postgresql/user"
)

// НАДО СДЕЛАТЬ ТАК, ЧТОБЫ USERID НЕПРИХОДИЛ В ДАННУЮ ФУНКЦИЮ
// НАДО ЧТОБЫ ЗАПРОС ПРОИЗВОДИЛСЯ ЗДЕСЬ,
// И ОТДАВАЛСЯ В ПАКЕТ PKG В МЕТОД CreateJWTPayload

type POSTUserRep struct {
	DataBase *sql.DB
}

// Подключаться к бд надо здесь

func (rep *POSTUserRep) GetUserId(userID string) (user.User, error) {
	var user user.User

	sqlReq := `SELECT id, username FROM users WHERE user_id = ?`
}
