package route

import (
	"github.com/gorilla/mux"
	controller "github.com/zakhaev26/dualipa/controllers"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home).Methods("GET")
	r.HandleFunc("/all-image", controller.GetAllImages).Methods("GET")
	r.HandleFunc("/push-one", controller.PushOneImage).Methods("POST")
	r.HandleFunc("/delete", controller.DeleteAll).Methods("GET")
	return r
}
