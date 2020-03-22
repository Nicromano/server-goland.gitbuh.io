package router

import (
	"log"
	"net/http"

	"../databases"
	"github.com/gorilla/mux"
)

var (
	router *mux.Router
)

type User struct {
	username string `json: "username, omitempty"`
	password string `json: "password, omitempty"`
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func init() {
	router = mux.NewRouter()
	databases.Conectar()
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/signin", singinUser).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodOptions)

	router.HandleFunc("/home", homeHandler).Methods("GET")

	router.Use(mux.CORSMethodMiddleware(router))

}
func singinUser(w http.ResponseWriter, req *http.Request) {

	setupResponse(&w, req)
	log.Println("holaa")

}

func home(res http.ResponseWriter, req *http.Request) {

	res.Write([]byte("Holaa"))
}

//Estructura para datos de persona
type Persona struct {
	cedula, rol, nombre, telefono, direccion interface{}
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	cliente := "Cliente"
	row, err := databases.ConsultaSQL("select * from persona where rol = ?", cliente)
	if err != nil {
		log.Println(err)
		return
	}
	/* personas := []Persona{} */

	var p Persona
	var contador int8
	for row.Next() {

		err := row.Scan(&p.cedula, &p.rol, &p.nombre, &p.telefono, &p.direccion)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s, %s, %s, %s, %s", p.cedula, p.rol, p.nombre, p.telefono, p.direccion)
		contador++
	}
	log.Println(contador)
	defer row.Close()
	res.Write([]byte("Hello word"))
}

//GetRouter Funcion para retornar el router
func GetRouter() *mux.Router {
	return router
}
