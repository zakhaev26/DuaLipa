package helper

import (
	"context"

	dbconn "github.com/zakhaev26/dualipa/dbconn"
	model "github.com/zakhaev26/dualipa/models"
)

func insertOneImage(dua model.DuaLipa) {
	
	inserted, err := dbconn.Collection.InsertOne(context.TODO(), dua)
	dbconn.HandleErr(err)
}
