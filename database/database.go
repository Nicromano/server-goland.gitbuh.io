package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Conectar a la base de datos
func Conectar() {

	user := "root"
	password := "pass123"
	database := "deposito"

	db, err := sql.Open("mysql", user+":"+password+"@/"+database)

	if db != nil {
		fmt.Println("Base de datos conectada")
	}
	if err != nil {
		fmt.Println(err)
	}

}
