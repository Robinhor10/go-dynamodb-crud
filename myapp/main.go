package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/gorilla/mux"
)

var db *dynamodb.DynamoDB

type Item struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/items", createItem).Methods("POST")
    router.HandleFunc("/items/{id}", getItem).Methods("GET")

    awsSession, err := session.NewSession(&aws.Config{
        Region:      aws.String("us-east-1"),
        Endpoint:    aws.String(os.Getenv("DYNAMODB_ENDPOINT")),
        Credentials: credentials.NewStaticCredentials("fakeAccessKeyID", "fakeSecretAccessKey", ""),
    })
    if err != nil {
        log.Fatalf("Falha para criar a sessão da aws: %v", err)
    }
    db = dynamodb.New(awsSession)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func createItem(w http.ResponseWriter, r *http.Request) {
    var item Item
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    av, err := dynamodbattribute.MarshalMap(item)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    input := &dynamodb.PutItemInput{
        TableName: aws.String("Items"),
        Item:      av,
    }

    _, err = db.PutItem(input)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(item)
}

func getItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    input := &dynamodb.GetItemInput{
        TableName: aws.String("Items"),
        Key: map[string]*dynamodb.AttributeValue{
            "id": {
                S: aws.String(id),
            },
        },
    }

    result, err := db.GetItem(input)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if result.Item == nil {
        http.Error(w, "Item não encontrado", http.StatusNotFound)
        return
    }

    var item Item
    err = dynamodbattribute.UnmarshalMap(result.Item, &item)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(item)
}
