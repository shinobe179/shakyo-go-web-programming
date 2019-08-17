package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Post struct {
	User string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "そのようなサービスはありません。他をあたってください。")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

// output
// $ curl -v http://localhost:8080/json
// *   Trying 127.0.0.1...
// * Connected to localhost (127.0.0.1) port 8080 (#0)
// > GET /json HTTP/1.1
// > Host: localhost:8080
// > User-Agent: curl/7.47.0
// > Accept: */*
// >
// < HTTP/1.1 200 OK
// < Content-Type: application/json
// < Date: Sat, 17 Aug 2019 14:11:44 GMT
// < Content-Length: 45
// <
// * Connection #0 to host localhost left intact
// {"User":"shinobe179","Threads":["1","2","3"]}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "shinobe179",
		Threads: []string{"1", "2", "3"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
