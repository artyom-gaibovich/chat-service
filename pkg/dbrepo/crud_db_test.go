package dbrepo

import (
	"chat-service/pkg/env"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestConnect(t *testing.T) {

	ctx := context.Background()
	dbDSN := env.GetString("DB_DSN", "host=localhost port=54321 dbname=postgres user=admin password=root sslmode=disable")

	con, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer con.Close(ctx)
	t.Run("It should be connected", func(t *testing.T) {
		assert.Equal(t, true, con.Ping(ctx))
	})

}

func TestBase_Exec(t *testing.T) {

	ctx := context.Background()
	dbDSN := env.GetString("DB_DSN", "host=localhost port=12345 dbname=postgres user=admin password=root sslmode=disable")

	con, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer con.Close(ctx)

	base := NewBase()
	rows, err := con.Query(ctx, `SELECT id from "user".users`)
	if err != nil {
		log.Fatalf("failed to select notes: %v", err)
	}
	defer rows.Close()

}
