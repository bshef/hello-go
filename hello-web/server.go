package main

import (
	"io"
	"net/http"
)

func setServerHeader(w http.ResponseWriter) {
    w.Header().Set("Server", "Hello Go Web Server");
}

func hello(w http.ResponseWriter, req *http.Request) {
    setServerHeader(w)
	io.WriteString(w, "Hello, world!")
}

func health(w http.ResponseWriter, req *http.Request) {
    setServerHeader(w)
    w.WriteHeader(200);
}

func main() {
	http.HandleFunc("/", hello)
    http.HandleFunc("/health", health)
	http.ListenAndServe(":8000", nil)
}
