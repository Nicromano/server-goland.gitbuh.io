package router

import (
	"fmt"
	"log"
	"net/http"

	"../databases"
	"github.com/gorilla/mux"
)

var (
	router *mux.Router
)

/* Estructura de respuesta */
type Response struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

/* Estructura de usuario */
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	Link     Link   `json:"link"`
}

/* Estructura de link */
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
	router.HandleFunc("/signin", singinUser).Methods("GET")
	router.HandleFunc("/signup", signupUser).Methods("POST")
	router.HandleFunc("/forgot", forgotPassword).Methods("POST")
	router.HandleFunc("/home", home).Methods("GET")
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(loggingMiddleware)

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
func controlError(err error) {
	if err != nil {
		log.Panic(err)
	}

}
func forgotPassword(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

}
func signupUser(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

	var user User

}
func singinUser(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

	links := databases.FindAll("links", "links")

	w.Write([]byte(fmt.Sprintf("%s", links)))

}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	setupResponse(&res, req)

	links := Link{Name: "Google", Url: "https://www.google.com"}

	id := databases.InsertOne("links", "links", links)
	res.Write([]byte("id ingresado" + id))
}

func home(res http.ResponseWriter, req *http.Request) {

	databases.DeleteAll("links", "links")

	res.Write([]byte("borrada"))
}

//GetRouter Funcion para retornar el router
func GetRouter() *mux.Router {
	return router
}
