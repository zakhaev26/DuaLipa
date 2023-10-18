package controller

import (
	"encoding/json"
	"net/http"

	helper "github.com/zakhaev26/dualipa/helpers"
	model "github.com/zakhaev26/dualipa/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllImages(w http.ResponseWriter, r *http.Request) {
	var allImages []primitive.M = helper.GetAllImages()
	json.NewEncoder(w).Encode(allImages)
}

func PushOneImage(w http.ResponseWriter, r *http.Request) {
	var dua model.DuaLipa
	_ = json.NewDecoder(r.Body).Decode(&dua)
	helper.InsertOneImage(dua)
	json.NewEncoder(w).Encode(dua)
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	num := helper.DeleteAllImage()
	var res = map[string]int{
		"deletedCount": num,
	}
	json.NewEncoder(w).Encode(res)
}
/*
func UpdateOneImage(w http.ResponseWriter, r *http.Request) {

	var dua model.DuaLipa
	_ = json.NewDecoder(r.Body).Decode(&dua)
	// helper.UpdateOneImage((dua.ID),dua.ImgSrc)
}
*/
