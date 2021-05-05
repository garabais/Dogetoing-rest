package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dogetoing!\n"))
}

func moviesHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query(context.Background(), "SELECT id, name, description, image_url, release_date FROM movie")
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in MovieHandler failed: %v", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("Movies query succesfull")

	var movies []*movie = make([]*movie, 0)

	for i := 0; rows.Next(); i++ {
		m := &movie{}
		err = rows.Scan(&m.Id, &m.Name, &m.Description, &m.ImageUrl, &m.ReleaseDate)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		movies = append(movies, m)
	}

	json.NewEncoder(w).Encode(movies)
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query(context.Background(), "SELECT id, name, description, image_url, release_date FROM game")
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in gamesHandler failed: %v", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("Games query succesfull")

	var games []*game = make([]*game, 0)

	for i := 0; rows.Next(); i++ {
		g := &game{}
		err = rows.Scan(&g.Id, &g.Name, &g.Description, &g.ImageUrl, &g.ReleaseDate)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		games = append(games, g)
	}

	json.NewEncoder(w).Encode(games)
}

func showsHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query(context.Background(), "SELECT id, name, description, image_url, release_date FROM show")
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in showsHandler failed: %v", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("Games query succesfull")

	var shows []*show = make([]*show, 0)

	for i := 0; rows.Next(); i++ {
		s := &show{}
		err = rows.Scan(&s.Id, &s.Name, &s.Description, &s.ImageUrl, &s.ReleaseDate)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		shows = append(shows, s)
	}

	json.NewEncoder(w).Encode(shows)
}


func singleMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	m := movie{}
	err := db.QueryRow(context.Background(), "SELECT id, name, description, image_url, release_date FROM movie WHERE id = $1", key).Scan(&m.Id, &m.Name, &m.Description, &m.ImageUrl, &m.ReleaseDate)

	if err == pgx.ErrNoRows {
		log.Printf("Movie query with id %v failed", key)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Query in singleMovieHandler failed: %v", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	log.Printf("Movie query with id %v succesfull", key)

	json.NewEncoder(w).Encode(m)
}

func singleGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	g := game{}
	err := db.QueryRow(context.Background(), "SELECT id, name, description, image_url, release_date FROM game WHERE id = $1", key).Scan(&g.Id, &g.Name, &g.Description, &g.ImageUrl, &g.ReleaseDate)

	if err == pgx.ErrNoRows {
		log.Printf("Game query with id %v failed", key)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Query in singleGameHandler failed: %v", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	log.Printf("Game query with id %v succesfull", key)

	json.NewEncoder(w).Encode(g)
}

func singleShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	s := show{}
	err := db.QueryRow(context.Background(), "SELECT id, name, description, image_url, release_date FROM show WHERE id = $1", key).Scan(&s.Id, &s.Name, &s.Description, &s.ImageUrl, &s.ReleaseDate)

	if err == pgx.ErrNoRows {
		log.Printf("Show query with id %v failed", key)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Query in singleShowHandler failed: %v", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	log.Printf("Show query with id %v succesfull", key)

	json.NewEncoder(w).Encode(s)
}
