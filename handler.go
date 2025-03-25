package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/get/{key}", handleGet)
	mux.HandleFunc("/put/{key}", handlePut)

	return mux
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed\n", http.StatusMethodNotAllowed)
		return
	}

	params := strings.Split(r.URL.Path, "/")

	key := params[2]

	val := get(key)
	w.Write(val)

	if val != nil {
		log.Printf("GET(%s,%s)", key, val)
	}
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Only PUT method is allowed\n", http.StatusMethodNotAllowed)
		return
	}
	params := strings.Split(r.URL.Path, "/")
	key := params[2]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body\n", http.StatusInternalServerError)
		return
	}
	put(key, body)
	log.Printf("PUT(%s,%s)", key, body)
}
