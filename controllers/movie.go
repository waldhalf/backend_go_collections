package controllers

import (
	"auth/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Title       string `json:"title"`
	Overview    string `json:"overview"`
	ReleaseDate string `json:"release_date"`
	PosterPath  string `json:"poster_path"`
	ID          int    `json:"id"`
}

// CreateMovie : Création d'un film
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	if (r).Method == "OPTIONS" {
		return
	}

	movie := &models.Movie{}
	json.NewDecoder(r.Body).Decode(movie)
	db.NewRecord(movie)
	createdMovie := db.Create(&movie).Debug()
	var errMessage = createdMovie.Error

	if createdMovie.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdMovie)
}

// GetMovieCollection : Récupére la totalité des films dans la base de données
func GetMovieCollection(w http.ResponseWriter, r *http.Request) {
	if (r).Method == "OPTIONS" {
		return
	}
	var movies []models.Movie

	allMovies := db.Find(&movies)
	var errMessage = allMovies.Error
	if allMovies.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(allMovies)

}

// DeleteMovie : Delete a movie with id param
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	if (r).Method == "OPTIONS" {
		return
	}
	params := mux.Vars(r)
	var id = params["id"]

	var movie models.Movie
	fmt.Println(id)
	db.Where("api_movie_id = ?", id).Unscoped().Delete(movie)

	json.NewEncoder(w).Encode("Movie deleted")
}
