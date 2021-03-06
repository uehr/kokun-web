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

	http.HandleFunc("/senryu", handler.SenryuApi)
	// http.HandleFunc("/upload-twitter", handler.UploadTwitter)

	http.Handle("/", http.FileServer(http.Dir("views")))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts/"))))
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media/"))))

	http.ListenAndServe(":"+port, nil)
}
