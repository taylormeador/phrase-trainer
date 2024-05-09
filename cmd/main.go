package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/jackc/pgx/v5"
	"phrasetrainer.tm.com/internal/data"
)

func main() {
	// Create a logger.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Open the DB connection.
	conn, err := openConn()
	if err != nil {
		logger.Error(err.Error())
	}
	defer conn.Close(context.Background())
	logger.Info("Successfully connected to db")

	// Create an Upload.
	upload := &data.Upload{ID: 0, UserID: 1, Timestamp: time.Now(), FileName: "file name", FileLabel: "file label", BlobName: "asdfwf-asdfawef-asdfawef-awsefasef"}

	// // Insert the Upload.
	// models := data.NewModels(conn)
	// err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
	// 	return models.Uploads.Insert(context.Background(), tx, upload)
	// })
	// if err != nil {
	// 	logger.Error(err.Error())
	// }

	// logger.Info("Insert was successful")

	models := data.NewModels(conn)

	// Get a known upload by ID.
	var uploadID int64 = 967147961206865921
	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		upload, err = models.Uploads.Get(context.Background(), tx, uploadID)
		return err
	})
	if err != nil {
		logger.Error(err.Error())
	}

	fmt.Printf("%+v\n", upload)
}

func openConn() (*pgx.Conn, error) {
	config, err := pgx.ParseConfig(os.Getenv("COCKROACH_DB_DSN"))
	if err != nil {
		return nil, err
	}

	config.RuntimeParams["phrase_trainer"] = "$ docs_simplecrud_gopgx"

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
