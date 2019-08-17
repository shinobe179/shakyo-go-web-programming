package main

import (
	"fmt"
	"net/http"
)

// output
// 
// $ curl -v localhost:8080/write
// *   Trying 127.0.0.1...
// * Connected to localhost (127.0.0.1) port 8080 (#0)
// > GET /write HTTP/1.1
// > Host: localhost:8080
// > User-Agent: curl/7.47.0
// > Accept: */*
// >
// < HTTP/1.1 200 OK
// < Date: Sat, 17 Aug 2019 13:23:31 GMT
// < Content-Length: 95
// < Content-Type: text/html; charset=utf-8
// <
// <html>
// <head><title>Go Web Programming</title></head>
// <body><h1>Hello World</h1></body>
// * Connection #0 to host localhost left intact

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

// output
// 
// $ curl -v localhost:8080/writeheader
// *   Trying 127.0.0.1...
// * Connected to localhost (127.0.0.1) port 8080 (#0)
// > GET /writeheader HTTP/1.1
// > Host: localhost:8080
// > User-Agent: curl/7.47.0
// > Accept: */*
// >
// < HTTP/1.1 501 Not Implemented
// < Date: Sat, 17 Aug 2019 13:22:47 GMT
// < Content-Length: 82
// < Content-Type: text/plain; charset=utf-8
// <
// そのようなサービスはありません。他をあたってください。
// * Connection #0 to host localhost left intact

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "そのようなサービスはありません。他をあたってください。")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	server.ListenAndServe()
}
