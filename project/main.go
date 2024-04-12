package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()
    app.Use(cors.New())
    
	app.Static("/", "./public/dist") 

    app.Get("/users", func(c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{
			"data": "Usuario del backend",
		})
    })

    app.Listen(":3000")
}