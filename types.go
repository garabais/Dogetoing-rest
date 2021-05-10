package main

import "time"

type user struct {
	Id           string    `json:"uid"`
	Name         string    `json:"name"`
	RegisterDate time.Time `json:"registerDate"`
}

type movie struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageURL"`
	ReleaseDate time.Time `json:"releaseDate"`
	Score       float64   `json:"score"`
}

type game struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageURL"`
	ReleaseDate time.Time `json:"releaseDate"`
	Score       float64   `json:"score"`
}

type show struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageURL"`
	ReleaseDate time.Time `json:"releaseDate"`
	Score       float64   `json:"score"`
}

type movieReview struct {
	UserId string `json:"userId"`
	Movie  movie  `json:"movie"`
	Score  string `json:"score"`
}
type gameReview struct {
	UserId string `json:"uid"`
	Game   game   `json:"game"`
	Score  string `json:"score"`
}
type showReview struct {
	UserId string `json:"uid"`
	Show   show   `json:"show"`
	Score  string `json:"score"`
}
