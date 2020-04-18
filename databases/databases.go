package databases

import (
	"context"
	"fmt"
	"log"
	"time"

	//driver para mongo

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Conectar a la base de datos
var (
	client *mongo.Client
	err    error
)

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

//Conectar a la base de datos
func Conectar() {

	host := "localhost"
	port := "27017"

	conn := fmt.Sprintf("mongodb://%s:%s", host, port)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(conn))
	controlError(err)
	if client != nil {
		fmt.Println("MongoDB conectada")
	}

}

// Insertat un elemento en una coleccion
func InsertOne(database string, document string, data interface{}) error {
	/* Crea una instancia de la coleccion */
	collection := client.Database(database).Collection(document)
	/* Abre un contexto */
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	/* Inserta un dato */
	res, err := collection.InsertOne(ctx, convertBSON(data))
	/*  chequea error*/
	controlError(err)
	fmt.Printf("%v", res.InsertedID)
	return err
}

func controlError(err error) {
	if err != nil {
		log.Println(err)
	}
	log.Println("----END OF ERROR----")
}

//Buscar todos los elementos
func FindAll(database string, document string) []interface{} {

	var results []interface{}
	collection := client.Database(database).Collection(document)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cursor, err := collection.Find(ctx, bson.M{})
	controlError(err)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result interface{}
		cursor.Decode(&result)
		results = append(results, result)
	}

	return results
}

//Encuentra uno
func FindOne(database, document, clave, valor string) interface{} {
	var result User
	collection := client.Database(database).Collection(document)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	filter := bson.D{{clave, valor}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		fmt.Printf("Dato %v no encontrado en la coleccion %s\n", filter, document)
		return nil
	}

	return result
}

//Actualiza un documento
func UpdateOne(database string, document string, find string, replace interface{}) {
	collection := client.Database(database).Collection(document)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	id, _ := primitive.ObjectIDFromHex(find)

	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", replace}})

	controlError(err)

	fmt.Printf("Actualizado %v", result.ModifiedCount)

}

//Actualiza algunos elementos
func UpdateMany(database string, document string, filter interface{}, replace interface{}) {
	collection := client.Database(database).Collection(document)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	result, err := collection.UpdateMany(ctx, filter, bson.D{{"$set", replace}})

	controlError(err)

	fmt.Printf("Actualizados %v documentos", result.ModifiedCount)
}

//Elimina todos los elementos de una coleccion
func DeleteAll(database string, document string) {
	collection := client.Database(database).Collection(document)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = collection.Drop(ctx)
	controlError(err)

}

//Elimina algunos elementos de una coleccion
func DeleteMany(database string, document string, busqueda interface{}) {
	collection := client.Database(database).Collection(document)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	deleteResult, err := collection.DeleteMany(ctx, busqueda)
	controlError(err)
	fmt.Printf("Fueron eliminados %v elementos de la coleccion %s", deleteResult.DeletedCount, document)

}

//Elimina un elemento de una coleccion
func DeleteOne(database string, document string, busqueda interface{}) {

	collection := client.Database(database).Collection(document)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	deleteResult, err := collection.DeleteOne(ctx, convertBSON(busqueda))
	controlError(err)
	fmt.Printf("Fue eliminado %v de la coleccion %s", deleteResult, document)

}

func convertBSON(data interface{}) []byte {
	b, err := bson.Marshal(data)
	controlError(err)
	return b
}
