package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:123@localhost:5432/boba_shop?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgxpool.New(context.Background(), dbSource)

	if err != nil {
		log.Fatal("Cannot cnnect to db: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
