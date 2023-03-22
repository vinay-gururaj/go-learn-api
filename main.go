package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Roll struct {
	ID          string `json:"id"`
	ImageNumber string `json:"imageNumber"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

var rolls []Roll

func main() {

	rolls = append(rolls, Roll{ID: "1", ImageNumber: "8", Name: "Spicy Tuna Roll", Ingredients: "Tuna, Chilli sauce, Nori, Rice"},
		Roll{ID: "2", ImageNumber: "9", Name: "Vegetarian Roll", Ingredients: "Avocado, Cucumber, Chilli sauce, Nori, Rice"},
		Roll{ID: "3", ImageNumber: "10", Name: "California Roll", Ingredients: "Crab, Avocado, Cucumber, Nori, Rice"})

	router := mux.NewRouter()

	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))

}

func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rolls)
}

func getRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range rolls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newRoll Roll
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll.ID = strconv.Itoa(len(rolls) + 1)
	rolls = append(rolls, newRoll)
	json.NewEncoder(w).Encode(newRoll)

}

func updateRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			var newRoll Roll
			json.NewDecoder(r.Body).Decode(&newRoll)
			newRoll.ID = params["id"]
			rolls = append(rolls, newRoll)
			json.NewEncoder(w).Encode(newRoll)
			return
		}
	}

}

func deleteRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			break

		}
	}
	json.NewEncoder(w).Encode(rolls)
}
