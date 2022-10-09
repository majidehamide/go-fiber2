package router

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/majidehamide/go-fiber2/model"
)

func Routes() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is API")
	})

	user := app.Group("user")

	user.Get("/", model.GetUsers)
	user.Post("/", model.CreateUser)
	user.Get("/:id", model.GetUser)
	user.Put("/:id", model.UpdateUser)
	user.Delete("/:id", model.DeleteUser)
	app.Listen(":3000")
}
