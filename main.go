package main

import (
	"axion/database"
	"axion/database/migration"
	"log"
	"os"

	_ "axion/docs"
	route "axion/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Auction API Documentation
// @version 1.0
// @description This is API documentation for Auction project
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email yahyatruth@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host axion-be-production.up.railway.app
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	database.DatabaseInit()
	migration.RunMigration()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization,Access-Control-Allow-Headers,Headers",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	route.RouteInit(app)
	port := "8080"

	errListen := app.Listen(":" + port)
	if errListen != nil {
		log.Println("Fail to listen go fiber server")
		os.Exit(1)
	}
}
