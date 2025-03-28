package utils

import (
	"errors"
	"net/http"
	"strings"
)

// validates request method and splits the url path into parameters. The first member of the []string is empty.
func ParseUrl(allowedMethod string, w http.ResponseWriter, r *http.Request) ([]string, error) {
	if r.Method != allowedMethod {
		msg := "ERROR: only " + allowedMethod + " method is allowed\n"
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return make([]string, 0), errors.New(msg)
	}

	trimmed := strings.TrimRight(r.URL.Path, "/") // remove trailing "/"
	params := strings.Split(trimmed, "/")

	return params, nil
}
