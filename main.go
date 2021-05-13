package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

const wait = 5 * time.Second

var db *pgxpool.Pool

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("$DATABASE_URL must be set")
	}

	var err error
	db, err = pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)

	r.HandleFunc("/movies", moviesHandler).Methods("GET")
	r.HandleFunc("/movies", addMovieHandler).Methods("POST")
	r.HandleFunc("/movies/{id}", singleMovieHandler).Methods("GET")

	r.HandleFunc("/games", gamesHandler).Methods("GET")
	r.HandleFunc("/games", addGameHandler).Methods("POST")
	r.HandleFunc("/games/{id}", singleGameHandler).Methods("GET")

	r.HandleFunc("/shows", showsHandler).Methods("GET")
	r.HandleFunc("/shows", addShowHandler).Methods("POST")
	r.HandleFunc("/shows/{id}", singleShowHandler).Methods("GET")

	r.HandleFunc("/users", userHandler).Methods("GET")
	r.HandleFunc("/users", addUserHandler).Methods("POST")
	r.HandleFunc("/users/{uid}", singleUserHandler).Methods("GET")

	r.HandleFunc("/users/{uid}/movies", UserMoviesHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/movies/{id}", singleUserMovieHandler).Methods("GET")

	r.HandleFunc("/users/{uid}/games", UserGamesHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/games/{id}", singleUserGameHandler).Methods("GET")

	r.HandleFunc("/users/{uid}/shows", UserShowsHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/shows/{id}", singleUserShowHandler).Methods("GET")
	
	r.HandleFunc("/users/{uid}/follows", UserFollowsHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/follows/{id}", singleUserFollowHandler).Methods("GET")

	// Create a server so you can gracefully shutdown it
	srv := &http.Server{
		Addr: ":" + port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("Listening in port %s", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM will not be caught.
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	log.Println("shutting down")
}
