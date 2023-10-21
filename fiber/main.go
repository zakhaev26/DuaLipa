package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to SSE Server!</h1>"))
}

func hello2(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"]);
	w.Write([]byte("ok"))
}

/*



*/

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello).Methods("GET")
	r.HandleFunc("/post", hello2).Methods("POST")
	http.ListenAndServe(":8080", r)
}
