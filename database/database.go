package database

import (
	"database/sql"
	"fmt"
)

/*InitDB Permite hacer la conexion a la BD oracle*/
func InitDB() *sql.DB {
	connectionString := "hector:gWz6%677@tcp(mysql-26699-0.cloudclusters.net:26699)/learning"
	databaseConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Conexion invalida a la BD")
		panic(err.Error()) // Error Handling = manejo de errores
	}
	return databaseConnection
}
