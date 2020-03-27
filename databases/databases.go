package databases

import (
	"database/sql"
	"fmt"

	//driver para mysql
	_ "github.com/go-sql-driver/mysql"
)

//Conectar a la base de datos
var (
	db  *sql.DB
	err error
)

//Conectar a la base de datos
func Conectar() {

	user := "root"
	password := "pass123"
	/* host := "tcp(172.0.0.1:3306)" */
	database := "link"
	conn := fmt.Sprintf("%s:%s@/%s", user, password, database)
	/* db, err = sql.Open("mysql", user+":"+password+"@/"+database) */
	db, err = sql.Open("mysql", conn)

	if db != nil {
		fmt.Println("Base de datos conectada")
	}
	if err != nil {
		fmt.Println(err)
	}

}

/* ConsultaSQL exportar funcion para   */
func ConsultaSQL(query string, datos ...interface{}) (*sql.Rows, error) {

	row, er := db.Query(query, datos...)
	return row, er
}
