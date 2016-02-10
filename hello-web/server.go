// Package server provides logic to listen and serve content and faciliate
// a RESTful API.
package server

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

//
//	CONSTANT Variables
//

//	port defines the port on which the server will listen and serve.
const port int = 8000

//	serverHeader defines the value of the "Server" header set on HTTP responses.
const serverHeader string = "Hello Go Web Server"

//	Logger variables
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

//
//	Type definitions
//

//	Type apiError contains important details about an API error.
type apiError struct {
	Error   error
	Message string
	Code    int
}

//	Type apiHandler handles an HTTP request and can return an error if necessary.
type apiHandler func(http.ResponseWriter, *http.Request) *apiError

//
//	API Handler Functions
//

//	hello responds with "Hello, world!"
func hello(w http.ResponseWriter, req *http.Request) *apiError {
	setServerHeader(w)
	io.WriteString(w, "Hello, world!")
	return nil
}

//	health responds with 200 OK
func health(w http.ResponseWriter, req *http.Request) *apiError {
	setServerHeader(w)
	w.WriteHeader(200)
	return nil
}

//
//	Server functions
//

//	initializeLogger sets up the logging capabilities of the application.
func initializeLogger(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {
	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	//	Usage:
	// Trace.Println("I have something standard to say")
	// Info.Println("Special Information")
	// Warning.Println("There is something you need to know about")
	// Error.Println("Something has failed")
}

//	ServeHTTP method calls the apiHandler function and displays the returned
//	error (if any).
func (fn apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil {
		Error.Println(e.Error)
		http.Error(w, e.Message, e.Code)
	} else {
		Info.Println("Request received ... ")
	}
}

//	setServerHeader sets the response header.
func setServerHeader(w http.ResponseWriter) {
	w.Header().Set("Server", serverHeader)
}

//	mapAPIFunctions maps API handler functions to the API path, as specified in a mapping.
func mapAPIFunctions(router *http.ServeMux, apiMap map[string]func(http.ResponseWriter, *http.Request) *apiError) {
	for path, apiFunction := range apiMap {
		router.Handle(path, apiHandler(apiFunction))
	}
}

//	startServer defines logic that occurs when server starts.
func startServer(port int, apiMap map[string]func(http.ResponseWriter, *http.Request) *apiError) {
	mapAPIFunctions(http.NewServeMux(), apiMap)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

//
//	Server main entry point
//
func main() {
	initializeLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	startServer(port, map[string]func(http.ResponseWriter, *http.Request) *apiError{
		"/":       hello,
		"/health": health,
	})
}
