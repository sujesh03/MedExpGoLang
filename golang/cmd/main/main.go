package main

import (
	"log"
	"net/http"

	"github.com/sujesh03/ExpenseTracker/pkg/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterETStoreRoutes(r)
	http.Handle("/", r)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Autherization", "access-token"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe("localhost:8010", handlers.CORS(headers, methods, origins)(r)))
}
