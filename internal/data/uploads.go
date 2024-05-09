package data

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type Upload struct {
	ID        int64
	UserID    int64
	Timestamp time.Time
	FileName  string
	FileLabel string
	BlobName  string
}

type UploadModel struct {
	Conn *pgx.Conn
}

// TODO: Implement Get(), Insert(), Validate()
func (u UploadModel) Insert(ctx context.Context, tx pgx.Tx, upload *Upload) error {
	sql := `
	INSERT INTO user_uploads
		(user_id, timestamp, file_name, file_label, blob_name)
	VALUES ($1, $2, $3, $4, $5);
	`
	args := []any{upload.UserID, upload.Timestamp, upload.FileName, upload.FileLabel, upload.BlobName}

	_, err := tx.Exec(ctx, sql, args...)

	return err // TODO will this be nil?
}

func (u UploadModel) Get() (*Upload, error) {
	sql := `
		SELECT id, user_id, timestamp, file_name, file_label, blob_name
		FROM user_uploads
		WHERE id = $1;
	`
	rows, err := u.Conn.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err) // TODO is this right?
	}
	defer rows.Close()

	var upload Upload
	err = rows.Scan(&upload.ID, &upload.UserID, &upload.Timestamp, &upload.FileName, &upload.FileLabel, &upload.BlobName)

	return &upload, err
}
