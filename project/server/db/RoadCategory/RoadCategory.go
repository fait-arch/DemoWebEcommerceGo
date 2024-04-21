package RoadCategory

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Category representa la estructura de un resultado de consulta
type Category struct {
	CategoryID int    `json:"Category_id"`
	Tipo        string `json:"tipo"`
}

// ObtenerCategoryesString realiza la consulta a la base de datos y devuelve los resultados en formato JSON como un string
func ObtenerCategoryes() (string, error) {
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
	rows, err := db.Query("SELECT CategoryID, Tipo FROM Categoryes")
	if err != nil {
		return "", fmt.Errorf("error al ejecutar la consulta: %w", err)
	}
	defer rows.Close()

	// Crear un slice para almacenar las Categoryes
	var Categoryes []Category

	// Iterar sobre las filas y almacenar los valores de cada columna
	for rows.Next() {
		var Category Category
		if err := rows.Scan(&Category.CategoryID, &Category.Tipo); err != nil {
			return "", fmt.Errorf("error al escanear fila: %w", err)
		}
		Categoryes = append(Categoryes, Category)
	}

	// Convertir el slice de Categoryes a formato JSON como string
	CategoryesJSON, err := json.Marshal(Categoryes)
	if err != nil {
		return "", fmt.Errorf("error al convertir a JSON: %w", err)
	}

	// Convertir el JSON byte slice a string
	CategoryesString := string(CategoryesJSON)

	return CategoryesString, nil
}
