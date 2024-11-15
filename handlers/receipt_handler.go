package handler

import (
	"encoding/json"
	"net/http"
	"receipt_processor/models"
	"receipt_processor/utils"
	"sync"

	"github.com/gorilla/mux"
)

var (
	receipts = make(map[string]models.Receipt)
	mu       sync.Mutex
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid receipt", http.StatusBadRequest)
		return
	}

	receipt.ID = utils.GenerateID()
	mu.Lock()
	receipt.Points = utils.CalculatePoints(receipt)
	receipts[receipt.ID] = receipt
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": receipt.ID})
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	mu.Lock()
	receipt, exists := receipts[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "No receipt found for that id", http.StatusNotFound)
		return
	}

	points := receipt.Points
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.PointsResponse{Points: points})
}
