package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// リクエストを解析してからFormフィールドにアクセスする
	//r.ParseForm()
	// クエリストリングとPOSTフォームの両方の値を参照したい時
	//fmt.Fprintln(w, r.Form)
	// POSTフォームの値だけを参照したい時
	//fmt.Fprintln(w, r.PostForm)
	// multipart/formdataの時
	//r.ParseMultipartForm(1024)
	//fmt.Fprintln(w, r.MultipartForm)
	// FormValueメソッドを使う時
	fmt.Fprintln(w, r.FormValue("hello"))
	// output: map[thread:[123] hello:[shinobe179 world] post:[456]]
	fmt.Fprintln(w, r.Form)

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
