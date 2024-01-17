package app

import "net/http"

type Handler struct {
    // dependencies like services
}

func NewHandler() *Handler {
    return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // handle requests
    w.Write([]byte("Hello RPG LMS"))
}