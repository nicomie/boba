package db

import (
	"boba/util"
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load conifg")
	}

	testDB, err = pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("Cannot cnnect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
