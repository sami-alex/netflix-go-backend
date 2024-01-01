package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sami-alex/netflix-go-backend/db"
)

func main() {
	router := gin.Default()
	client, err := db.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Disconnect(context.Background())
	router.Run(":3000")
}
