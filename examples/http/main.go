package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("content-type", "application/json")
		_, _ = fmt.Fprintf(w, `{"ok": true}`)
	})

	_ = http.ListenAndServe(":8080", mux)
}
