package main

import (
	"fmt"
	"log"
	"strconv"

	dbRoadProduct "RoadProduct"
	dbRoadPropiedades "RoadPropiedades"
	whereProduct "whereProduct"

	WritingCart "WritingCart"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
    // Inicializa el servidor de Fiber
    app := fiber.New()
    app.Use(cors.New())
    app.Use(logger.New())

    // Ruta POST para guardar un string
    app.Post("/addrocart", WritingCart.AddToCartHandler) // Asigna el handler AddToCartHandler para la ruta POST /cart
    app.Post("/clearcart", WritingCart.ClearCartHandler)

    // Agrega un nuevo endpoint para obtener el contenido del carrito
    app.Get("/getCartContent", func(c fiber.Ctx) error {
        // Obtén el contenido del carrito desde WritingCart de manera concurrente
        cartIDsChan := make(chan []string)
        go func() {
            // Obtén los IDs del carrito de compras
            cartIDsInt := WritingCart.GetCartIDs()
            // Convierte los IDs a strings
            cartIDs := make([]string, len(cartIDsInt))
            for i, id := range cartIDsInt {
                cartIDs[i] = strconv.Itoa(id)
            }
            cartIDsChan <- cartIDs
        }()

    // Obtén los IDs del carrito de compras
    cartIDs := <-cartIDsChan

    // Devuelve el contenido del carrito como respuesta en formato JSON
    return c.JSON(cartIDs)
})

    // Agrega un nuevo endpoint para obtener el contenido del productos segun el ID
    app.Get("/productID", func(c fiber.Ctx) error {
        // Obtener los IDs del carrito de compras de manera concurrente
        cartIDsChan := make(chan []string)
        go func() {
            // Obtén los IDs del carrito de compras
            cartIDsInt := WritingCart.GetCartIDs()
            // Convierte los IDs a strings
            cartIDs := make([]string, len(cartIDsInt))
            for i, id := range cartIDsInt {
                cartIDs[i] = strconv.Itoa(id)
            }
            cartIDsChan <- cartIDs
        }()

        // Obtén los IDs del carrito de compras
        cartIDs := <-cartIDsChan

        // Convierte los IDs del carrito de strings a enteros
        cartIDsInt := make([]int, len(cartIDs))
        for i, idStr := range cartIDs {
            idInt, err := strconv.Atoi(idStr)
            if err != nil {
                log.Printf("Error al convertir ID a entero: %v", err)
                // Puedes manejar el error aquí según tus necesidades
                return c.Status(fiber.StatusInternalServerError).SendString("Error al convertir ID a entero")
            }
            cartIDsInt[i] = idInt
        }

        // Llama a la función ObtenerProductosPorIDs para obtener los productos por esos IDs de manera concurrente
        productIDJSONChan := make(chan interface{})
        go func() {
            productIDJSON, err := whereProduct.ObtenerProductosPorIDs(cartIDsInt)
            if err != nil {
                log.Printf("Error al obtener productos: %v", err)
                productIDJSONChan <- err
                return
            }
            productIDJSONChan <- productIDJSON
        }()

        // Obtén el JSON de productos de manera concurrente
        productIDJSON := <-productIDJSONChan

        // Devuelve el JSON como respuesta
        return c.JSON(productIDJSON)
    })

    // Ruta para manejar la solicitud y devolver el JSON de propiedades
    app.Get("/propiedades", func(c fiber.Ctx) error {
        // Llama a la función ObtenerPropiedades para obtener el JSON de propiedades de manera concurrente
        propiedadesJSONChan := make(chan interface{})
        go func() {
            propiedadesJSON, err := dbRoadPropiedades.ObtenerPropiedades()
            if err != nil {
                log.Printf("Error al obtener propiedades: %v", err)
                propiedadesJSONChan <- err
                return
            }
            propiedadesJSONChan <- propiedadesJSON
        }()

        // Obtén el JSON de propiedades de manera concurrente
        propiedadesJSON := <-propiedadesJSONChan

        // Devuelve el JSON como respuesta
        return c.JSON(propiedadesJSON)
    })

    // Ruta para manejar la solicitud y devolver el JSON de productos
    app.Get("/product", func(c fiber.Ctx) error {
        // Llama a la función ObtenerProduct para obtener el JSON de productos de manera concurrente
        productJSONChan := make(chan interface{})
        go func() {
            productJSON, err := dbRoadProduct.ObtenerProduct()
            if err != nil {
                log.Printf("Error al obtener productos: %v", err)
                productJSONChan <- err
                return
            }
            productJSONChan <- productJSON
        }()

        // Obtén el JSON de productos de manera concurrente
        productJSON := <-productJSONChan

        // Devuelve el JSON como respuesta
        return c.JSON(productJSON)
    })

    // Lanza el servidor Fiber
    fmt.Println("Iniciando servidor Fiber...")
    log.Fatal(app.Listen(":3000"))
}
