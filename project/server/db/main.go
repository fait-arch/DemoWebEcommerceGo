package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Configura los detalles de conexión
	dbUser := ""						//tu_usuario
	dbPass := ""					//tu_contraseña
	dbHost := ""	//servidor_de_base_de_datos:3306
	dbName := ""		//nombre_de_la_base_de_datos

	// Crea la cadena de conexión
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	fmt.Println(connStr)



	// Conéctate a la base de datos
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	// Consulta en la base de datos
	rows, err := db.Query("SELECT tipo FROM Propiedades")
	if err != nil {
		log.Fatal("Error al ejecutar la consulta:", err)
	}
	defer rows.Close()

	// Si llegamos aquí, la conexión fue exitosa
	fmt.Println("Conexión exitosa a la base de datos")

	// Iterar sobre las filas y mostrar los valores de la columna "tipo"
	for rows.Next() {
		var tipo string
		if err := rows.Scan(&tipo); err != nil {
			log.Fatal("Error al escanear fila:", err)
		}
		fmt.Printf("Tipo: %s\n", tipo)
	}
	if err := rows.Err(); err != nil {
		log.Fatal("Error al iterar sobre las filas:", err)
	}
}



