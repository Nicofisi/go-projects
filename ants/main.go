package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"math/rand"
	"net/http"
)

const Size = 20
const TargetAntCount = 30

const EmptyCell = "⬜"
const AntCell = "🐜"

var Board = make([][]string, Size)

func generateAnts() {
	antCount := 0
	for antCount < TargetAntCount {
		i := rand.Intn(Size)
		j := rand.Intn(Size)
		if Board[i][j] == EmptyCell {
			Board[i][j] = AntCell
			antCount += 1
		}
	}
}

func resetBoard() {
	for i := range Board {
		Board[i] = make([]string, Size)
		for j, _ := range Board[i] {
			Board[i][j] = EmptyCell
		}
	}
}

func restartSimulation() {
	resetBoard()
	generateAnts()
}

func getBoardHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(Board)
}

func restartHandler(w http.ResponseWriter, req *http.Request) {
	restartSimulation()
	getBoardHandler(w, req)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	restartSimulation()

	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc("/board", getBoardHandler)
	router.HandleFunc("/restart", restartHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	http.ListenAndServe(":8080", handler)
}
