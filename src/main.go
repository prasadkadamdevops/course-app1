package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Course struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
    data := []Course{
        {ID: 1, Name: "DevOps Basics"},
        {ID: 2, Name: "Kubernetes Beginner"},
        {ID: 3, Name: "Docker & CI/CD"},
    }
    json.NewEncoder(w).Encode(data)
}

func main() {
    http.HandleFunc("/courses", coursesHandler)
    log.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}

