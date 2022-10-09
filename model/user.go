package model

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/majidehamide/go-fiber2/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Address string
}

func GetUsers(c *fiber.Ctx) error {
	var users []User

	db, _ := database.ConnectDB()

	db.Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	db, _ := database.ConnectDB()
	var user = new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON("Wrong format")
	}

	db.Create(&user)
	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	db, _ := database.ConnectDB()
	var user User

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON("Param not found")
	}

	db.First(&user, id)

	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON("Id not found")
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	db, _ := database.ConnectDB()
	var user User

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON("Param not found")
	}

	db.First(&user, id)

	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON("Id not found")
	}

	var update = new(User)
	if err := c.BodyParser(update); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON("Wrong format")
	}

	db.Model(&update).Where("id = ?", id).Updates(&update)

	return c.JSON("Success update")
}

func DeleteUser(c *fiber.Ctx) error {
	db, _ := database.ConnectDB()
	var user User

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON("Param not found")
	}

	db.First(&user, id)

	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON("Id not found")
	}

	db.Delete(&user)

	return c.JSON("Success delete")
}
