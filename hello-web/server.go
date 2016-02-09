package main

import (
	"io"
	"net/http"
	"strconv"
)

// CONSTANT Variables
const port int = 8000
const serverHeader string = "Hello Go Web Server"

//
//	API Handler Functions
//

//	Responds with "Hello, world!"
func hello(w http.ResponseWriter, req *http.Request) {
	setServerHeader(w)
	io.WriteString(w, "Hello, world!")
}

//	Responds with 200 OK
func health(w http.ResponseWriter, req *http.Request) {
	setServerHeader(w)
	w.WriteHeader(200)
}

//
//	Server functions
//

//	Sets the response header
func setServerHeader(w http.ResponseWriter) {
	w.Header().Set("Server", serverHeader)
}

//	Maps API handler functions to the API path, as specified in a mapping
func mapAPIfunctions(router *http.ServeMux, apiMap map[string]func(http.ResponseWriter, *http.Request)) {
	for path, apiFunction := range apiMap {
		router.HandleFunc(path, apiFunction)
	}
}

//	Logic that occurs when server starts
func startServer(port int, apiMap map[string]func(http.ResponseWriter, *http.Request)) {
	mapAPIfunctions(http.NewServeMux(), apiMap)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

//
//	Server main entry point
//
func main() {
	startServer(port, map[string]func(http.ResponseWriter, *http.Request){
		"/":       hello,
		"/health": health,
	})
}
