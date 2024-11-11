package main

import (
	"chat-service/pkg/env"
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4"
	"log"
)

func main() {
	ctx := context.Background()
	dbDSN := env.GetString("DB_DSN", "host=localhost port=12345 dbname=postgres user=admin password=root sslmode=disable")

	con, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer con.Close(ctx)

	res, err := con.Exec(ctx, `INSERT INTO "user".users (email, password, name) VALUES ($1, $2, $3)`, gofakeit.City(), gofakeit.Address().Street, gofakeit.Name())
	if err != nil {
		log.Fatalf("failed to insert note: %v", err)
	}

	log.Printf("inserted %d rows", res.RowsAffected())

	rows, err := con.Query(ctx, `SELECT id from "user".users`)
	if err != nil {
		log.Fatalf("failed to select notes: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int

		err = rows.Scan(&id)
		if err != nil {
			log.Fatalf("failed to scan note: %v", err)
		}

		log.Printf("id: %d, title: %s, body: %s, created_at: %v, updated_at: %v\n", id)
	}
}
