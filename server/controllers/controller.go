package controller

import (
	"encoding/json"
	"net/http"

	hanErr "github.com/zakhaev26/dualipa/errors"
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
	err := json.NewDecoder(r).Decode(&dua)
	hanErr.

}
