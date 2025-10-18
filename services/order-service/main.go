package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Order struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Status    string  `json:"status"`
}

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/health", healthCheck).Methods("GET")
	r.HandleFunc("/api/orders", getOrders).Methods("GET")
	r.HandleFunc("/api/orders", createOrder).Methods("POST")

	port := getEnv("PORT", "3003")
	log.Printf("Order service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "healthy",
		"service": "order-service",
		"version": "1.0.0",
	}
	json.NewEncoder(w).Encode(response)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	orders := []Order{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	order.Status = "pending"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
