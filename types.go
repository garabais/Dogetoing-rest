package main

import "time"

type user struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	RegisterDate time.Time `json:"registerDate"`
}

type movie struct {
	Id          int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageURL"`
	ReleaseDate time.Time `json:"releaseDate"`
}

type game struct {
	Id          int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageURL"`
	ReleaseDate time.Time `json:"releaseDate"`
}

type show struct {
	Id          int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageURL"`
	ReleaseDate time.Time `json:"releaseDate"`
}

type movieReview struct {
	UserId string `json:"userId"`
	Movie  movie  `json:"movie"`
	Score  string `json:"score"`
}
type gameReview struct {
	UserId string `json:"userId"`
	Game   game   `json:"game"`
	Score  string `json:"score"`
}
type showReview struct {
	UserId string `json:"userId"`
	Show   show   `json:"show"`
	Score  string `json:"score"`
}
