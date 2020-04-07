package databases

import (
	"context"
	"fmt"
	"log"
	"time"

	//driver para mongo

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Conectar a la base de datos
var (
	client *mongo.Client
	err    error
)

//Conectar a la base de datos
func Conectar() {

	host := "localhost"
	port := "27017"

	conn := fmt.Sprintf("mongodb://%s:%s", host, port)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))

	if client != nil {
		fmt.Println("MongoDB conectada")
	}
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

/* Insertat un elemento en una coleccion */
func InsertOne(database string, document string, data interface{}) {
	/* Crea una instancia de la coleccion */
	collection := client.Database(database).Collection(document)
	/* Abre un contexto */
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	/* Inserta un dato */
	res, err := collection.InsertOne(ctx, data)
	/*  chequea error*/
	checkError(err)
	id := res.InsertedID

	fmt.Println(id)
}
