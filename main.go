package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/config"
	"github.com/habbazettt/jobseek-go/routes"
)

func main() {
	config.ConnectDB()

	config.SetupCloudinary()

	r := gin.Default()

	routes.AuthRoutes(r)

	port := "8080"
	fmt.Println("Server running on port " + port)
	log.Fatal(r.Run(":" + port))
}
