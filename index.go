package main

import (
	"fmt"
	"net/http"
	"database"
)

func main() {

	//routes
	http.HandleFunc("/", homeHandler)

	//Escuchando en el puerto 3000
	fmt.Println("Servidor escuchando en el puerto 3000")
	database.Conectar()
	http.ListenAndServe(":3000", nil)
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello word"))
}
