package main

import (
	"net/http"
	"tentangKode/controllers"
	"tentangKode/middleware"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")

	router.HandleFunc("/product", middleware.MiddlewareJWTAuthorization(controllers.GetAllProducts)).Methods("GET")
	router.HandleFunc("/product/{id}", middleware.MiddlewareJWTAuthorization(controllers.GetProductById)).Methods("GET")
	router.HandleFunc("/product", middleware.MiddlewareJWTAuthorization(controllers.AddProduct)).Methods("POST")
	router.HandleFunc("/product/{id}", middleware.MiddlewareJWTAuthorization(controllers.EditProduct)).Methods("PUT")
	router.HandleFunc("/product/{id}", middleware.MiddlewareJWTAuthorization(controllers.DeleteProduct)).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
