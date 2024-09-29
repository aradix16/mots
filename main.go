package main

import (
	"log"
	"net/http"

	"github.com/aradix16/mots/config"
	"github.com/aradix16/mots/web"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	r := http.NewServeMux()

	fs := http.FS(web.StaticFS)
	r.Handle("GET /static/", http.FileServer(fs))
	r.Handle("GET app.localhost/static/", http.FileServer(fs))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to mots.wiki!"))
	})

	addr := ":" + cfg.Port
	log.Printf("Starting the server on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
