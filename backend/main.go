package main

import (
    "encoding/json"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{
        "message": "StartTech API Running",
    })
}

func health(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{
        "status": "healthy",
    })
}

func main() {
    http.HandleFunc("/", home)
    http.HandleFunc("/health", health)

    http.ListenAndServe(":8080", nil)
}