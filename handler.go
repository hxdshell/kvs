package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PONG"))
	})

	mux.HandleFunc("/get/{key}", handleGet)
	mux.HandleFunc("/set/{key}", handleSet)
	mux.HandleFunc("/list", handleList)
	mux.HandleFunc("/delete/{key}", handleDelete)
	return mux
}

func handleList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed\n", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list())
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

func handleSet(w http.ResponseWriter, r *http.Request) {
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
	set(key, body)
	log.Printf("SET(%s,%s)", key, body)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Only DELETE method is allowed\n", http.StatusMethodNotAllowed)
		return
	}

	params := strings.Split(r.URL.Path, "/")
	key := params[2]
	remove(key)
}
