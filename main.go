package main

import (
	"gocord/internal/database"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct{
	database *database.Queries
}

func myHandler(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "text/plain; charset=utf-8")
	write.WriteHeader(http.StatusOK)
	write.Write([]byte("Welcome to Go-Cord"))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		http.StripPrefix("/go-cord", next).ServeHTTP(write, request)
	})
}

func main() {
	godotenv.Load()

	mux := http.NewServeMux()

	dbURL := os.Getenv("DB_URL")

	fmt.Printf( "url %s: \n", dbURL )

	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)

	cfg := Config{
		database : dbQueries,
	}

	//dbQueries.CreateChatlog( context.Background(), "Hello World!" )

	handler := http.FileServer(http.Dir("."))
	mux.Handle("/go-cord/", middleware(handler))
	mux.HandleFunc("GET /api/stuff", myHandler)

	port := os.Getenv("PORT")

	//mux.HandleFunc("POST /api/message", handleMessage)
	mux.HandleFunc("POST /api/message", cfg.handleMessage)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Printf("Running Server on port %s...\n", port)

	err = server.ListenAndServe()
	log.Fatal(err)
}
