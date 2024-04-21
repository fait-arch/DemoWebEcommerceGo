package main

import (
	"fmt"
	"log"

	db "db"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
    // Inicializa el servidor de Fiber
    app := fiber.New()
    app.Use(cors.New())
    app.Use(logger.New())
    app.Static("", "../../client/dist")

    // Ruta para manejar la solicitud y devolver el JSON de propiedades
    app.Get("/propiedades", func(c fiber.Ctx) error {
        // Llama a la función ObtenerPropiedades para obtener el JSON de propiedades
        propiedadesJSON, err := db.ObtenerPropiedades()
        if err != nil {
            log.Printf("Error al obtener propiedades: %v", err)
            return c.Status(fiber.StatusInternalServerError).SendString("Error al obtener propiedades")
        }
        // Devuelve el JSON como respuesta
        return c.JSON(propiedadesJSON)
    })

    // Lanza el servidor Fiber
    fmt.Println("Iniciando servidor Fiber...")
    log.Fatal(app.Listen(":3000"))
}
