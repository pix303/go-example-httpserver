package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Genre struct {
	Name      string            `json:"name"`
	Rate      int               `json:"rate"`
	Subgenres map[string]string `json:"subgenres"`
}

var genres = map[string]Genre{
	"metal": {
		Name: "metal",
		Rate: 6,
		Subgenres: map[string]string{
			"OO": "doom",
			"BT": "british",
			"TH": "Thrash",
		},
	},
	"pop": {
		Name: "pop",
		Rate: 6,
		Subgenres: map[string]string{
			"KP": "kpop",
			"MI": "melodica italiana",
			"60": "sixties",
		},
	},
}

func genreHandler(w http.ResponseWriter, r *http.Request) {

	value := r.URL.Query().Get("name")
	var result []byte
	var err error

	if value == "metal" || value == "pop" {
		result, err = json.Marshal(genres[value])
	} else {
		http.Error(w, "Unknown genre", http.StatusBadRequest)
	}

	if err != nil {
		http.Error(w, "Error on json", 500)
	}
	w.Write([]byte(result))
}

func main() {
	http.HandleFunc("/genre", genreHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
