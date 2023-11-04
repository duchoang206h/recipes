package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zeimedee/go-postgres/database"
	"github.com/zeimedee/go-postgres/routes"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	routes.SetUpRoutes(app)

	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
