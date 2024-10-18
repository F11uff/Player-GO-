package testReg

import "database/sql"

type Model interface {
	AddUser(string, string) error
}

type PGRepository struct {
	db *sql.DB
}
