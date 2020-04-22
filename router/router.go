package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"../bcrypt"
	"../databases"
	"../model"
	"github.com/gorilla/mux"
)

var (
	router *mux.Router
)

// Estructura de respuesta */
type Response struct {
	Title       string `json:"title"`
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
	router.HandleFunc("/signin", singinUser).Methods("POST")
	router.HandleFunc("/signup", signupUser).Methods("POST")
	router.HandleFunc("/forgot", forgotPassword).Methods("POST")
	router.HandleFunc("/addLink/{user}", addLink).Methods("POST")
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

	var user model.User
	_ = json.NewDecoder(req.Body).Decode(&user)
	hash, err := bcrypt.GeneratePassword(user.Password)
	if err != nil {
		fmt.Println(err)
	}

	user.Password = string(hash)
	_, err = databases.FindOne("links", "User", "username", user.Username)

	if err == nil {

		//Encontró
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
	var user model.User
	_ = json.NewDecoder(req.Body).Decode(&user)

	findUser, err := databases.FindOne("links", "User", "username", user.Username)
	if err == nil {
		//Se encontró el usuario
		//Desencriptar contraseña
		compare := bcrypt.ComparatePassword(findUser.Password, user.Password)
		if compare {
			json.NewEncoder(w).Encode(Response{Title: "USER FOUND"})

		} else {
			json.NewEncoder(w).Encode(Response{Title: "INCORRECT PASSWORD"})

		}

	} else {
		json.NewEncoder(w).Encode(Response{Title: "USER NOT FOUND"})

	}

}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	setupResponse(&res, req)

}

func addLink(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

	params := mux.Vars(req)
	user := params["user"]
	var link model.Link
	_ = json.NewDecoder(req.Body).Decode(&link)
	link.Timestamp = time.Now()
	err := databases.FindOneAndUpdate("links", "User", "username", user, link)
	if err != nil {
		//No se encontro el
			json.NewEncoder(w).Encode(Response{Title: "USER NOT UPDATED"})

	} else{
			// se encontró
			json.NewEncoder(w).Encode(Response{Title: "UPDATED USER"})

	}
/*	fmt.Println(userFound)
	var links []model.Link
	links = userFound.Links
	fmt.Println(links)
	links = append(links, link)
	fmt.Println(links)
	userFound.Links = links
	fmt.Println(userFound)*/
}

//GetRouter Funcion para retornar el router
func GetRouter() *mux.Router {
	return router
}
