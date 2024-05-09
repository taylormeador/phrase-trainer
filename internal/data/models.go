package data

import (
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Uploads UploadModel
}

func NewModels(conn *pgx.Conn) Models {
	return Models{
		Uploads: UploadModel{Conn: conn},
	}
}
