package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

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
	mux := http.NewServeMux()

	handler := http.FileServer(http.Dir("."))
	mux.Handle("/go-cord/", middleware(handler))
	mux.HandleFunc("GET /api/stuff", myHandler)

	port := os.Getenv("PORT")

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Printf("Running Server on port %s...", port)

	err := server.ListenAndServe()
	log.Fatal(err)
}
