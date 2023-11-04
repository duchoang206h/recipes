package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/handler"
)

func SetUpRoutes(app *fiber.App) {
	api := app.Group("/api")
	book := api.Group("/books")
	book.Get("/", handler.AllBooks)
	book.Post("/", handler.AddBook)
	book.Patch("/:bookID", handler.UpdateBook)
	book.Delete("/:bookID", handler.Delete)
}
