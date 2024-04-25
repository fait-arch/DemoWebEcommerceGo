package WritingCart

import (
	"fmt"
	"strconv"
	"strings"
	"sync" // Importar el paquete sync para la sincronización segura

	"github.com/gofiber/fiber/v3"
)

var (
    cartIDs    []int      // Slice para almacenar los IDs de los elementos del carrito
    cartString string     // Variable global para almacenar el string del carrito
    mutex      sync.Mutex // Mutex para garantizar acceso seguro a las variables globales en entornos concurrentes
)

// ClearCartHandler maneja las solicitudes POST a "/clearCart" en Fiber.
func ClearCartHandler(c fiber.Ctx) error {
    // Obtener el contenido del carrito como un JSON antes de limpiarlo
    cartContent := GetCartContentJSON()

    // Llamar a la función para limpiar el carrito
    ClearCartIDs()

    // Responder al cliente con el contenido del carrito en formato JSON
    return c.JSON(cartContent)
}

// AddToCartHandler maneja las solicitudes POST a "/addToCart" en Fiber.
func AddToCartHandler(c fiber.Ctx) error {
    // Leer el cuerpo de la solicitud
    body := c.Body()

    // Convertir el slice de bytes a un string
    bodyString := string(body)

    // Bloquear el mutex para acceso seguro a las variables globales
    mutex.Lock()
    defer mutex.Unlock()

    // Guardar el string en la variable global
    cartString = bodyString

    // Parsear los IDs de los productos enviados como una cadena separada por comas
    idStrings := strings.Split(bodyString, ",")
    for _, idStr := range idStrings {
        id, err := strconv.Atoi(idStr)
        if err != nil {
            fmt.Println("Error al convertir ID a entero:", err)
            continue
        }
        // Verificar si el ID ya está en la lista, si no, agregarlo
        if !contains(cartIDs, id) {
            cartIDs = append(cartIDs, id)
        }
    }

    // Imprimir la lista de IDs del carrito
    fmt.Println("Lista de IDs del carrito:", cartIDs)

    // Responder al cliente con un mensaje simple
    return c.SendString("String recibido y guardado correctamente en el carrito")
}

// ClearCartIDs vacía la lista de IDs del carrito
func ClearCartIDs() {
    // Bloquear el mutex para acceso seguro a la variable global
    mutex.Lock()
    defer mutex.Unlock()

    // Vaciar la lista de IDs del carrito
    cartIDs = nil
}

// GetCartContentJSON devuelve el contenido del carrito como un JSON antes de limpiarlo
func GetCartContentJSON() map[string]interface{} {
    // Bloquear el mutex para acceso seguro a la variable global
    mutex.Lock()
    defer mutex.Unlock()

    // Crear un mapa para almacenar el contenido del carrito
    cartContent := make(map[string]interface{})

    // Agregar el contenido del carrito al mapa
    cartContent["ids"] = cartIDs

    return cartContent
}

// GetCartString devuelve el string guardado en el carrito
func GetCartString() string {
    // Bloquear el mutex para acceso seguro a la variable global
    mutex.Lock()
    defer mutex.Unlock()

    // Devolver el string guardado en el carrito
    return cartString
}

// GetCartIDs devuelve la lista de IDs del carrito
func GetCartIDs() []int {
    // Bloquear el mutex para acceso seguro a la variable global
    mutex.Lock()
    defer mutex.Unlock()

    // Devolver una copia de la lista de IDs del carrito para evitar modificaciones inesperadas
    idsCopy := make([]int, len(cartIDs))
    copy(idsCopy, cartIDs)
    return idsCopy
}

// contains verifica si un slice contiene un valor específico
func contains(slice []int, value int) bool {
    for _, item := range slice {
        if item == value {
            return true
        }
    }
    return false
}
