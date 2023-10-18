package dbconn

import (
	"context"
	"fmt"

	hanErr "github.com/zakhaev26/dualipa/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://b422056:<password>@testcluster.t47d55j.mongodb.net/?retryWrites=true&w=majority"
const dbName = "DuaLipaImages"
const colName = "Picsource"

var Collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	hanErr.HandleErr(err)
	fmt.Println("Connection to MongoDB successful")

	Collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection Instance is ready!!")
}
