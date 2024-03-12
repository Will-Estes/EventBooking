package main

import (
	"eventbooking/db"
	"eventbooking/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
