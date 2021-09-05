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

func main() {

	mux := http.NewServeMux()
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
