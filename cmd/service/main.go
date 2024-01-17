package main

import (
	"log"
	"net/http"

	"github.com/studioaugustus/rpglms_go/internal/app"
)

func main() {
    // Create a new instance of the application handler
    handler := app.NewHandler()

    // Register the handler for the root path
    http.Handle("/", handler)

    // Start the HTTP server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}