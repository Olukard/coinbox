package handlers

import (
	"fmt"
	"net/http"
)

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("Host: %s\nPath: %s\nMethod: %s",
		r.Host, r.URL.Path, r.Method)
	w.Write([]byte(s))
}
