package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {

	users := []User{
		{
			ID:    1,
			Name:  "Anil",
			Email: "anil@example.com",
		},
		{
			ID:    2,
			Name:  "Rahul",
			Email: "rahul@example.com",
		},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("user-service healthy"))
}

func main() {

	http.HandleFunc("/users", usersHandler)

	http.HandleFunc("/health", healthHandler)

	log.Println("User service running on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
