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

type Review struct {
	UserId  string `json:"uid"`
	Id int    `json:"id"`
	Score   int    `json:"score"`
}

type follow struct {
	Follower  string `json:"uid"`
	Following string `json:"followUid"`
}
