package main

import (
	"net/http"
)

func main() {

	// ServeMux := http.NewServeMux()

	Server := &http.Server{
		Handler: nil,
		Addr:    ":8080",
	}

	Server.ListenAndServe()
}
