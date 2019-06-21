package database
import (
	"fmt"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func CreateConnectionHandler()(*mongo.Database, error){
	uri := "mongodb://datastore:27017";
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil { 
		fmt.Println("error: ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil { 
		fmt.Println(err)
	}

	DB := client.Database("secretsquirrelNEST")
	return DB, nil
}

func CreateConnection(){
	db, err := CreateConnectionHandler()
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("after db, err in main")
	DB = db
}