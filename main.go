package main

import (
	"fmt"

	"example.com/go_api/db"
	"example.com/go_api/server"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	err := server.SetupRoutes(gin.Default())
	if err != nil {
		fmt.Println("Failed to setup routes due to : " + err.Error())
	}
}
