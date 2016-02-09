package main

import (
	"io"
	"net/http"
	"strconv"
)

// Variables
const port int = 8000
var portString string = ":" + strconv.Itoa(port)

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
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
    mux.HandleFunc("/health", health)
	http.ListenAndServe(portString, nil)
}
