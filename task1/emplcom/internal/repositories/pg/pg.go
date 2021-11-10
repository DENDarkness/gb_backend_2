package pg

import (
	"database/sql"
)

type PG struct {
	db *sql.DB
}

func NewPG(db *sql.DB) *PG {
	return &PG{db: db}
}

