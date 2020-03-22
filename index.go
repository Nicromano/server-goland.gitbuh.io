package main

import (
	"fmt"
	"net/http"

	"./router"
)

func main() {
	r := router.GetRouter()

	fmt.Println("Servidor escuchando en el puerto 3000")
	
	http.ListenAndServe(":3000", r)
}
