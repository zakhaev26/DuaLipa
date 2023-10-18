package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	route "github.com/zakhaev26/dualipa/routes"
)

func main() {
	fmt.Println("DuAPI is Starting...")
	router := route.Router()
	// Handle CORS
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":6969", handlers.CORS(headers, methods, origins)(router)))
	fmt.Println("API is Live @ http://localhost:9030")
}
