package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// リクエストを解析してからFormフィールドにアクセスする
	r.ParseForm()
	// クエリストリングとPOSTフォームの両方の値を参照したい時
	//fmt.Fprintln(w, r.Form)
	// POSTフォームの値だけを参照したい時
	fmt.Fprintln(w, r.PostForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
