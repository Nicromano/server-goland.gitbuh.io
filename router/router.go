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

/* Estructura para respuestas al cliente */
type Response struct {
	Title       string `json:title`
	Description string `json:description`
}

/* Estructura para almacenamiento de datos de un usuario */
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Image    string `json:image`
}

type Link struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	(*w).Header().Set("Content-Type", "application/json")
	/* (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization") */
}

func init() {
	router = mux.NewRouter().StrictSlash(true)
	databases.Conectar()
	router.HandleFunc("/", home).Methods("GET")
	/* 	router.HandleFunc("/signin", singinUser).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodOptions) */
	router.HandleFunc("/signin", singinUser).Methods("POST")
	router.HandleFunc("/signup", signupUser).Methods("POST")
	router.HandleFunc("/forgot", forgotPassword).Methods("POST")
	router.HandleFunc("/home", homeHandler).Methods("GET")
	router.Use(mux.CORSMethodMiddleware(router))

}

func controlError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func forgotPassword(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

}
func signupUser(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

}
func singinUser(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

	 link := Link{Name:'Google', Url:'https://www.google.com'} 
	 
	databases.InsertOne('links', 'links' )

	res.Write([]byte("Hello word"))
}

func home(res http.ResponseWriter, req *http.Request) {

	res.Write([]byte("Holaa"))
}

//GetRouter Funcion para retornar el router
func GetRouter() *mux.Router {
	return router
}
