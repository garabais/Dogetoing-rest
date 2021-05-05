package main

import (
	"context"
	"database/sql"

	"encoding/json"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dogetoing!\n"))
}

func moviesHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query(context.Background(), "SELECT id, name, description, image_url, release_date FROM movie")
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Query failed", http.StatusInternalServerError)
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
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Query failed", http.StatusInternalServerError)
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
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Query failed", http.StatusInternalServerError)
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

// func singleArticleHandler(w http.ResponseWriter, r *http.Request) {
//     vars := mux.Vars(r)
//     key := vars["id"]

//     for _, a := range articles {
//         if a.Id == key {
//             json.NewEncoder(w).Encode(a)
//             return
//         }
//     }
//     http.Error(w, "Item not found", http.StatusNoContent)
// }
