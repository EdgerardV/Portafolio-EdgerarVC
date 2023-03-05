package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	r := http.NewServeMux()

	r.HandleFunc("/route1", index)
	r.HandleFunc("/route2", index)
	buildHandler := http.FileServer(http.Dir("../_Frontend/build"))
	r.Handle("/", buildHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8083",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server started on PORT 8083")
	log.Fatal(srv.ListenAndServe())
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../_Frontend/build/index.html")
}
