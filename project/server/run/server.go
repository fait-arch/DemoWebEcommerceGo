package main

import (
	"fmt"
	"log"

	dbRoadProduct "RoadProduct"
	dbRoadPropiedades "RoadPropiedades"

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
        propiedadesJSON, err := dbRoadPropiedades.ObtenerPropiedades()
        if err != nil {
            log.Printf("Error al obtener propiedades: %v", err)
            return c.Status(fiber.StatusInternalServerError).SendString("Error al obtener propiedades")
        }
        // Devuelve el JSON como respuesta
        return c.JSON(propiedadesJSON)
    })

    // Ruta para manejar la solicitud y devolver el JSON de productos
    app.Get("/product", func(c fiber.Ctx) error {
        // Llama a la función ObtenerPropiedades para obtener el JSON de productos
        productJSON, err := dbRoadProduct.ObtenerProduct()
        if err != nil {
            log.Printf("Error al obtener propiedades: %v", err)
            return c.Status(fiber.StatusInternalServerError).SendString("Error al obtener propiedades")
        }
        // Devuelve el JSON como respuesta
        return c.JSON(productJSON)
    })

    // Lanza el servidor Fiber
    fmt.Println("Iniciando servidor Fiber...")
    log.Fatal(app.Listen(":3000"))
}
