package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.HandleFunc("/hello", Hello)

	log.Println("Server started...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("ERROR: Failed to start Server (err: %v)", err)
	}
}

// Hello returns text message
func Hello(w http.ResponseWriter, _ *http.Request) {
	text := "<html><head></head><body>Hi from <b>Go</b> on Heroku!!!</body><html>"

	count, err := w.Write([]byte(text))
	if err != nil {
		log.Printf("ERROR: %v", err)
		return
	}
	log.Printf("SUCCESS: %d byte sent", count)
}
