package controller

import (
	"database/sql"
	"net/http"
	"synk/gateway/app/model"
)

type X struct {
	model *model.X
}

type HandleXAuthResponse struct {
	Resource ResponseHeader        `json:"resource"`
	Data     UserXAuthDataResponse `json:"x"`
}

type UserXAuthDataResponse struct {
	Token string `json:"token"`
}

func NewX(db *sql.DB) *X {
	x := X{
		model: model.NewX(db),
	}

	return &x
}

func (u *X) HandleAuth(w http.ResponseWriter, r *http.Request) {
	SetJsonContentType(w)

	response := HandleXAuthResponse{
		Resource: ResponseHeader{
			Ok: true,
		},
		Data: UserXAuthDataResponse{},
	}

	WriteSuccessResponse(w, response)
}
