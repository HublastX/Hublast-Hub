package main

import (
	"github.com/HublastX/HubLast-Hub/internal/api/routes"
	"github.com/HublastX/HubLast-Hub/pkg/database"
	"github.com/gin-gonic/gin"

	_ "github.com/HublastX/HubLast-Hub/docs"
)

// @title           HubLast-Hub API
// @version         1.0
// @description     Project Management System API

// @contact.name   API Support
// @contact.email  support@hublast.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5000
// @BasePath  /api
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	database.Connect()

	router := gin.Default()

	routes.SetupRoutes(router)

	database.SeedTechnologies(database.DB)

	router.Run(":5000")
}
