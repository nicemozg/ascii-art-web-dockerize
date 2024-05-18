package main

import (
	"ascii-art-web/services"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ascii-art", services.AsciiArtHandler)
	mux.HandleFunc("/", services.AsciiArtPageHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./templates/static"))))
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}
