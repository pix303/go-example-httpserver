package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	fmt.Println(metal)
	result, err := json.Marshal(metal)
	if err != nil {
		http.Error(w, "Error on json", 500)
	}
	w.Write([]byte(result))
}

func main() {
	http.HandleFunc("/metal", metalHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
