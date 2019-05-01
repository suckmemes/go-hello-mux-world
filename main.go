package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello gorilla/mux world!\n")
}

func getSampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET /sample\n")
}

func postSampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST /sample\n")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/sample", postSampleHandler).Methods("POST")
	r.HandleFunc("/sample", getSampleHandler).Methods("GET")

	// Server configuration
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Run
	log.Fatal(srv.ListenAndServe())
}
