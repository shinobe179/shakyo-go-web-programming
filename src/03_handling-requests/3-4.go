package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}
	// ListenAndServeTLSを使う
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
