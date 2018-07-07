package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my new Chat Server Under Construction.")
}


func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

		mux.HandleFunc("/", index)

	server := &http.Server {
		Addr:		"0.0.0.0:8080",
		Handler:	mux,
	}
	server.ListenAndServe()

}