package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// http://127.0.0.1:8000/redirect?url=http://www.google.com
	http.HandleFunc("/redirect", RedirectHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("[*] Starting server on port " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	url := query.Get("url")
	http.Redirect(w, r, url, 301)
}
