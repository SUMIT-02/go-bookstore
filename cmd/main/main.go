package main

import (
	"log"
	"net/http"

	"github.com/SUMIT-02/go-bookstore/pkg/routes"
	_ "github.com/SUMIT-02/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
