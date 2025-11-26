package model

import (
	"database/sql"
)

type Mastodon struct {
	db *sql.DB
}

func NewMastodon(db *sql.DB) *Mastodon {
	mastodon := Mastodon{db: db}

	return &mastodon
}
