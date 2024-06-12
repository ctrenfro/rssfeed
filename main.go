package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/ctrenfro/rssfeed/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// by default, godotenv will look for a file named .env in the current directory
	godotenv.Load(".env")

	dbURL := os.Getenv("dbURL")
	if dbURL == "" {
		log.Fatal("dbURL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	apicfg := apiConfig{
		DB: dbQueries,
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", handlerHealthzGet)
	mux.HandleFunc("GET /v1/err", errGet)
	mux.HandleFunc("POST /v1/users", apicfg.createUser)
	mux.HandleFunc("GET /v1/users", apicfg.middlewareAuth(apicfg.handlerGetUser))
	mux.HandleFunc("POST /v1/feeds", apicfg.middlewareAuth(apicfg.handlerCreateFeed))
	mux.HandleFunc("GET /v1/feeds", apicfg.handlerGetFeeds)
	mux.HandleFunc("POST /v1/feed_follows", apicfg.middlewareAuth(apicfg.handlerCreateFeedFollows))
	mux.HandleFunc("GET /v1/feed_follows", apicfg.middlewareAuth(apicfg.handlerGetFeedFollows))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowsID}", apicfg.middlewareAuth(apicfg.handlerDeleteFeedFollows))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
