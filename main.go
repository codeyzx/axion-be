package main

import (
	"axion/database"
	"axion/database/migration"
	"log"
	"os"

	_ "axion/docs"
	route "axion/routes"

	"github.com/gofiber/fiber/v2"
)

// @title Auction API Documentation
// @version 1.0
// @description This is API documentation for Auction project
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email yahyatruth@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	database.DatabaseInit()
	migration.RunMigration()
	app := fiber.New()
	route.RouteInit(app)
	port := "8080"

	errListen := app.Listen(":" + port)
	if errListen != nil {
		log.Println("Fail to listen go fiber server")
		os.Exit(1)
	}
}
