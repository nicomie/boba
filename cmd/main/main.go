package main

import (
	"boba/api"
	db "boba/db/sqlc"
	"boba/util"
	"context"
	"log"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load conifg")
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("Cannot cnnect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
