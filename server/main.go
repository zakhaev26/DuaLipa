package main

import (
	"fmt"
	"log"
	"net/http"

	route "github.com/zakhaev26/dualipa/routes"
)

func main() {
	fmt.Println("DuAPI is Starting...")
	router := route.Router()
	log.Fatal(http.ListenAndServe(":9030", router))
	fmt.Println("API is Live @ http://localhost:9030")
}
