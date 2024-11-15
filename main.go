package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "receipt_processor/handlers"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/receipts/process", handler.ProcessReceipt).Methods("POST")
    r.HandleFunc("/receipts/{id}/points", handler.GetPoints).Methods("GET")

    http.ListenAndServe(":8080", r)
}