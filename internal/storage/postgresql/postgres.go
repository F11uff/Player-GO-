package postgresql

import (
	"database/sql"
	"fmt"
	"player/internal/config"
)

func OpenDB() (*sql.DB, error) {
	cnf := config.DefaultConfig()
	connStr := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=%s",
		cnf.DBConfig.Port, cnf.DBConfig.User, cnf.DBConfig.Password, cnf.DBConfig.DBName, cnf.DBConfig.SslMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
