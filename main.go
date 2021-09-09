package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Genre struct {
	Name      string            `json:"name"`
	Rate      int               `json:"rate"`
	Subgenres map[string]string `json:"subgenres"`
}

func metalHandler(w http.ResponseWriter, r *http.Request) {
	metal := Genre{
		Name: "metal",
		Rate: 6,
		Subgenres: map[string]string{
			"OO": "doom",
			"BT": "british",
			"TH": "Thrash",
		},
	}

	result, err := json.Marshal(metal)
	if err != nil {
		http.Error(w, "Error on json", 500)
	}
	w.Write([]byte(result))
}

var start time.Time

func logRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start = time.Now()
		log.Printf("Request on process: %s - %s\r", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	}
}

func logResponse(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		log.Printf("Request processed in %d\r", time.Since(start).Milliseconds())
	}
}

func main() {
	http.HandleFunc("/metal", logRequest(logResponse(metalHandler)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
