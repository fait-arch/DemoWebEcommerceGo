package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
    app := fiber.New()
    app.Use(cors.New())
    // Initialize default config
	app.Use(logger.New())
	
	app.Static("", "../client/dist") 

	app.Get("/users", func(c fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "Usuario del backend",
		})
	})


    app.Listen(":3000")
}

