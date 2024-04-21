package RoadPropiedades

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Propiedad representa la estructura de un resultado de consulta
type Propiedad struct {
	PropiedadID int    `json:"propiedad_id"`
	Tipo        string `json:"tipo"`
}

// ObtenerPropiedadesString realiza la consulta a la base de datos y devuelve los resultados en formato JSON como un string
func ObtenerPropiedades() (string, error) {
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
	rows, err := db.Query("SELECT PropiedadID, Tipo FROM Propiedades")
	if err != nil {
		return "", fmt.Errorf("error al ejecutar la consulta: %w", err)
	}
	defer rows.Close()

	// Crear un slice para almacenar las propiedades
	var propiedades []Propiedad

	// Iterar sobre las filas y almacenar los valores de cada columna
	for rows.Next() {
		var propiedad Propiedad
		if err := rows.Scan(&propiedad.PropiedadID, &propiedad.Tipo); err != nil {
			return "", fmt.Errorf("error al escanear fila: %w", err)
		}
		propiedades = append(propiedades, propiedad)
	}

	// Convertir el slice de propiedades a formato JSON como string
	propiedadesJSON, err := json.Marshal(propiedades)
	if err != nil {
		return "", fmt.Errorf("error al convertir a JSON: %w", err)
	}

	// Convertir el JSON byte slice a string
	propiedadesString := string(propiedadesJSON)

	return propiedadesString, nil
}
