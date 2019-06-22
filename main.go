package main

import (
	"log"
	"net/http"
	"os"

	"github.com/uehr/kokun/pkg/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/senryu", handler.SenryuApi)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts/"))))
	http.ListenAndServe(":"+port, nil)
}
