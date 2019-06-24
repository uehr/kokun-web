package main

import (
	"log"
	"net/http"
	"os"

	"github.com/uehr/kokun-web/handler"
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
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media/"))))

	http.ListenAndServe(":"+port, nil)
}
