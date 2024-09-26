package postgresql

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func New(dataBasePath string) (*Storage, error) {
	const errorMessage = "postgres.postgres_go.New"

	db, err := sql.Open("postgres", dataBasePath)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", errorMessage, err)
	}
}
