package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	_ "github.com/jackc/pgx/v5"
)

func main() {
	r := gin.Default()
	r.GET("/")

	fmt.Println("Hi")
}
