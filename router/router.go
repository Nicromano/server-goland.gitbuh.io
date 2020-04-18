package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../crypto"
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
	Username  string        `json:"username"    bson:"username" `
	Password  string        `json:"password"    bson:"password"`
	Email     string        `json:"email"    	bson:"email"`
	Image     string        `json:"image"   	bson:"image"`
	Links     []Link        `json:"links"    	bson:"links"`
	Follow    []interface{} `json:"follow"   	bson:"follow"`
	Followers []interface{} `json:"followers"   bson:"followers"`
}

/* Estructura de link */
type Link struct {
	Name        string    `json:"name"  		bson:"name" 	`
	Url         string    `json:"url"  			bson:"url" 		`
	Description string    `json:"description"  	bson:"description" `
	Comments    []Comment `json:"comments"  	bson:"comments"`
	Like        uint32    `json:"like"  		bson:"like"`
	Dislike     uint32    `json:"dislike"  		bson:"dislike"`
}
type Comment struct {
	IdUser  string `json:"iduser"  		bson:"iduser"`
	Content string `json:"content"  	bson:"content"`
	Like    uint32 `json:"like"  		bson:"like"`
	Dislike uint32 `json:"dislike"  	bson:"dislike"`
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
	_ = json.NewDecoder(req.Body).Decode(&user)

	fmt.Printf("%s\n", user)

	findUser := databases.FindOne("links", "User", "username", user.Username)
	if findUser != nil {
		//Encontró
		fmt.Println(findUser)
		json.NewEncoder(w).Encode(Response{Title: "USER FOUND"})
	} else {
		//No encontró
		err := databases.InsertOne("links", "User", user)
		if err != nil {
			//No ingreso datos
			json.NewEncoder(w).Encode(Response{Title: "UNREGISTERED USER"})
		} else {
			//ingresó datos
			json.NewEncoder(w).Encode(Response{Title: "REGISTERED USER"})
		}
	}
}

func singinUser(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)

	findUser := databases.FindOne("links", "User", "username", user.Username)
	fmt.Printf("%v", findUser)
	if findUser != nil {
		//Se encontró el usuario
		//Desencriptar contraseña

		json.NewEncoder(w).Encode(Response{Title: "USER FOUND"})

	} else {
		json.NewEncoder(w).Encode(Response{Title: "USER NOT FOUND"})

	}

}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	setupResponse(&res, req)

	links := Link{Name: "Google", Url: "https://www.google.com"}

	err := databases.InsertOne("links", "links", links)
	res.Write([]byte("id ingresado" + err.Error()))
}

func home(res http.ResponseWriter, req *http.Request) {

	text := "Hola como esta"

	hash := crypto.Encrypt(text)

	fmt.Printf("%v\n", hash)

	textPlain := crypto.Decrypt(hash)

	fmt.Printf("%v", textPlain)
	res.Write([]byte("Hola"))

}

//GetRouter Funcion para retornar el router
func GetRouter() *mux.Router {
	return router
}
