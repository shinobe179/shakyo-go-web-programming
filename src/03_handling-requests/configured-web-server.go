package main

import (
	"net/http"
)

func main() {
	// Server構造体で細かな設定変更ができる
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}
