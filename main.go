package main

import (
	"github.com/dannyjhall/gotyper/socket"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/connect", socket.Handler)
	http.Handle("/", http.FileServer(http.Dir("./www")))
	err := http.ListenAndServe(":"+getPort(), nil)

	if err != nil {
		panic("Error: " + err.Error())
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT not set, using 8080")
		port = "8080"
	}
	return port
}
