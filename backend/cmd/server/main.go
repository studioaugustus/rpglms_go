package main

import (
	"log"
	"net/http"
)

func main() {
    handler := rpglms.NewHandler()
    http.HandleFunc("/", handler.ServeHTTP)

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}