package routes

import (
	"github.com/SUMIT-02/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreatBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.GetBookId).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")

}
