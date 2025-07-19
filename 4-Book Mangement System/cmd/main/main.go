package main

import (
	"book-management-system/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// initialize the router
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	// start the server on port 9010
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
