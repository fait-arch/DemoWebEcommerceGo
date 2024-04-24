package WritingCart

import (
	"github.com/gofiber/fiber/v3"
)

// SaveToCart is the handler to save a string to the cart
func SaveToCart(c fiber.Ctx) error {
    // Struct to store the text
    type Request struct {
        Text string `json:"text"`
    }

    // Parse JSON body into the request struct
    var request Request
    if err := c.BodyParser(&request); err != nil {
        return err
    }

    // Here you can save the string to your database or perform any other action
    // In this example, we simply return it in the response
    return c.JSON(fiber.Map{"message": "Text saved to cart", "text": request.Text})
}
