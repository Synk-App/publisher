package model

import (
	"database/sql"
)

type X struct {
	db *sql.DB
}

func NewX(db *sql.DB) *X {
	x := X{db: db}

	return &x
}
