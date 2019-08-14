package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
	// 特定ヘッダー(Accept-Encoding)だけ出力する
	// fmt.Fprintln(w, h["Accept-Encoding"])
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", headers)
	server.ListenAndServe()
}
