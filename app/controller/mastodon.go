package controller

import (
	"database/sql"
	"synk/gateway/app/model"
)

type Mastodon struct {
	model *model.Mastodon
}

func NewMastodon(db *sql.DB) *Mastodon {
	mastodon := Mastodon{
		model: model.NewMastodon(db),
	}

	return &mastodon
}
