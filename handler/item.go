package handler

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
    "go-dynamodb-crud/dynamo"
    "go-dynamodb-crud/model"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
    var item model.Item
    _ = json.NewDecoder(r.Body).Decode(&item)
    err := dynamo.PutItem(item)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(item)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    item, err := dynamo.GetItem(params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if item == nil {
        http.NotFound(w, r)
        return
    }
    json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    err := dynamo.DeleteItem(params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func ListItems(w http.ResponseWriter, r *http.Request) {
    items, err := dynamo.ListItems()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(items)
}
