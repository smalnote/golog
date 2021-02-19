package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed static
var static embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(static)))

	log.Println("listening on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
