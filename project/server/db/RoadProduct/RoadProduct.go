package RoadProduct

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Producto representa la estructura de un producto
type Producto struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Href        string  `json:"href"`
	ImageSrc    string  `json:"imageSrc"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Color       string  `json:"color"`
	Quantity    int     `json:"quantity"`
}

// ObtenerProductos realiza la consulta a la base de datos y devuelve los resultados en formato JSON como un string
func ObtenerProduct() (string, error) {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("error al cargar el archivo .env: %w", err)
	}

	// Obtener las credenciales de la base de datos desde las variables de entorno
	dbUser := os.Getenv("DbUser")
	dbPass := os.Getenv("DbPass")
	dbHost := os.Getenv("DbHost")
	dbName := os.Getenv("DbName")

	// Crea la cadena de conexión
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	// Conéctate a la base de datos
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return "", fmt.Errorf("error al conectar a la base de datos: %w", err)
	}
	defer db.Close()

	// Consulta en la base de datos
	rows, err := db.Query("SELECT id, name, href, imageSrc, description, price, color, quantity FROM products")
	if err != nil {
		return "", fmt.Errorf("error al ejecutar la consulta: %w", err)
	}
	defer rows.Close()

	// Crear un slice para almacenar los productos
	var productos []Producto

	// Iterar sobre las filas y almacenar los valores de cada columna en un Producto
	for rows.Next() {
		var producto Producto
		if err := rows.Scan(&producto.ID, &producto.Name, &producto.Href, &producto.ImageSrc, &producto.Description, &producto.Price, &producto.Color, &producto.Quantity); err != nil {
			return "", fmt.Errorf("error al escanear fila: %w", err)
		}
		productos = append(productos, producto)
	}

	// Convertir el slice de productos a formato JSON como string
	productosJSON, err := json.Marshal(productos)
	if err != nil {
		return "", fmt.Errorf("error al convertir a JSON: %w", err)
	}

	// Convertir el JSON byte slice a string
	productosString := string(productosJSON)

	return productosString, nil
}
