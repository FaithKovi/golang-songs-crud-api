package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Song struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Artist *Artist `json:"artist"`
}

type Artist struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var songs []Song

func getSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}

func deleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range songs {

		if item.ID == params["id"] {
			songs = append(songs[:index], songs[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(songs)
}

func getSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range songs {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var song Song
	_ = json.NewDecoder(r.Body).Decode(&song)
	song.ID = strconv.Itoa(rand.Intn(100000000))
	songs = append(songs, song)
	json.NewEncoder(w).Encode(song)
}

func updateSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range songs {
		if item.ID == params["id"] {
			songs = append(songs[:index], songs[index+1:]...)
			var song Song
			_ = json.NewDecoder(r.Body).Decode(&song)
			song.ID = params["id"]
			songs = append(songs, song)
			json.NewEncoder(w).Encode(song)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	songs = append(songs, Song{ID: "1", Isbn: "324522", Title: "Yahweh", Artist: &Artist{FirstName: "John", LastName: "Johnny"}})
	songs = append(songs, Song{ID: "2", Isbn: "542981", Title: "Home", Artist: &Artist{FirstName: "Sady", LastName: "Jule"}})
	r.HandleFunc("/songs", getSongs).Methods("GET")
	r.HandleFunc("/songs/{id}", getSong).Methods("GET")
	r.HandleFunc("/songs", createSong).Methods("POST")
	r.HandleFunc("/songs/{id}", updateSong).Methods("PUT")
	r.HandleFunc("/songs/{id}", deleteSong).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
