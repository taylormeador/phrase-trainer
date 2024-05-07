package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func Connect() *pgx.Conn {
	dsn := os.Getenv("COCKROACH_DSN")
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	var now time.Time
	err = conn.QueryRow(ctx, "SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	fmt.Println("Successfully connected to db", now)

	return conn
}

func InitTable(ctx context.Context, tx pgx.Tx) error {
	if _, err := tx.Exec(ctx, userUploadsDDL); err != nil {
		return err
	}
	return nil
}

func InsertRow(ctx context.Context, tx pgx.Tx, userID int, timestamp time.Time, fileLabel string, blobName string) error {
	sql := `
		INSERT INTO user_uploads (user_id, timestamp, file_label, blob_name)
		VALUES ($1, $2, $3, $4);
	`

	if _, err := tx.Exec(ctx, sql, userID, timestamp, fileLabel, blobName); err != nil {
		return err
	}
	return nil
}
