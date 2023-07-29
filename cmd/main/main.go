package main

import (
	"boba/api"
	db "boba/db/sqlc"
	"context"
	"log"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgresql://root:123@localhost:5432/boba_shop?sslmode=disable"
	serverAdress = "0.0.0.0:8080"
)

func main() {

	conn, err := pgxpool.New(context.Background(), dbSource)

	if err != nil {
		log.Fatal("Cannot cnnect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
