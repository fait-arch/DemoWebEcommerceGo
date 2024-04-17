## Comandos de inicio
Ejecuta estos comandos en la terminal en orden 

| COMANDO            | EXPLICACON                                                                                | RUTA DE EJECUCION |
| ------------------ | ----------------------------------------------------------------------------------------- | ----------------- |
| `npm install`      | Instalar dependecias de nodejs                                                            |                   |
| `npm run build`    | Construir la aplicación de JavaScript para producción que leerá el server en go.          |                   |
| `go run server.go` | Correr servidor "server.go" local en el puerto :3000                                      |                   |
| `air`              | Correr servidor "server.go" y se actualiza en tiempo real los cambios en el puerto :3000. |                   |
| `npm run dev`      | Correr servidor local en el puerto :5137 con el motor de ejecución de nodejs              |                   |

### Tecnologias
- Vite
- React
- nodejs
- Go
- MySQL
- FireBase

### Librería de apoyo
- [Tailwind]()
- [Air](https://github.com/cosmtrek/air)
- 

### Diagrama de Pantallas 
Se hizo bocetos iniciales de como se vería gráficamente la interface grafica con los componentes de react ["archivo"🔗](img/ScreenDiagrams.excalidraw.png) 

### Diagramas ER
Documentación de uso de [" MERMAID 🔗"](https://mermaid.js.org/intro/)
Para ver de forma grafica la distribución de las bases de datos el diagrama esta en el  ["archivo"🔗](img/ER.png) 
#### CodigoMermaidER
	erDiagram
	PRODUCTS_DATA ||--o{ CATEGORY : Belongs_to
	PRODUCTS_DATA {
		id_Products INT pk
		name_Products VARCHA
		image_Products
		imageAltProducts_Products
		price_Products FLOAT 
		description_Products TEXT 
		category_id INT fk 
	    }
	CATEGORY {
		id_Category INT pk
		name_Category VARCHAR
	    }
	USERS  {
		id_User  INT pk
		username_User  VARCHA
		email_User  VARCHA
		password_User  VARCHA
		created_at_User  VARCHA
	    }
#### CodigoSQLER
	CREATE TABLE CATEGORY (
	    id_Category INT PRIMARY KEY,
	    name_Category VARCHAR(255)
	);
	
	CREATE TABLE PRODUCTS_DATA (
	    id_Products INT PRIMARY KEY,
	    name_Products VARCHAR(255),
	    image_Products VARCHAR(255),
	    imageAltProducts_Products VARCHAR(255),
	    price_Products FLOAT,
	    description_Products TEXT,
	    category_id INT,
	    FOREIGN KEY (category_id) REFERENCES CATEGORY(id_Category)
	);
	
	CREATE TABLE USERS (
		id_User INT PRIMARY KEY AUTO_INCREMENT,
		username_User VARCHAR(255),
		email_User VARCHAR(255),
		password_User VARCHAR(255),
		created_a_Usert TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);