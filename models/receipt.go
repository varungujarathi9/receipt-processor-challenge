package models
type Receipt struct {
    ID           string    `json:"id"`
    Retailer     string    `json:"retailer"`
    PurchaseDate string    `json:"purchaseDate"`
    PurchaseTime string    `json:"purchaseTime"`
    Items        []Item    `json:"items"`
    Total        string    `json:"total"`
}