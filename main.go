package rpglms_go

import (
	"log"
	"net/http"

	"github.com/studioaugustus/rpglms_go/internal/app"
)

func main() {
    handler := app.NewHandler()
    http.HandleFunc("/", handler.ServeHTTP)

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}