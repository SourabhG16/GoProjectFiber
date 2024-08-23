package main

import (
	"log"

	"github.com/SourabhG16/goProject1/database"
	"github.com/SourabhG16/goProject1/database/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Fiber Tutorial!")
}
func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)

	//usere routes
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/getUsers", routes.GetUsers)
	app.Get("/api/findUser/:id", routes.FindUser)
	app.Post("/api/UpdateUser/:id", routes.UpdateUser)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
