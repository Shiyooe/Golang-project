
package main

import (
    "log"
    "net/http"
    "simple-api/config"
    "simple-api/internal/handler"
    "simple-api/internal/service"
)

func main() {
    // Load configuration
    cfg := config.Load()

    // Initialize service
    userService := service.NewUserService()

    // Initialize handler
    userHandler := handler.NewUserHandler(userService)

    // Setup routes
    http.HandleFunc("/users", userHandler.GetUsers)
    http.HandleFunc("/user", userHandler.CreateUser)

    // Start server
    log.Printf("Server starting on port %s", cfg.Port)
    if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
        log.Fatal(err)
    }
}