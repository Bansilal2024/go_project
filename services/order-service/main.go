package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Order struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Item   string `json:"item"`
	Status string `json:"status"`
}

func ordersHandler(w http.ResponseWriter, r *http.Request) {
	orders := []Order{
		{
			ID:     101,
			UserID: 1,
			Item:   "Laptop",
			Status: "COMPLETED",
		},
		{
			ID:     102,
			UserID: 2,
			Item:   "Keyboard",
			Status: "PENDING",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("order-service healthy"))
}

func main() {
	http.HandleFunc("/orders", ordersHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("Order service running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}