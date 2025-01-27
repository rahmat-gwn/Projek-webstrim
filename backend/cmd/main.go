package main

import (
    "log"
    "net/http"

    "backend/config"
    "backend/middleware"
    "backend/routes"
)

func main() {
    config.ConnectDatabase()

    router := routes.SetupRoutes()

    // Tambahkan CORS middleware
    handler := middleware.CORSMiddleware(router)

    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}
