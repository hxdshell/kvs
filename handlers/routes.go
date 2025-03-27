package handlers

import (
	"net/http"
)

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/get/", handleGet)
	mux.HandleFunc("/set/", handleSet)
	mux.HandleFunc("/delete/", handleDelete)

	mux.HandleFunc("/list", handleList)
	mux.HandleFunc("/list/", handleList)

	mux.HandleFunc("/flush", handleFlush)
	mux.HandleFunc("/flush/", handleFlush)

	mux.HandleFunc("/inc/", handleInc)
	mux.HandleFunc("/dec/", handleDec)

	return mux
}
