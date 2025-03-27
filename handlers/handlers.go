package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"kvs/core"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// TODO : Handler logic is pretty redundant, make a better wrapper
func handleFlush(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "ERROR : Only DELETE method is allowed\n", http.StatusMethodNotAllowed)
		return
	}

	trimmed := strings.TrimRight(r.URL.Path, "/") // remove trailing "/"
	params := strings.Split(trimmed, "/")

	if len(params) > 2 {
		http.Error(w, "ERROR : usage: /flushdb \n", http.StatusNotFound)
		return
	}

	core.Flushdb()
	log.Println("FLUSHDB")
}

func handleList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "ERROR : Only GET method is allowed\n", http.StatusMethodNotAllowed)
		return
	}

	trimmed := strings.TrimRight(r.URL.Path, "/") // remove trailing "/"
	params := strings.Split(trimmed, "/")

	if len(params) > 2 {
		http.Error(w, "ERROR : usage: /list \n", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	kvList := core.List()
	json.NewEncoder(w).Encode(kvList)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "ERROR : Only GET method is allowed\n", http.StatusMethodNotAllowed)
		return
	}

	trimmed := strings.TrimRight(r.URL.Path, "/") // remove trailing "/" ex: <host>/get/key/
	params := strings.Split(trimmed, "/")

	if len(params) != 3 {
		http.Error(w, "ERROR : usage: /get/{key}\n", http.StatusNotFound)
		return
	}

	key := params[2]
	val := core.Get(key)

	if len(val) == 0 {
		http.Error(w, "ERROR: key not found\n", http.StatusNotFound)
		return
	}

	w.Write(val)

	if val != nil {
		log.Printf("GET(%s,%s)", key, val)
	}
}

func handleSet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "ERROR : Only PUT method is allowed\n", http.StatusMethodNotAllowed)
		return
	}

	trimmed := strings.TrimRight(r.URL.Path, "/") // remove trailing "/"  ex: <host>/set/key/
	params := strings.Split(trimmed, "/")

	if len(params) != 3 {
		http.Error(w, "ERROR : usage: /get/{key}", http.StatusNotFound)
		return
	}

	key := params[2]
	if key == "" {
		fmt.Println("boring")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "ERROR : Failed to read request body\n", http.StatusInternalServerError)
		return
	}
	if len(body) == 0 {
		http.Error(w, "ERROR : NULL VALUE\n", http.StatusBadRequest)
		return
	}
	core.Set(key, body)
	log.Printf("SET(%s,%s)", key, body)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "ERROR : Only DELETE method is allowed\n", http.StatusMethodNotAllowed)
		return
	}

	trimmed := strings.TrimRight(r.URL.Path, "/")
	params := strings.Split(trimmed, "/")

	if len(params) != 3 {
		http.Error(w, "ERROR : usage: /delete/{key}", http.StatusNotFound)
		return
	}

	key := params[2]
	core.Remove(key)
}

func handleInc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "ERROR : Only GET method is allowed\n", http.StatusMethodNotAllowed)
		return
	}

	trimmed := strings.TrimRight(r.URL.Path, "/") // remove trailing "/" ex: <host>/inc/counter/
	params := strings.Split(trimmed, "/")

	if len(params) != 3 && len(params) != 4 {
		http.Error(w, "ERROR : usage: /inc/{key}/{magnitude}\n", http.StatusNotFound)
		return
	}

	key := params[2]
	magnitude := 1

	if len(params) == 4 {

		result, err := strconv.Atoi(params[3])
		if err != nil {
			http.Error(w, "ERROR: magnitude should be an integer\n", http.StatusBadRequest)
			return
		}

		magnitude = result
	}

	err := core.IncDec(key, magnitude, true)
	if err != nil {
		http.Error(w, "ERROR : Value is either non-existent or not an integer\n", http.StatusBadRequest)
		return
	}
}

func handleDec(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "ERROR : Only GET method is allowed\n", http.StatusMethodNotAllowed)
		return
	}

	trimmed := strings.TrimRight(r.URL.Path, "/") // remove trailing "/" ex: <host>/dec/counter/
	params := strings.Split(trimmed, "/")

	if len(params) != 3 && len(params) != 4 {
		http.Error(w, "ERROR : usage: /dec/{key}/{magnitude}\n", http.StatusNotFound)
		return
	}

	key := params[2]
	magnitude := 1

	if len(params) == 4 {
		result, err := strconv.Atoi(params[3])
		if err != nil {
			http.Error(w, "ERROR: magnitude should be an integer\n", http.StatusBadRequest)
			return
		}

		magnitude = result
	}

	err := core.IncDec(key, magnitude, false)
	if err != nil {
		http.Error(w, "ERROR : Value is either non-existent or not an integer\n", http.StatusBadRequest)
		return
	}
}
