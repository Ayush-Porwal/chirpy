package main

import (
	"fmt"
	"net/http"
)

func main() {

	serveMux := http.NewServeMux()

	serveMux.Handle("/", http.FileServer(http.Dir(".")))

	server := &http.Server{
		Handler: serveMux,
		Addr:    ":8080",
	}

	fmt.Println("Starting server now")

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Something went wrong")
	}
}
