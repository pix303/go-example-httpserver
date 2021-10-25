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

func addGenreHandler(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	if method != http.MethodPost {
		http.Error(w, "Method not accepted", http.StatusMethodNotAllowed)
		return
	}
	// var item []byte
	// itemSize, err := r.Body.Read(item)
	// defer r.Body.Close()

	var g Genre
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		http.Error(w, "Error on read body request", http.StatusBadRequest)
		return
	}
	log.Println(g)

	w.Write([]byte("Genre added"))
}

func logHeaders(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("---Headers for %s\r", r.URL.Path)
		for k, v := range r.Header {
			log.Println(k, v)
		}
		log.Printf("---Headers-------------end-----\n\n")
		h.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/genre", logHeaders(genreHandler))
	http.HandleFunc("/genre/add", logHeaders(addGenreHandler))
	log.Println("Server running...")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
