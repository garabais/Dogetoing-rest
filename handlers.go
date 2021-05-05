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

	var movies []*movie

	for i:= 0; rows.Next(); i++ {
		m := &movie{}
		err = rows.Scan(&m.Id, &m.Name, &m.Description, &m.ImageUrl, &m.ReleaseDate)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		movies = append(movies, m)
		log.Print(i)
	}

	json.NewEncoder(w).Encode(movies)
}

// func allArticlesHandler(w http.ResponseWriter, r *http.Request) {
//     log.Println("Endpoint Hit: returnAllArticles")
//     json.NewEncoder(w).Encode(articles)
// }

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
