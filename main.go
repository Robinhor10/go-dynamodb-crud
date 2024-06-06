package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "go-dynamodb-crud/handler"
    "go-dynamodb-crud/dynamo"
    "log"
)

func main() {
    dynamo.InitDynamo()
    r := mux.NewRouter()
    r.HandleFunc("/items", handler.CreateItem).Methods("POST")
    r.HandleFunc("/items/{id}", handler.GetItem).Methods("GET")
    r.HandleFunc("/items/{id}", handler.DeleteItem).Methods("DELETE")
    r.HandleFunc("/items", handler.ListItems).Methods("GET")

    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}
