package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dogetoing!\n"))
}

func moviesHandler(w http.ResponseWriter, r *http.Request) {

	//select name, avg(r.score) from movie m left outer join movie_review r on (m.id = r.movie_id) group by m.id;
	// query := "SELECT id, name, description, image_url, release_date FROM movie"
	query := "SELECT m.id, m.name, m.description, m.image_url, m.release_date, COALESCE(avg(r.score), -1) FROM movie m LEFT OUTER JOIN movie_review r ON (m.id = r.movie_id) GROUP BY m.id ORDER BY m.id"
	rows, err := db.Query(context.Background(), query)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in MovieHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("Movies query succesfull")

	var movies []*movie = make([]*movie, 0)

	for rows.Next() {
		m := &movie{}
		err = rows.Scan(&m.Id, &m.Name, &m.Description, &m.ImageUrl, &m.ReleaseDate, &m.Score)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		movies = append(movies, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {

	query := "SELECT g.id, g.name, g.description, g.image_url, g.release_date, COALESCE(avg(r.score), -1) FROM game g LEFT OUTER JOIN game_review r ON (g.id = r.game_id) GROUP BY g.id ORDER BY g.id"
	rows, err := db.Query(context.Background(), query)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in gamesHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("Games query succesfull")

	var games []*game = make([]*game, 0)

	for rows.Next() {
		g := &game{}
		err = rows.Scan(&g.Id, &g.Name, &g.Description, &g.ImageUrl, &g.ReleaseDate, &g.Score)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		games = append(games, g)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func showsHandler(w http.ResponseWriter, r *http.Request) {

	query := "SELECT s.id, s.name, s.description, s.image_url, s.release_date, COALESCE(avg(r.score), -1) FROM show s LEFT OUTER JOIN show_review r ON (s.id = r.show_id) GROUP BY s.id ORDER BY s.id"
	rows, err := db.Query(context.Background(), query)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in showsHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("Games query succesfull")

	var shows []*show = make([]*show, 0)

	for rows.Next() {
		s := &show{}
		err = rows.Scan(&s.Id, &s.Name, &s.Description, &s.ImageUrl, &s.ReleaseDate, &s.Score)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		shows = append(shows, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shows)
}

func singleMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	m := movie{}
	query := "SELECT m.id, m.name, m.description, m.image_url, m.release_date, COALESCE(avg(r.score), -1) FROM movie m LEFT OUTER JOIN movie_review r ON (m.id = r.movie_id) WHERE m.id = $1 GROUP BY m.id"
	err := db.QueryRow(context.Background(), query, key).Scan(&m.Id, &m.Name, &m.Description, &m.ImageUrl, &m.ReleaseDate, &m.Score)

	if err == pgx.ErrNoRows {
		log.Printf("Movie query with id %v failed\n", key)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Query in singleMovieHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	log.Printf("Movie query with id %v succesfull\n", key)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}

func singleGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	g := game{}
	query := "SELECT g.id, g.name, g.description, g.image_url, g.release_date, COALESCE(avg(r.score), -1) FROM game g LEFT OUTER JOIN game_review r ON (g.id = r.game_id) WHERE g.id = $1 GROUP BY g.id"
	err := db.QueryRow(context.Background(), query, key).Scan(&g.Id, &g.Name, &g.Description, &g.ImageUrl, &g.ReleaseDate, &g.Score)

	if err == pgx.ErrNoRows {
		log.Printf("Game query with id %v failed\n", key)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Query in singleGameHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	log.Printf("Game query with id %v succesfull\n", key)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g)
}

func singleShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	s := show{}
	query := "SELECT s.id, s.name, s.description, s.image_url, s.release_date, COALESCE(avg(r.score), -1) FROM show s LEFT OUTER JOIN show_review r ON (s.id = r.show_id) WHERE s.id = $1 GROUP BY s.id"
	err := db.QueryRow(context.Background(), query, key).Scan(&s.Id, &s.Name, &s.Description, &s.ImageUrl, &s.ReleaseDate, &s.Score)

	if err == pgx.ErrNoRows {
		log.Printf("Show query with id %v failed\n", key)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Query in singleShowHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	log.Printf("Show query with id %v succesfull\n", key)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func addMovieHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached addMovieHandler")
	decoder := json.NewDecoder(r.Body)

	var t struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ImageUrl    string `json:"imageURL"`
		ReleaseDate string `json:"releaseDate"`
	}
	err := decoder.Decode(&t)
	if err != nil {
		log.Printf("Error decoding movie json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}
	d, err := time.Parse("2006-1-2", t.ReleaseDate)
	if err != nil {
		log.Printf("Error decoding date from json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}

	m := movie{Name: strings.ToLower(t.Name), Description: t.Description, ImageUrl: t.ImageUrl, ReleaseDate: d}
	query := "INSERT INTO movie (name, release_date, description, image_url) VALUES ($1, $2, $3, $4) RETURNING id"
	err = db.QueryRow(context.Background(), query, m.Name, m.ReleaseDate, m.Description, m.ImageUrl).Scan(&m.Id)
	if err != nil {
		log.Printf("Error inserting value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error adding movie", http.StatusBadRequest)
		} else {
			http.Error(w, "Error adding movie", http.StatusInternalServerError)

		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(m)
}

func addGameHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached addGameHandler")
	decoder := json.NewDecoder(r.Body)

	var t struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ImageUrl    string `json:"imageURL"`
		ReleaseDate string `json:"releaseDate"`
	}
	err := decoder.Decode(&t)
	if err != nil {
		log.Printf("Error decoding game json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}
	d, err := time.Parse("2006-1-2", t.ReleaseDate)
	if err != nil {
		log.Printf("Error decoding date from json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}

	g := game{Name: strings.ToLower(t.Name), Description: t.Description, ImageUrl: t.ImageUrl, ReleaseDate: d}
	query := "INSERT INTO game (name, release_date, description, image_url) VALUES ($1, $2, $3, $4) RETURNING id"
	err = db.QueryRow(context.Background(), query, g.Name, g.ReleaseDate, g.Description, g.ImageUrl).Scan(&g.Id)
	if err != nil {

		log.Printf("Error inserting value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error adding game", http.StatusBadRequest)
		} else {
			http.Error(w, "Error adding game", http.StatusInternalServerError)

		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(g)
}

func addShowHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached addShowHandler")
	decoder := json.NewDecoder(r.Body)

	var t struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ImageUrl    string `json:"imageURL"`
		ReleaseDate string `json:"releaseDate"`
	}
	err := decoder.Decode(&t)
	if err != nil {
		log.Printf("Error decoding show json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}
	d, err := time.Parse("2006-1-2", t.ReleaseDate)
	if err != nil {
		log.Printf("Error decoding date from json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}

	s := show{Name: strings.ToLower(t.Name), Description: t.Description, ImageUrl: t.ImageUrl, ReleaseDate: d}
	query := "INSERT INTO show (name, release_date, description, image_url) VALUES ($1, $2, $3, $4) RETURNING id"
	err = db.QueryRow(context.Background(), query, s.Name, s.ReleaseDate, s.Description, s.ImageUrl).Scan(&s.Id)
	if err != nil {
		log.Printf("Error inserting value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error adding show", http.StatusBadRequest)
		} else {
			http.Error(w, "Error adding show", http.StatusInternalServerError)

		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached UserHandler")

	query := "SELECT u.id, u.name, u.register_date FROM user u ORDER BY u.id"
	rows, err := db.Query(context.Background(), query)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("Movies query succesfull")

	var users []*user = make([]*user, 0)

	for rows.Next() {
		u := &user{}
		err = rows.Scan(&u.Id, &u.Name, &u.RegisterDate)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached addUserHandler")
	decoder := json.NewDecoder(r.Body)

	var u user
	decoder.Decode(&u)

	query := "INSERT INTO user (id, name) VALUES ($1, $2) RETURNING register_date"
	err := db.QueryRow(context.Background(), query, u.Id, u.Name).Scan(&u.RegisterDate)
	if err != nil {
		log.Printf("Error inserting value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error adding user", http.StatusBadRequest)
		} else {
			http.Error(w, "Error adding user", http.StatusInternalServerError)

		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}
