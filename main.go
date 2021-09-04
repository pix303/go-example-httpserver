package main

import (
	"net/http"
)

type srv string

func (s srv) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/metal":
		w.Write([]byte("sei metal"))
	case "/dark":
		w.Write([]byte("sei dark"))
	case "/pop":
		w.Write([]byte("sei pop"))
	}
}

func main() {
	var s srv

	http.ListenAndServe(":8080", s)

}
