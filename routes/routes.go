package routes

import (
    "github.com/gorilla/mux"
    "github.com/username/go_rest_api_crud/handlers"
)

func SetupRoutes(r *mux.Router, productHandler *handlers.ProductHandler) {
    r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
    r.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
    r.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
    r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
    r.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")
}