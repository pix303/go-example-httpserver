package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GenreResponse struct {
	Genre string `json:"genre"`
	Rate  int    `json:"rate"`
}

type BluesHandler struct{}

func (bh BluesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mockResult := GenreResponse{"Blues", 8}
	result, err := json.Marshal(mockResult)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error on parsing data: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

func main() {

	mux := http.NewServeMux()
	bh := BluesHandler{}
	mux.Handle("/blues", bh)
	mux.HandleFunc("/metal", func(w http.ResponseWriter, r *http.Request) {
		mockResult := GenreResponse{
			"Britsh Metal", 6,
		}
		result, err := json.Marshal(mockResult)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error on data: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		w.Write(result)
	})
	mux.HandleFunc("/pop", func(w http.ResponseWriter, r *http.Request) {
		mockResult := GenreResponse{
			"K Pop", 3,
		}
		result, err := json.Marshal(mockResult)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error on data: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		w.Write(result)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))

}
