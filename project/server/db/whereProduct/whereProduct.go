package RoadProduct

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	WritingCart "WritingCart" // Importa el paquete WritingCart

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

// ObtenerProductosPorCarrito realiza la consulta a la base de datos y devuelve los resultados en formato JSON como un []Producto, filtrando por los IDs del carrito
func ObtenerProductosPorCarrito() ([]Producto, error) {
	// Obtener los IDs del carrito desde la API
	cartIDs := WritingCart.GetCartIDs()

	// Verificar si el carrito está vacío
	if len(cartIDs) == 0 {
		return nil, fmt.Errorf("el carrito está vacío")
	}

	// Llamar a la función para obtener los productos por IDs
	return ObtenerProductosPorIDs(cartIDs)
}

// ObtenerProductosPorIDs realiza la consulta a la base de datos y devuelve los resultados en formato JSON como un []Producto, filtrando por IDs específicos
func ObtenerProductosPorIDs(ids []int) ([]Producto, error) {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error al cargar el archivo .env: %w", err)
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
		return nil, fmt.Errorf("error al conectar a la base de datos: %w", err)
	}
	defer db.Close()

	// Construir la consulta SQL con los IDs específicos
	query := "SELECT id, name, href, imageSrc, description, price, color, quantity FROM products WHERE id IN ("
	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "?"
	}
	query += strings.Join(placeholders, ",") + ")"

	// Convertir los IDs de tipo int a tipo interface{}
	interfaceIDs := make([]interface{}, len(ids))
	for i, v := range ids {
		interfaceIDs[i] = v
	}

	// Ejecutar la consulta en la base de datos
	rows, err := db.Query(query, interfaceIDs...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %w", err)
	}
	defer rows.Close()

	// Crear un slice para almacenar los productos
	var productos []Producto

	// Iterar sobre las filas y almacenar los valores de cada columna en un Producto
	for rows.Next() {
		var producto Producto
		if err := rows.Scan(&producto.ID, &producto.Name, &producto.Href, &producto.ImageSrc, &producto.Description, &producto.Price, &producto.Color, &producto.Quantity); err != nil {
			return nil, fmt.Errorf("error al escanear fila: %w", err)
		}
		productos = append(productos, producto)
	}

	return productos, nil
}
