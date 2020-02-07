package main

import (
	"net/http"
)

const (
	serverAddress string = ":8080"
)

func main() {
	mux := http.NewServeMux()

	srv := http.Server{
		Addr: serverAddress,

		Handler: mux,
	}
}
