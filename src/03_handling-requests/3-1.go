package main

import (
	"net/http"
)

func main() {
	// ポート番号を指定しないと80が使われる
	// ハンドラがnilの場合はDefaultServeMuxが使われる
	http.ListenAndServe("", nil)
}
