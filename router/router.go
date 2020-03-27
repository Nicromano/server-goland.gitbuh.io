package router

import (
	"encoding/json"
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
	router.HandleFunc("/home", homeHandler).Methods("GET")
	router.Use(mux.CORSMethodMiddleware(router))

}

func controlError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func signupUser(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)

	log.Println(user)

	_, err := databases.ConsultaSQL("INSERT INTO USUARIO(USERNAME, PASSWORD, EMAIL) VALUES(?, AES_ENCRYPT(?, 'pato'), ?)", user.Username, user.Password, user.Email)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Title: "NO", Description: "Usuario no registrado"})
		log.Fatal(err)
	} else {
		json.NewEncoder(w).Encode(Response{Title: "YES", Description: "Usuario registrado"})

	}

}
func singinUser(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)

	log.Println(user)

	row, err := databases.ConsultaSQL("SELECT * FROM USUARIO WHERE USERNAME = ? AND AES_DECRYPT(PASSWORD, 'pato') = ?", user.Username, user.Password)
	controlError(err)
	if row.Next() {
		json.NewEncoder(w).Encode(Response{Title: "OK", Description: "Usuario encontrado"})
	} else {

		json.NewEncoder(w).Encode(Response{Title: "ERROR", Description: "Usuario no encontrado"})
	}

}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	cliente := "Cliente"
	row, err := databases.ConsultaSQL("select * from persona where rol = ?", cliente)
	if err != nil {
		log.Println(err)
		return
	}

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

func home(res http.ResponseWriter, req *http.Request) {

	res.Write([]byte("Holaa"))
}

//Estructura para datos de persona
type Persona struct {
	cedula, rol, nombre, telefono, direccion interface{}
}

//GetRouter Funcion para retornar el router
func GetRouter() *mux.Router {
	return router
}
