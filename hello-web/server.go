//	Package server provides logic to listen and serve content and faciliate
//	a RESTful API.
package server

import (
	"io"
	"net/http"
	"strconv"
)

//
//	CONSTANT Variables
//

//	port defines the port on which the server will listen and serve.
const port int = 8000

//	serverHeader defines the value of the "Server" header set on HTTP responses.
const serverHeader string = "Hello Go Web Server"

//
//	API Handler Functions
//

//	hello responds with "Hello, world!"
func hello(w http.ResponseWriter, req *http.Request) {
	setServerHeader(w)
	io.WriteString(w, "Hello, world!")
}

//	health responds with 200 OK
func health(w http.ResponseWriter, req *http.Request) {
	setServerHeader(w)
	w.WriteHeader(200)
}

//
//	Server functions
//

//	setServerHeader sets the response header.
func setServerHeader(w http.ResponseWriter) {
	w.Header().Set("Server", serverHeader)
}

//	mapAPIFunctions maps API handler functions to the API path, as specified in a mapping.
func mapAPIFunctions(router *http.ServeMux, apiMap map[string]func(http.ResponseWriter, *http.Request)) {
	for path, apiFunction := range apiMap {
		router.HandleFunc(path, apiFunction)
	}
}

//	startServer defines logic that occurs when server starts.
func startServer(port int, apiMap map[string]func(http.ResponseWriter, *http.Request)) {
	mapAPIFunctions(http.NewServeMux(), apiMap)
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
