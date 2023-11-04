package handler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
	"github.com/zeimedee/go-postgres/models"
	"gorm.io/gorm"
)

// AddBook
func AddBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&book)

	return c.Status(200).JSON(book)
}

// AllBooks
func AllBooks(c *fiber.Ctx) error {
	books := []models.Book{}
	title := c.Query("title", "")
	query := database.DB.Db.Select("id", "title", "author")
	if title != "" {
		query.Where("title like ?", "%"+title+"%")
	}
	query.Find(&books)
	return c.Status(200).JSON(fiber.Map{"result": books})
}

// Update
func UpdateBook(c *fiber.Ctx) error {
	bookID, _ := strconv.Atoi(c.Params("bookID"))
	book := models.Book{}
	result := database.DB.Db.First(&book, bookID)
	fmt.Println("book", book)
	fmt.Println("bookID", book)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"message": "Book not found",
		})
	}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result = database.DB.Db.Save(book)
	if result.Error != nil {
		return c.Status(400).JSON(result.Error)
	}
	return c.Status(400).JSON("updated")
}

// Delete
func Delete(c *fiber.Ctx) error {
	bookID, _ := strconv.Atoi(c.Params("bookID"))
	book := new(models.Book)
	query := database.DB.Db.First(book, "id = ?", bookID)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"message": "Book not found",
		})
	}
	database.DB.Db.Delete(book)

	return c.Status(200).JSON("deleted")
}
