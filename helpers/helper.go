package helper

import (
	"context"
	"fmt"

	dbconn "github.com/zakhaev26/dualipa/dbconn"
	hanErr "github.com/zakhaev26/dualipa/errors"
	model "github.com/zakhaev26/dualipa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// insert one Image
func InsertOneImage(dua model.DuaLipa) {

	inserted, err := dbconn.Collection.InsertOne(context.TODO(), dua)
	hanErr.HandleErr(err)

	fmt.Println("Inserted a Dua Lipa Pic with id ::", inserted.InsertedID)
}

// updateoneimage
func UpdateOneImage(duaId string, newSrc string) {
	objectID, err := primitive.ObjectIDFromHex(duaId)
	hanErr.HandleErr(err)
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"imgsrc": newSrc}}

	res, err := dbconn.Collection.UpdateOne(context.Background(), filter, update)
	hanErr.HandleErr(err)
	fmt.Println("Modified Count ::", res.ModifiedCount)
}

// deleteONEimage
func DeleteOneImage(duaId string) {
	objectId, err := primitive.ObjectIDFromHex(duaId)
	hanErr.HandleErr(err)
	filter := bson.M{"_id": objectId}
	deleteCount, err := dbconn.Collection.DeleteOne(context.Background(), filter)
	hanErr.HandleErr(err)
	fmt.Println("Delete Count ::", deleteCount)
}

//deleteAllRecords [DESTRUCTIVE]

func DeleteAllImage() int {
	deleteRes, err := dbconn.Collection.DeleteMany(context.Background(), bson.M{}, nil)
	hanErr.HandleErr(err)
	fmt.Println("Number of sources deleted : ", deleteRes.DeletedCount)
	return int(deleteRes.DeletedCount)
}

/*
The primitive.M type is often used to represent MongoDB documents retrieved from the database. It allows you to work with the data in a convenient map-like structure, where the keys are the field names in the document and the values are the corresponding field values.
*/

func GetAllImages() []primitive.M {
	cur, err := dbconn.Collection.Find(context.Background(), bson.D{})
	hanErr.HandleErr(err)

	var ImgSrcs []primitive.M

	for cur.Next(context.Background()) {
		var dua bson.M
		err := cur.Decode(&dua)
		hanErr.HandleErr(err)
		ImgSrcs = append(ImgSrcs, dua)
	}

	defer cur.Close(context.Background())
	return ImgSrcs
}
