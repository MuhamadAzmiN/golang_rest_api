package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/username/go_rest_api_crud/common/db"
	"github.com/username/go_rest_api_crud/repo"
	"github.com/username/go_rest_api_crud/handlers"
    "github.com/username/go_rest_api_crud/routes"
	"github.com/username/go_rest_api_crud/service"
)

func main() {
    // Database configuration
    dbConfig := &db.Config{
        Host:     "localhost",
        Port:     "3306",
        User:     "root",
        Password: "",
        DBName:   "go_rest",
    }

    // Initialize database connection
    database := db.NewConnection(dbConfig)

    // Initialize repositories
    productRepo := repo.NewProductRepository(database)

    // Initialize services
    productService := services.NewProductService(productRepo)
    // Initialize router
    router := mux.NewRouter()
    
    // Setup routes
    handler := handlers.NewProductHandler(productService)
    handlers.NewProductHandler(productService)
    routes.SetupRoutes(router, handler)

    // Start server
    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}