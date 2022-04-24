package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(Board)
}

const Size = 20

var Board = make([][]string, Size)

func resetBoard() {
	for i := range Board {
		Board[i] = make([]string, Size)
		for j, _ := range Board[i] {
			Board[i][j] = "-"
		}
	}
}

func main() {
	resetBoard()

	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc("/hello", hello)

	http.ListenAndServe(":8080", router)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
