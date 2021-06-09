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

	r.HandleFunc("/movies", nameQueryMoviesHandler).Queries("name", "{name}").Methods("GET")
	r.HandleFunc("/movies", moviesHandler).Methods("GET")
	r.HandleFunc("/movies", addMovieHandler).Methods("POST")
	r.HandleFunc("/movies/{id:[0-9]+}", singleMovieHandler).Methods("GET")
	r.HandleFunc("/movies/{id:[0-9]+}", deleteMovieHandler).Methods("DELETE")
	r.HandleFunc("/movies/{id:[0-9]+}", updateMovieHandler).Methods("PUT")

	r.HandleFunc("/games", nameQueryGamesHandler).Queries("name", "{name}").Methods("GET")
	r.HandleFunc("/games", gamesHandler).Methods("GET")
	r.HandleFunc("/games", addGameHandler).Methods("POST")
	r.HandleFunc("/games/{id:[0-9]+}", singleGameHandler).Methods("GET")
	r.HandleFunc("/games/{id:[0-9]+}", deleteGameHandler).Methods("DELETE")
	r.HandleFunc("/games/{id:[0-9]+}", updateGameHandler).Methods("PUT")

	r.HandleFunc("/shows", nameQueryShowsHandler).Queries("name", "{name}").Methods("GET")
	r.HandleFunc("/shows", showsHandler).Methods("GET")
	r.HandleFunc("/shows", addShowHandler).Methods("POST")
	r.HandleFunc("/shows/{id:[0-9]+}", singleShowHandler).Methods("GET")
	r.HandleFunc("/shows/{id:[0-9]+}", deleteShowHandler).Methods("DELETE")
	r.HandleFunc("/shows/{id:[0-9]+}", updateShowHandler).Methods("PUT")

	r.HandleFunc("/users", nameFollowQueryUserHandler).Queries("name", "{name}", "nf", "{uid}").Methods("GET")
	r.HandleFunc("/users", nameAdminQueryUserHandler).Queries("name", "{name}", "admin", "{admin}").Methods("GET")
	r.HandleFunc("/users", nameQueryUserHandler).Queries("name", "{name}").Methods("GET")
	r.HandleFunc("/users", userHandler).Methods("GET")
	r.HandleFunc("/users", addUserHandler).Methods("POST")
	r.HandleFunc("/users/{uid}", singleUserHandler).Methods("GET")
	r.HandleFunc("/users/{uid}", changeUserNameHandler).Methods("PUT")
		
	r.HandleFunc("/users/{uid}/feed/movies", userActivityMoviesHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/feed/games", userActivityGamesHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/feed/shows", userActivityShowsHandler).Methods("GET")

	r.HandleFunc("/users/{uid}/movies", nameQueryUserMoviesHandler).Queries("name", "{name}").Methods("GET")
	r.HandleFunc("/users/{uid}/movies", userMoviesHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/movies", addUserMovieHandler).Methods("POST")
	r.HandleFunc("/users/{uid}/movies/{id:[0-9]+}", singleUserMovieHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/movies/{id:[0-9]+}", deleteUserMovieHandler).Methods("DELETE")
	r.HandleFunc("/users/{uid}/movies/{id:[0-9]+}", updateUserMovieHandler).Methods("PUT")

	r.HandleFunc("/users/{uid}/games", nameQueryUserGamesHandler).Queries("name", "{name}").Methods("GET")
	r.HandleFunc("/users/{uid}/games", userGamesHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/games", addUserGameHandler).Methods("POST")
	r.HandleFunc("/users/{uid}/games/{id:[0-9]+}", singleUserGameHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/games/{id:[0-9]+}", deleteUserGamesHandler).Methods("DELETE")
	r.HandleFunc("/users/{uid}/games/{id:[0-9]+}", updateUserGameHandler).Methods("PUT")

	r.HandleFunc("/users/{uid}/shows", nameQueryUserShowsHandler).Queries("name", "{name}").Methods("GET")
	r.HandleFunc("/users/{uid}/shows", userShowsHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/shows", addUserShowHandler).Methods("POST")
	r.HandleFunc("/users/{uid}/shows/{id:[0-9]+}", singleUserShowHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/shows/{id:[0-9]+}", deleteUserShowsHandler).Methods("DELETE")
	r.HandleFunc("/users/{uid}/shows/{id:[0-9]+}", updateUserShowsHandler).Methods("PUT")

	r.HandleFunc("/users/{uid}/follows", nameUserFollowsHandler).Methods("GET").Queries("name", "{name}")
	r.HandleFunc("/users/{uid}/follows", userFollowsHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/follows", addUserFollowHandler).Methods("POST")
	r.HandleFunc("/users/{uid}/follows/{id}", singleUserFollowHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/follows/{id}", deleteUserFollowHandler).Methods("DELETE")

	r.HandleFunc("/users/{uid}/followers", UserFollowersHandler).Methods("GET")
	r.HandleFunc("/users/{uid}/follows/{id}", singleUserFollowersHandler).Methods("GET")

	r.HandleFunc("/admin", addAdminHandler).Methods("POST")

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
