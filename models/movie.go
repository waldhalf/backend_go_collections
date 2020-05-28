package models

import (
	"github.com/jinzhu/gorm"
)

//Movie struct declaration
type Movie struct {
	gorm.Model

	Title       string `json:"title"`
	Overview    string `json:"overview"`
	ReleaseDate string `json:"release_date"`
	PosterPath  string `json:"poster_path"`
	APIMovieID  int    `json:"id" gorm:"unique;not null" `
}
