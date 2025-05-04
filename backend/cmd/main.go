package main

import (
	"github.com/HublastX/HubLast-Hub/internal/api/routes"
	"github.com/HublastX/HubLast-Hub/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":3005")
}
