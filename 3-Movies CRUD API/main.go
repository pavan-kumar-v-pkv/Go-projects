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

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie // slice to hold movie data

// getMovies handles GET requests to retrieve all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Encode the movies slice to JSON and write it to the response
	json.NewEncoder(w).Encode(movies)
}

// deleteMovie handles DELETE requests to remove a movie by ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// remove the movie from the slice, by appending the slice before and after the index
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// getMovie handles GET requests to retrieve a single movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			// Encode the found movie to JSON and write it to the response
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// createMovie handles POST requests to add a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)  // Decode the request body into the movie variable
	movie.ID = strconv.Itoa(rand.Intn(1000000)) // Generate a random ID for the movie
	movies = append(movies, movie)              // Append the new movie to the movies slice
	json.NewEncoder(w).Encode(movie)            // Encode the new movie to JSON and write it to the response
}

// updateMovie handles PUT requests to update an existing movie by ID
func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set the response header to application/json
	w.Header().Set("Content-Type", "application/json")
	// params
	params := mux.Vars(r)
	// loop over the movies slice
	// find the movie with the matching ID and delete it
	// add a new movie with the same ID with the updated data
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = item.ID // keep the same ID
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie) // Encode the updated movie to JSON and write it to the response
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	// Initialize the movies slice with some sample data
	movies = append(movies, Movie{
		ID:       "1",
		Isbn:     "438734",
		Title:    "Movie One",
		Director: &Director{Firstname: "John", Lastname: "Doe"},
	})
	movies = append(movies, Movie{
		ID:       "2",
		Isbn:     "438735",
		Title:    "Movie Two",
		Director: &Director{Firstname: "Jane", Lastname: "Smith"},
	})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000...\n")
	// Start the server
	// ListenAndServe is a blocking call, so it will not return until the server is stopped
	// Use ListenAndServe to start the server on port 8000
	// If there is an error starting the server, log it and exit
	// Use log.Fatal to log the error and exit the program
	// Use http.ListenAndServe to start the server
	// Use r as the router for the server
	// Use ":8000" as the address to listen on
	log.Fatal(http.ListenAndServe(":8000", r))

}
