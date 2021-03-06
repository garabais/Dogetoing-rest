package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Endpoints")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "Movies")
	fmt.Fprintln(w, "\tGET:  /movies")
	fmt.Fprintln(w, "\tGET:  /movies?name={name}")
	fmt.Fprintln(w, "\tPOST: /movies")
	fmt.Fprintln(w, "\tGET:  /movies/{id}")
	fmt.Fprintln(w, "\tPUT:  /movies/{id}")
	fmt.Fprintln(w, "\tDELETE:  /movies/{id}")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "Games")
	fmt.Fprintln(w, "\tGET:  /games")
	fmt.Fprintln(w, "\tGET:  /games?name={name}")
	fmt.Fprintln(w, "\tPOST: /games")
	fmt.Fprintln(w, "\tGET:  /games/{id}")
	fmt.Fprintln(w, "\tPUT:  /games/{id}")
	fmt.Fprintln(w, "\tDELETE:  /games/{id}")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "Shows")
	fmt.Fprintln(w, "\tGET:  /shows")
	fmt.Fprintln(w, "\tGET:  /shows?name={name}")
	fmt.Fprintln(w, "\tPOST: /shows")
	fmt.Fprintln(w, "\tGET:  /shows/{id}")
	fmt.Fprintln(w, "\tPUT:  /shows/{id}")
	fmt.Fprintln(w, "\tDELETE:  /shows/{id}")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "Users")
	fmt.Fprintln(w, "\tGET:  /users")
	fmt.Fprintln(w, "\tPOST: /users")
	fmt.Fprintln(w, "\tGET:  /users/{uid}")
	fmt.Fprintln(w, "\tPUT:  /users/{uid}")
	fmt.Fprintln(w, "\tGET:  /users/{uid}/feed/movies")
	fmt.Fprintln(w, "\tGET:  /users/{uid}/feed/games")
	fmt.Fprintln(w, "\tGET:  /users/{uid}/feed/shows")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "User Movies")
	fmt.Fprintln(w, "\tGET:     /users/{uid}/movies")
	fmt.Fprintln(w, "\tGET:     /users/{uid}/movies?name={name}")
	fmt.Fprintln(w, "\tPOST:    /users/{uid}/movies")
	fmt.Fprintln(w, "\tGET:     /users/{uid}/movies/{id}")
	fmt.Fprintln(w, "\tDELETE:  /users/{uid}/movies/{id}")
	fmt.Fprintln(w, "\tPUT:     /users/{uid}/movies/{id}")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "User Games")
	fmt.Fprintln(w, "\tGET:     /users/{uid}/games")
	fmt.Fprintln(w, "\tGET:     /users/{uid}/games?name={name}")
	fmt.Fprintln(w, "\tPOST:    /users/{uid}/games")
	fmt.Fprintln(w, "\tGET:     /users/{uid}/games/{id}")
	fmt.Fprintln(w, "\tDELETE:  /users/{uid}/games/{id}")
	fmt.Fprintln(w, "\tPUT:     /users/{uid}/games/{id}")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "User Shows")
	fmt.Fprintln(w, "\tGET:     /users/{uid}/shows")
	fmt.Fprintln(w, "\tGET:     /users/{uid}/shows?name={name}")
	fmt.Fprintln(w, "\tPOST:    /users/{uid}/shows")
	fmt.Fprintln(w, "\tGET:     /users/{uid}/shows/{id}")
	fmt.Fprintln(w, "\tDELETE:  /users/{uid}/shows/{id}")
	fmt.Fprintln(w, "\tPUT:     /users/{uid}/shows/{id}")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "User Follows")
	fmt.Fprintln(w, "\tGET:     /users/{uid}/follows")
	fmt.Fprintln(w, "\tPOST:    /users/{uid}/follows")
	fmt.Fprintln(w, "\tGET:     /users/{uid}/follows/{id}")
	fmt.Fprintln(w, "\tDELETE:  /users/{uid}/follows/{id}")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "User Followers")
	fmt.Fprintln(w, "\tGET:     /users/{uid}/followers")
	fmt.Fprintln(w, "\tGET:     /users/{uid}/followers/{id}")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "POST EXAMPLES")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "JSON add movie/game/show")
	fmt.Fprintln(w, "\t{")
	fmt.Fprintln(w, "\t\t\"name\":\"toy story\",")
	fmt.Fprintln(w, "\t\t\"description\":\"Toys\",")
	fmt.Fprintln(w, "\t\t\"imageURL\":\"https://url.com/imagen.png\",")
	fmt.Fprintln(w, "\t\t\"releaseDate\":\"2016-09-17\"")
	fmt.Fprintln(w, "\t}")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "JSON add user")
	fmt.Fprintln(w, "\t{")
	fmt.Fprintln(w, "\t\t\"uid\":\"Jyu2oXXi9XQZf2CJz6ZWyeNycAB2\",")
	fmt.Fprintln(w, "\t\t\"name\":\"Ari\"")
	fmt.Fprintln(w, "\t}")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "JSON add follower")
	fmt.Fprintln(w, "\t{")
	fmt.Fprintln(w, "\t\t\"followUid\":\"Jyu2oXXi9XQZf2CJz6ZWyeNycAB2\"")
	fmt.Fprintln(w, "\t}")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "JSON add user game/movie/show review")
	fmt.Fprintln(w, "\t{")
	fmt.Fprintln(w, "\t\t\"id\": 1,")
	fmt.Fprintln(w, "\t\t\"score\":10")
	fmt.Fprintln(w, "\t}")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "PUT EXAMPLES")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "JSON update user game/movie/show review")
	fmt.Fprintln(w, "\t{")
	fmt.Fprintln(w, "\t\t\"score\":10")
	fmt.Fprintln(w, "\t}")
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

	log.Print("Shows query succesfull")

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

func nameQueryMoviesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	query := "SELECT m.id, m.name, m.description, m.image_url, m.release_date, COALESCE(avg(r.score), -1) FROM movie m LEFT OUTER JOIN movie_review r ON (m.id = r.movie_id) WHERE name ~ $1 GROUP BY m.id ORDER BY m.id"
	rows, err := db.Query(context.Background(), query, name)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in nameQueryMovieHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("NameQueryMovies query succesfull")

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

func nameQueryGamesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	query := "SELECT g.id, g.name, g.description, g.image_url, g.release_date, COALESCE(avg(r.score), -1) FROM game g LEFT OUTER JOIN game_review r ON (g.id = r.game_id) WHERE name ~ $1 GROUP BY g.id ORDER BY g.id"
	rows, err := db.Query(context.Background(), query, name)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in nameQueryGamesHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("NameQueryGames query succesfull")

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

func nameQueryShowsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	query := "SELECT s.id, s.name, s.description, s.image_url, s.release_date, COALESCE(avg(r.score), -1) FROM show s LEFT OUTER JOIN show_review r ON (s.id = r.show_id) WHERE name ~ $1 GROUP BY s.id ORDER BY s.id"
	rows, err := db.Query(context.Background(), query, name)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in nameQueryShowsHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("NameQueryShows query succesfull")

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
	id := vars["id"]

	m := movie{}
	query := "SELECT m.id, m.name, m.description, m.image_url, m.release_date, COALESCE(avg(r.score), -1) FROM movie m LEFT OUTER JOIN movie_review r ON (m.id = r.movie_id) WHERE m.id = $1 GROUP BY m.id"
	err := db.QueryRow(context.Background(), query, id).Scan(&m.Id, &m.Name, &m.Description, &m.ImageUrl, &m.ReleaseDate, &m.Score)

	if err == pgx.ErrNoRows {
		log.Printf("Movie query with id %v failed\n", id)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Query in singleMovieHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	log.Printf("Movie query with id %v succesfull\n", id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}

func singleGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	g := game{}
	query := "SELECT g.id, g.name, g.description, g.image_url, g.release_date, COALESCE(avg(r.score), -1) FROM game g LEFT OUTER JOIN game_review r ON (g.id = r.game_id) WHERE g.id = $1 GROUP BY g.id"
	err := db.QueryRow(context.Background(), query, id).Scan(&g.Id, &g.Name, &g.Description, &g.ImageUrl, &g.ReleaseDate, &g.Score)

	if err == pgx.ErrNoRows {
		log.Printf("Game query with id %v failed\n", id)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Query in singleGameHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	log.Printf("Game query with id %v succesfull\n", id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g)
}

func singleShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	s := show{}
	query := "SELECT s.id, s.name, s.description, s.image_url, s.release_date, COALESCE(avg(r.score), -1) FROM show s LEFT OUTER JOIN show_review r ON (s.id = r.show_id) WHERE s.id = $1 GROUP BY s.id"
	err := db.QueryRow(context.Background(), query, id).Scan(&s.Id, &s.Name, &s.Description, &s.ImageUrl, &s.ReleaseDate, &s.Score)

	if err == pgx.ErrNoRows {
		log.Printf("Show query with id %v failed\n", id)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Query in singleShowHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	log.Printf("Show query with id %v succesfull\n", id)

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
	log.Println(t)
	if err != nil {
		log.Printf("Error decoding movie json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}
	if t.Id != 0 || t.Name == "" || t.Description == "" || t.ImageUrl == "" || t.ReleaseDate == "" {
		log.Printf("Error incomplete json: %v\n", err)
		http.Error(w, "Unable incomplete json", http.StatusBadRequest)
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
	if t.Id != 0 || t.Name == "" || t.Description == "" || t.ImageUrl == "" || t.ReleaseDate == "" {
		log.Printf("Error incomplete json: %v\n", err)
		http.Error(w, "Unable incomplete json", http.StatusBadRequest)
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
	if t.Id != 0 || t.Name == "" || t.Description == "" || t.ImageUrl == "" || t.ReleaseDate == "" {
		log.Printf("Error incomplete json: %v\n", err)
		http.Error(w, "Unable incomplete json", http.StatusBadRequest)
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

func deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Printf("Delete Movie with id %v\n", id)

	query := "DELETE FROM movie WHERE id = $1"
	_, err := db.Exec(context.Background(), query, id)

	if err != nil {
		log.Printf("Error deleting value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error deleting movie", http.StatusBadRequest)
		} else {
			http.Error(w, "Error deleting movie", http.StatusInternalServerError)

		}
		return
	}

	log.Printf("Delete succesfull id %v succesfull\n", id)

	w.WriteHeader(http.StatusNoContent)
}

func deleteGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Printf("Delete game with id %v\n", id)

	query := "DELETE FROM game WHERE id = $1"
	_, err := db.Exec(context.Background(), query, id)

	if err != nil {
		log.Printf("Error deleting value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error deleting game", http.StatusBadRequest)
		} else {
			http.Error(w, "Error deleting game", http.StatusInternalServerError)

		}
		return
	}

	log.Printf("Delete succesfull id %v succesfull\n", id)

	w.WriteHeader(http.StatusNoContent)
}

func deleteShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Printf("Delete show with id %v\n", id)

	query := "DELETE FROM show WHERE id = $1"
	_, err := db.Exec(context.Background(), query, id)

	if err != nil {
		log.Printf("Error deleting value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error deleting show", http.StatusBadRequest)
		} else {
			http.Error(w, "Error deleting show", http.StatusInternalServerError)

		}
		return
	}

	log.Printf("Delete succesfull id %v succesfull\n", id)

	w.WriteHeader(http.StatusNoContent)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached UserHandler")

	query := "SELECT u.id, u.name, u.register_date, u.is_admin FROM account u ORDER BY u.register_date"
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
		err = rows.Scan(&u.Id, &u.Name, &u.RegisterDate, &u.Admin)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func nameQueryUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached UserHandler")

	vars := mux.Vars(r)
	name := vars["name"]

	name = strings.ToLower(name)

	query := "SELECT u.id, u.name, u.register_date, u.is_admin FROM account u WHERE name ~ $1 ORDER BY u.register_date"
	rows, err := db.Query(context.Background(), query, name)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("User name query succesfull")

	var users []*user = make([]*user, 0)

	for rows.Next() {
		u := &user{}
		err = rows.Scan(&u.Id, &u.Name, &u.RegisterDate, &u.Admin)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func nameAdminQueryUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached UserHandler")

	vars := mux.Vars(r)
	name := vars["name"]
	ad := vars["admin"]

	var admin bool

	if strings.ToLower(ad) == "true" {
		admin = true
	}

	name = strings.ToLower(name)

	query := "SELECT u.id, u.name, u.register_date, u.is_admin FROM account u WHERE name ~ $1 AND u.is_admin = $2 ORDER BY u.register_date"
	rows, err := db.Query(context.Background(), query, name, admin)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("User name Admin query succesfull")

	var users []*user = make([]*user, 0)

	for rows.Next() {
		u := &user{}
		err = rows.Scan(&u.Id, &u.Name, &u.RegisterDate, &u.Admin)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func nameFollowQueryUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached UserHandler")

	vars := mux.Vars(r)
	name := vars["name"]
	uid := vars["uid"]


	name = strings.ToLower(name)
	query := "SELECT * FROM account WHERE name ~ $1 AND NOT id IN (SELECT following_id AS id from follow WHERE follower_id = $2)"
	// query := "SELECT a.id, a.name, a.register_date, a.is_admin FROM follow f, account a WHERE a.id = f.following_id AND name ~ $1 AND a.id != $2 GROUP BY a.id"
	// query := "SELECT a.id, a.name, a.register_date, a.is_admin FROM follow f, account a WHERE a.id = f.following_id AND name ~ $1 and f.follower_id != $2"
	// query := "SELECT u.id, u.name, u.register_date, u.is_admin FROM account u WHERE name ~ $1 AND u.is_admin = $2 ORDER BY u.register_date"
	rows, err := db.Query(context.Background(), query, name, uid)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("User name follow query succesfull")

	var users []*user = make([]*user, 0)

	for rows.Next() {
		u := &user{}
		err = rows.Scan(&u.Id, &u.Name, &u.RegisterDate, &u.Admin)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}


func noFollowQueryUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached UserHandler")

	vars := mux.Vars(r)
	// name := vars["name"]
	uid := vars["uid"]


	query := "SELECT * FROM account WHERE NOT id IN (SELECT following_id AS id from follow WHERE follower_id = $1)"
	// query := "SELECT a.id, a.name, a.register_date, a.is_admin FROM follow f, account a WHERE a.id = f.following_id AND name ~ $1 AND a.id != $2 GROUP BY a.id"
	// query := "SELECT a.id, a.name, a.register_date, a.is_admin FROM follow f, account a WHERE a.id = f.following_id AND name ~ $1 and f.follower_id != $2"
	// query := "SELECT u.id, u.name, u.register_date, u.is_admin FROM account u WHERE name ~ $1 AND u.is_admin = $2 ORDER BY u.register_date"
	rows, err := db.Query(context.Background(), query, uid)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("User name follow query succesfull")

	var users []*user = make([]*user, 0)

	for rows.Next() {
		u := &user{}
		err = rows.Scan(&u.Id, &u.Name, &u.RegisterDate, &u.Admin)
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

	u.Name = strings.ToLower(u.Name)

	query := "INSERT INTO account (id, name) VALUES ($1, $2) RETURNING register_date"
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

func singleUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached SingleUserHandler")

	vars := mux.Vars(r)
	uid := vars["uid"]

	u := user{}

	query := "SELECT u.id, u.name, u.register_date, u.is_admin FROM account u WHERE u.id = $1"
	err := db.QueryRow(context.Background(), query, uid).Scan(&u.Id, &u.Name, &u.RegisterDate, &u.Admin)

	if err == pgx.ErrNoRows {
		log.Printf("User query with id %v failed\n", uid)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Query in singleMovieHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}

	log.Printf("User query with id %v succesfull\n", uid)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)

}

func userMoviesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	query := "SELECT m.id, m.name, m.description, m.image_url, m.release_date, r.score FROM movie m LEFT OUTER JOIN movie_review r ON (m.id = r.movie_id) WHERE r.account_id = $1 ORDER BY r.score DESC"
	rows, err := db.Query(context.Background(), query, uid)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserMovieHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("UserMovies query succesfull")

	var movies []*movie = make([]*movie, 0)

	for rows.Next() {
		m := &movie{}
		var score int
		err = rows.Scan(&m.Id, &m.Name, &m.Description, &m.ImageUrl, &m.ReleaseDate, &score)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		m.Score = float64(score)
		movies = append(movies, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func userGamesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	query := "SELECT g.id, g.name, g.description, g.image_url, g.release_date, r.score FROM game g LEFT OUTER JOIN game_review r ON (g.id = r.game_id) WHERE r.account_id = $1 ORDER BY r.score DESC"
	rows, err := db.Query(context.Background(), query, uid)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserGamesHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("UserGames query succesfull")

	var games []*game = make([]*game, 0)

	for rows.Next() {
		g := &game{}
		var score int
		err = rows.Scan(&g.Id, &g.Name, &g.Description, &g.ImageUrl, &g.ReleaseDate, &score)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		g.Score = float64(score)
		games = append(games, g)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func userShowsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	query := "SELECT s.id, s.name, s.description, s.image_url, s.release_date, r.score FROM show s LEFT OUTER JOIN show_review r ON (s.id = r.show_id) WHERE r.account_id = $1 ORDER BY r.score DESC"
	rows, err := db.Query(context.Background(), query, uid)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserShowsHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("UserShows query succesfull")

	var shows []*show = make([]*show, 0)

	for rows.Next() {
		s := &show{}
		var score int
		err = rows.Scan(&s.Id, &s.Name, &s.Description, &s.ImageUrl, &s.ReleaseDate, &score)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		s.Score = float64(score)
		shows = append(shows, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shows)
}

func nameQueryUserMoviesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

	query := "SELECT m.id, m.name, m.description, m.image_url, m.release_date, r.score FROM movie m LEFT OUTER JOIN movie_review r ON (m.id = r.movie_id) WHERE r.account_id = $1 AND name ~ $2 ORDER BY r.score DESC"
	rows, err := db.Query(context.Background(), query, uid, name)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in nameQueryUserMovieHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("NameQueryUserMovies query succesfull")

	var movies []*movie = make([]*movie, 0)

	for rows.Next() {
		m := &movie{}
		var score int
		err = rows.Scan(&m.Id, &m.Name, &m.Description, &m.ImageUrl, &m.ReleaseDate, &score)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		m.Score = float64(score)
		movies = append(movies, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func nameQueryUserGamesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

	query := "SELECT g.id, g.name, g.description, g.image_url, g.release_date, r.score FROM game g LEFT OUTER JOIN game_review r ON (g.id = r.game_id) WHERE r.account_id = $1 AND name ~ $2 ORDER BY r.score DESC"

	rows, err := db.Query(context.Background(), query, uid, name)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in nameQueryUserGamesHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("NameQueryUserGames query succesfull")

	var games []*game = make([]*game, 0)

	for rows.Next() {
		g := &game{}
		var score int
		err = rows.Scan(&g.Id, &g.Name, &g.Description, &g.ImageUrl, &g.ReleaseDate, &score)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		g.Score = float64(score)
		games = append(games, g)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func nameQueryUserShowsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

	query := "SELECT s.id, s.name, s.description, s.image_url, s.release_date, r.score FROM show s LEFT OUTER JOIN show_review r ON (s.id = r.show_id) WHERE r.account_id = $1 AND name ~ $2 ORDER BY r.score DESC"
	rows, err := db.Query(context.Background(), query, uid, name)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in nameQueryUserShowsHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("NameQueryUserShows query succesfull")

	var shows []*show = make([]*show, 0)

	for rows.Next() {
		s := &show{}
		var score int
		err = rows.Scan(&s.Id, &s.Name, &s.Description, &s.ImageUrl, &s.ReleaseDate, &score)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		s.Score = float64(score)
		shows = append(shows, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shows)
}

func userFollowsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	query := "SELECT f.follower_id, f.following_id, a.name FROM follow f, account a WHERE f.follower_id = $1 and a.id = f.following_id"
	// query := "SELECT f.follower_id, f.following_id FROM follow f WHERE f.follower_id = $1"
	// query := "SELECT f.follower_id, f.following_id FROM follow f WHERE f.follower_id = $1"
	rows, err := db.Query(context.Background(), query, uid)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserShowsHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("UserShows query succesfull")

	var follows []*follow = make([]*follow, 0)

	for rows.Next() {
		f := &follow{}
		err = rows.Scan(&f.Follower, &f.Following, &f.Name)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		follows = append(follows, f)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(follows)
}

func nameUserFollowsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

	query := "SELECT f.follower_id, f.following_id, a.name FROM follow f, account a WHERE f.follower_id = $1 and a.id = f.following_id and name ~ $2"
	// query := "SELECT f.follower_id, f.following_id FROM follow f WHERE f.follower_id = $1"
	// query := "SELECT f.follower_id, f.following_id FROM follow f WHERE f.follower_id = $1"
	rows, err := db.Query(context.Background(), query, uid, name)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserShowsHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("UserShows query succesfull")

	var follows []*follow = make([]*follow, 0)

	for rows.Next() {
		f := &follow{}
		err = rows.Scan(&f.Follower, &f.Following, &f.Name)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		follows = append(follows, f)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(follows)
}

func userFollowersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	// query := "SELECT f.follower_id, f.following_id, a.name FROM follow f, account a WHERE f.follower_id = $1 and a.id = f.following_id"
	query := "SELECT f.follower_id, f.following_id, a.name FROM follow f, account a WHERE f.following_id = $1 and a.id = follower_id"
	rows, err := db.Query(context.Background(), query, uid)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserShowsHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("UserShows query succesfull")

	var follows []*follow = make([]*follow, 0)

	for rows.Next() {
		f := &follow{}
		err = rows.Scan(&f.Follower, &f.Following, &f.Name)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		follows = append(follows, f)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(follows)
}

func nameUserFollowersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	name := vars["name"]

	// query := "SELECT f.follower_id, f.following_id, a.name FROM follow f, account a WHERE f.follower_id = $1 and a.id = f.following_id"
	query := "SELECT f.follower_id, f.following_id, a.name FROM follow f, account a WHERE f.following_id = $1 and a.id = follower_id and name ~ $2"
	rows, err := db.Query(context.Background(), query, uid, name)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserShowsHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("UserShows query succesfull")

	var follows []*follow = make([]*follow, 0)

	for rows.Next() {
		f := &follow{}
		err = rows.Scan(&f.Follower, &f.Following, &f.Name)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		follows = append(follows, f)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(follows)
}

func singleUserMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid := vars["uid"]

	m := movie{}
	var score int
	query := "SELECT m.id, m.name, m.description, m.image_url, m.release_date, r.score FROM movie m LEFT OUTER JOIN movie_review r ON (m.id = r.movie_id) WHERE m.id = $1 AND r.account_id = $2"
	err := db.QueryRow(context.Background(), query, id, uid).Scan(&m.Id, &m.Name, &m.Description, &m.ImageUrl, &m.ReleaseDate, &score)

	if err == pgx.ErrNoRows {
		log.Printf("UserMovie query with uid %v and id %v failed\n", uid, id)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Query in singleUserMovieHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	m.Score = float64(score)
	log.Printf("SingleUserMovie query with uid %v and id %v succesfull\n", uid, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}

func singleUserGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid := vars["uid"]

	g := game{}
	var score int
	query := "SELECT g.id, g.name, g.description, g.image_url, g.release_date, r.score FROM game g LEFT OUTER JOIN game_review r ON (g.id = r.game_id) WHERE g.id = $1 and r.account_id = $2"
	err := db.QueryRow(context.Background(), query, id, uid).Scan(&g.Id, &g.Name, &g.Description, &g.ImageUrl, &g.ReleaseDate, &score)

	if err == pgx.ErrNoRows {
		log.Printf("Game query with uid %v and id %v failed\n", uid, id)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Query in singleGameHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	g.Score = float64(score)
	log.Printf("Game query with uid %v and id %v succesfull\n", uid, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g)
}

func singleUserShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid := vars["uid"]

	s := show{}
	var score int
	query := "SELECT s.id, s.name, s.description, s.image_url, s.release_date, r.score FROM show s LEFT OUTER JOIN show_review r ON (s.id = r.show_id) WHERE s.id = $1 and r.account_id = $2"
	err := db.QueryRow(context.Background(), query, id, uid).Scan(&s.Id, &s.Name, &s.Description, &s.ImageUrl, &s.ReleaseDate, &score)

	if err == pgx.ErrNoRows {
		log.Printf("Show query with uid %v and id %v failed\n", uid, id)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Query in singleShowHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	s.Score = float64(score)
	log.Printf("Show query with uid %v and id %v succesfull\n", uid, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func singleUserFollowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid := vars["uid"]

	f := follow{}
	query := "SELECT f.follower_id, f.following_id, a.name FROM follow f, account a WHERE f.follower_id = $1 and f.following_id = $2 and a.id = f.following_id"
	err := db.QueryRow(context.Background(), query, uid, id).Scan(&f.Follower, &f.Following, &f.Name)

	if err == pgx.ErrNoRows {
		log.Printf("Follow query with uid %v and id %v failed\n", uid, id)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Follow in singleShowHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	log.Printf("Follow query with uid %v and id %v succesfull\n", uid, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(f)
}

func singleUserFollowersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid := vars["uid"]

	f := follow{}
	query := "SELECT f.follower_id, f.following_id FROM follow f WHERE f.follower_id = $1 and f.following_id = $2"
	err := db.QueryRow(context.Background(), query, id, uid).Scan(&f.Follower, &f.Following)

	if err == pgx.ErrNoRows {
		log.Printf("Follow query with uid %v and id %v failed\n", uid, id)
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Follow in singleShowHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	log.Printf("Follow query with uid %v and id %v succesfull\n", uid, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(f)
}

func addUserMovieHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached addUserMovieHandler")
	decoder := json.NewDecoder(r.Body)

	vars := mux.Vars(r)
	uid := vars["uid"]

	d := Review{}

	err := decoder.Decode(&d)
	if err != nil {
		log.Printf("Error decoding movie review json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}
	d.UserId = uid

	query := "INSERT INTO movie_review (account_id, movie_id, score) VALUES ($1, $2, $3)"
	_, err = db.Exec(context.Background(), query, d.UserId, d.Id, d.Score)
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
	json.NewEncoder(w).Encode(d)
}

func addUserGameHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached addUserGameHandler")
	decoder := json.NewDecoder(r.Body)

	vars := mux.Vars(r)
	uid := vars["uid"]

	d := Review{}

	err := decoder.Decode(&d)
	if err != nil {
		log.Printf("Error decoding movie review json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}
	d.UserId = uid

	query := "INSERT INTO game_review (account_id, game_id, score) VALUES ($1, $2, $3)"
	_, err = db.Exec(context.Background(), query, d.UserId, d.Id, d.Score)
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
	json.NewEncoder(w).Encode(d)
}

func addUserShowHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached addUserShowHandler")
	decoder := json.NewDecoder(r.Body)

	vars := mux.Vars(r)
	uid := vars["uid"]

	d := Review{}

	err := decoder.Decode(&d)
	if err != nil {
		log.Printf("Error decoding show review json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}
	d.UserId = uid

	query := "INSERT INTO show_review (account_id, show_id, score) VALUES ($1, $2, $3)"
	_, err = db.Exec(context.Background(), query, d.UserId, d.Id, d.Score)
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
	json.NewEncoder(w).Encode(d)
}

func addUserFollowHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached addUserFollowHandler")
	decoder := json.NewDecoder(r.Body)

	vars := mux.Vars(r)
	uid := vars["uid"]

	d := follow{}

	err := decoder.Decode(&d)
	if err != nil {
		log.Printf("Error decoding follow review json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}
	d.Follower = uid

	query := "INSERT INTO follow (follower_id, following_id) VALUES ($1, $2)"
	_, err = db.Exec(context.Background(), query, d.Follower, d.Following)
	if err != nil {
		log.Printf("Error inserting value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error adding follow", http.StatusBadRequest)
		} else {
			http.Error(w, "Error adding follow", http.StatusInternalServerError)

		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(d)
}

func deleteUserMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid := vars["uid"]

	log.Printf("Delete User Movie with uid %v and id %v\n", uid, id)

	query := "DELETE FROM movie_review WHERE account_id = $1 AND movie_id = $2"
	_, err := db.Exec(context.Background(), query, uid, id)

	if err != nil {
		log.Printf("Error deleting value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error deleting movie", http.StatusBadRequest)
		} else {
			http.Error(w, "Error deleting movie", http.StatusInternalServerError)

		}
		return
	}

	log.Printf("Delete succesfull with uid %v and id %v succesfull\n", uid, id)

	w.WriteHeader(http.StatusNoContent)
}

func deleteUserGamesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid := vars["uid"]

	log.Printf("Delete User game with uid %v and id %v\n", uid, id)

	query := "DELETE FROM game_review WHERE account_id = $1 AND game_id = $2"
	_, err := db.Exec(context.Background(), query, uid, id)

	if err != nil {
		log.Printf("Error deleting value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error deleting game", http.StatusBadRequest)
		} else {
			http.Error(w, "Error deleting game", http.StatusInternalServerError)

		}
		return
	}

	log.Printf("Delete succesfull with uid %v and id %v succesfull\n", uid, id)

	w.WriteHeader(http.StatusNoContent)
}

func deleteUserShowsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid := vars["uid"]

	log.Printf("Delete User show with uid %v and id %v\n", uid, id)

	query := "DELETE FROM show_review WHERE account_id = $1 AND show_id = $2"
	_, err := db.Exec(context.Background(), query, uid, id)

	if err != nil {
		log.Printf("Error deleting value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error deleting show", http.StatusBadRequest)
		} else {
			http.Error(w, "Error deleting show", http.StatusInternalServerError)

		}
		return
	}

	log.Printf("Delete succesfull with uid %v and id %v succesfull\n", uid, id)

	w.WriteHeader(http.StatusNoContent)
}

func deleteUserFollowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid := vars["uid"]

	log.Printf("Delete User follow with uid %v and id %v\n", uid, id)

	query := "DELETE FROM follow WHERE follower_id = $1 AND following_id = $2"
	_, err := db.Exec(context.Background(), query, uid, id)

	if err != nil {
		log.Printf("Error deleting value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error deleting follow", http.StatusBadRequest)
		} else {
			http.Error(w, "Error deleting follow", http.StatusInternalServerError)

		}
		return
	}

	log.Printf("Delete succesfull with uid %v and id %v succesfull\n", uid, id)

	w.WriteHeader(http.StatusNoContent)
}

func updateUserMovieHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached updateUserMovieHandler")
	decoder := json.NewDecoder(r.Body)

	vars := mux.Vars(r)
	uid := vars["uid"]
	id := vars["id"]

	d := Review{}

	err := decoder.Decode(&d)
	if err != nil {
		log.Printf("Error decoding movie review json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}
	d.UserId = uid
	d.Id, _ = strconv.Atoi(id)

	query := "UPDATE movie_review SET score = $1 WHERE account_id = $2 AND movie_id =$3"
	c, err := db.Exec(context.Background(), query, d.Score, d.UserId, d.Id)
	if err != nil {
		log.Printf("Error updating value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error updating movie", http.StatusBadRequest)
		} else {
			http.Error(w, "Error updating movie", http.StatusInternalServerError)

		}
		return
	} else if c.RowsAffected() == 0 {
		http.Error(w, "Review not found", http.StatusNotFound)
		log.Printf("Error updating: review with id %v not found for user %v\n", d.Id, d.UserId)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func updateUserGameHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached updateUserGameHandler")
	decoder := json.NewDecoder(r.Body)

	vars := mux.Vars(r)
	uid := vars["uid"]
	id := vars["id"]

	d := Review{}

	err := decoder.Decode(&d)
	if err != nil {
		log.Printf("Error decoding games review json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}
	d.UserId = uid
	d.Id, _ = strconv.Atoi(id)

	query := "UPDATE game_review SET score = $1 WHERE account_id = $2 AND game_id =$3"
	c, err := db.Exec(context.Background(), query, d.Score, d.UserId, d.Id)
	if err != nil {
		log.Printf("Error updating value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error updating game", http.StatusBadRequest)
		} else {
			http.Error(w, "Error updating game", http.StatusInternalServerError)

		}
		return
	} else if c.RowsAffected() == 0 {
		http.Error(w, "Review not found", http.StatusNotFound)
		log.Printf("Error updating: review with id %v not found for user %v\n", d.Id, d.UserId)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func updateUserShowsHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached updateUserShowsHandler")
	decoder := json.NewDecoder(r.Body)

	vars := mux.Vars(r)
	uid := vars["uid"]
	id := vars["id"]

	d := Review{}

	err := decoder.Decode(&d)
	if err != nil {
		log.Printf("Error decoding show review json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}
	d.UserId = uid
	d.Id, _ = strconv.Atoi(id)

	query := "UPDATE show_review SET score = $1 WHERE account_id = $2 AND show_id =$3"
	c, err := db.Exec(context.Background(), query, d.Score, d.UserId, d.Id)
	if err != nil {
		log.Printf("Error updating value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error updating show", http.StatusBadRequest)
		} else {
			http.Error(w, "Error updating show", http.StatusInternalServerError)

		}
		return
	} else if c.RowsAffected() == 0 {
		http.Error(w, "Review not found", http.StatusNotFound)
		log.Printf("Error updating: review with id %v not found for user %v\n", d.Id, d.UserId)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func userActivityMoviesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	query := "select account.id, account.name, movie.id, movie.name, movie_review.score from follow , movie_review, movie, account  where follower_id = $1 and account_id = following_id and movie.id = movie_review.movie_id and follow.following_id = account.id"
	rows, err := db.Query(context.Background(), query, uid)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserMovieHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("UserMovies query succesfull")

	type t struct {
		FollowingUid  string `json:"followingUid"`
		FollowingName string `json:"followingName"`
		Name          string `json:"name"`
		Id            int    `json:"id"`
		Score         int    `json:"score"`
	}

	var resp = make([]t, 0)

	for rows.Next() {
		var fUid, fName, name string
		var id, score int
		err = rows.Scan(&fUid, &fName, &id, &name, &score)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		m := t{fUid, fName, name, id, score}
		resp = append(resp, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func userActivityGamesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	query := "select account.id, account.name, game.id, game.name, game_review.score from follow , game_review, game, account  where follower_id = $1 and account_id = following_id and game.id = game_review.game_id and follow.following_id = account.id"
	rows, err := db.Query(context.Background(), query, uid)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserMovieHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("UserMovies query succesfull")

	type t struct {
		FollowingUid  string `json:"followingUid"`
		FollowingName string `json:"followingName"`
		Name          string `json:"name"`
		Id            int    `json:"id"`
		Score         int    `json:"score"`
	}

	var resp = make([]t, 0)

	for rows.Next() {
		var fUid, fName, name string
		var id, score int
		err = rows.Scan(&fUid, &fName, &id, &name, &score)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		m := t{fUid, fName, name, id, score}
		resp = append(resp, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func userActivityShowsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	query := "select account.id, account.name, show.id, show.name, show_review.score from follow , show_review, show, account  where follower_id = $1 and account_id = following_id and show.id = show_review.show_id and follow.following_id = account.id"
	rows, err := db.Query(context.Background(), query, uid)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Query in UserMovieHandler failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	log.Print("UserMovies query succesfull")

	type t struct {
		FollowingUid  string `json:"followingUid"`
		FollowingName string `json:"followingName"`
		Name          string `json:"name"`
		Id            int    `json:"id"`
		Score         int    `json:"score"`
	}

	var resp = make([]t, 0)

	for rows.Next() {
		var fUid, fName, name string
		var id, score int
		err = rows.Scan(&fUid, &fName, &id, &name, &score)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		m := t{fUid, fName, name, id, score}
		resp = append(resp, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func addAdminHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached addAdminHandler")
	decoder := json.NewDecoder(r.Body)

	d := make(map[string]string)

	err := decoder.Decode(&d)
	if err != nil {
		log.Printf("Error decoding follow review json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}
	uid, ok := d["uid"]
	if !ok {
		log.Println("JSON doesn't have uid")
		http.Error(w, "Missing uid", http.StatusBadRequest)
		return
	}

	query := "UPDATE account SET is_admin = true WHERE id = $1"
	_, err = db.Exec(context.Background(), query, uid)
	if err != nil {
		log.Printf("Error updating value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error making admin", http.StatusBadRequest)
		} else {
			http.Error(w, "Error making admin", http.StatusInternalServerError)

		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
func changeUserNameHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached changeUserNameHandler")
	vars := mux.Vars(r)
	uid := vars["uid"]

	decoder := json.NewDecoder(r.Body)

	d := make(map[string]string)

	err := decoder.Decode(&d)
	if err != nil {
		log.Printf("Error decoding follow review json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}

	name, ok := d["name"]
	if !ok {
		log.Println("JSON doesn't have name")
		http.Error(w, "Missing name", http.StatusBadRequest)
		return
	}

	u := &user{Id: uid, Name: name}
	log.Println(u)

	query := "UPDATE account SET name = $1 WHERE id = $2 RETURNING register_date, is_admin"
	err = db.QueryRow(context.Background(), query, name, uid).Scan(&u.RegisterDate, &u.Admin)
	if err != nil {
		log.Printf("Error updating value: %T %v\n", err, err)
		if _, ok := err.(*pgconn.PgError); ok {
			http.Error(w, "Error making admin", http.StatusBadRequest)
		} else {
			http.Error(w, "Error making admin", http.StatusInternalServerError)

		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached updateMovieHandler")
	vars := mux.Vars(r)
	id := vars["id"]
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

	t.Id, _ = strconv.Atoi(id)
	if t.Id == 0 || t.Name == "" || t.Description == "" || t.ImageUrl == "" || t.ReleaseDate == "" {
		log.Printf("Error incomplete json: %v\n", err)
		http.Error(w, "Unable incomplete json", http.StatusBadRequest)
		return
	}

	d, err := time.Parse("2006-1-2", t.ReleaseDate)
	if err != nil {
		log.Printf("Error decoding date from json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}

	m := movie{Name: strings.ToLower(t.Name), Description: t.Description, ImageUrl: t.ImageUrl, ReleaseDate: d, Id: t.Id}
	query := "UPDATE movie SET name = $1, release_date = $2, description = $3, image_url = $4 WHERE id = $5"
	c, err := db.Exec(context.Background(), query, m.Name, m.ReleaseDate, m.Description, m.ImageUrl, m.Id)
	if err != nil {
		log.Printf("Error updating value: %T %v\n", err, err)
		if e, ok := err.(*pgconn.PgError); ok {
			if e.Code == "23505" {
				http.Error(w, "Repeting name", http.StatusBadRequest)
			} else {
				http.Error(w, "Error updating movie", http.StatusBadRequest)
			}
		} else {
			http.Error(w, "Error updating movie", http.StatusInternalServerError)

		}
		return
	} else if c.RowsAffected() == 0 {
		http.Error(w, "movie not found", http.StatusNotFound)
		log.Printf("Error updating: movie with id %v \n", m.Id)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

func updateGameHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached updateGameHandler")
	vars := mux.Vars(r)
	id := vars["id"]
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

	t.Id, _ = strconv.Atoi(id)
	if t.Id == 0 || t.Name == "" || t.Description == "" || t.ImageUrl == "" || t.ReleaseDate == "" {
		log.Printf("Error incomplete json: %v\n", err)
		http.Error(w, "Unable incomplete json", http.StatusBadRequest)
		return
	}

	d, err := time.Parse("2006-1-2", t.ReleaseDate)
	if err != nil {
		log.Printf("Error decoding date from json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}

	g := game{Name: strings.ToLower(t.Name), Description: t.Description, ImageUrl: t.ImageUrl, ReleaseDate: d, Id: t.Id}
	query := "UPDATE game SET name = $1, release_date = $2, description = $3, image_url = $4 WHERE id = $5"
	c, err := db.Exec(context.Background(), query, g.Name, g.ReleaseDate, g.Description, g.ImageUrl, g.Id)
	if err != nil {
		log.Printf("Error updating value: %T %v\n", err, err)
		if e, ok := err.(*pgconn.PgError); ok {
			if e.Code == "23505" {
				http.Error(w, "Repeting name", http.StatusBadRequest)
			} else {
				http.Error(w, "Error updating game", http.StatusBadRequest)
			}
		} else {
			http.Error(w, "Error updating game", http.StatusInternalServerError)

		}
		return
	} else if c.RowsAffected() == 0 {
		http.Error(w, "game not found", http.StatusNotFound)
		log.Printf("Error updating: game with id %v \n", g.Id)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(g)
}

func updateShowHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Reached updateShowHandler")
	vars := mux.Vars(r)
	id := vars["id"]
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
		log.Printf("Error decoding Show json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}

	t.Id, _ = strconv.Atoi(id)
	if t.Id == 0 || t.Name == "" || t.Description == "" || t.ImageUrl == "" || t.ReleaseDate == "" {
		log.Printf("Error incomplete json: %v\n", err)
		http.Error(w, "Unable incomplete json", http.StatusBadRequest)
		return
	}

	d, err := time.Parse("2006-1-2", t.ReleaseDate)
	if err != nil {
		log.Printf("Error decoding date from json: %v\n", err)
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}

	s := show{Name: strings.ToLower(t.Name), Description: t.Description, ImageUrl: t.ImageUrl, ReleaseDate: d, Id: t.Id}
	query := "UPDATE show SET name = $1, release_date = $2, description = $3, image_url = $4 WHERE id = $5"
	c, err := db.Exec(context.Background(), query, s.Name, s.ReleaseDate, s.Description, s.ImageUrl, s.Id)
	if err != nil {
		log.Printf("Error updating value: %T %v\n", err, err)
		if e, ok := err.(*pgconn.PgError); ok {
			if e.Code == "23505" {
				http.Error(w, "Repeting name", http.StatusBadRequest)
			} else {
				http.Error(w, "Error updating show", http.StatusBadRequest)
			}
		} else {
			http.Error(w, "Error updating game", http.StatusInternalServerError)

		}
		return
	} else if c.RowsAffected() == 0 {
		http.Error(w, "show not found", http.StatusNotFound)
		log.Printf("Error updating: show with id %v \n", s.Id)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s)
}
