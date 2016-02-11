// Package server provides logic to listen and serve content and faciliate
// a RESTful API.
package main

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/op/go-logging"
)

//
//	CONSTANT Variables
//

//	port defines the port on which the server will listen and serve.
const port int = 8000

//	serverHeader defines the value of the "Server" header set on HTTP responses.
const serverHeader string = "Hello Go Web Server"

//	log is the global logging object
var log = logging.MustGetLogger("main")

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

func initializeLogger() {
	//	format string which defines the log output format.
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} [%{level:.10s}] %{shortfunc} %{color:reset} %{message}`,
	)

	// 	Create two backend for os.Stderr.
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	//	For messages written to backend2 we want to add some additional
	// 	information to the output, including the used log level and the name of
	// 	the function.
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	//	Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")

	//	Set the backends to be used.
	logging.SetBackend(backend1Leveled, backend2Formatter)

	//	log usage:
	// log.Debug("debug")
	// log.Info("info")
	// log.Notice("notice")
	// log.Warning("warning")
	// log.Error("err")
	// log.Critical("crit")
}

//	ServeHTTP method calls the apiHandler function and displays the returned
//	error (if any).
func (apiHandlerFunction apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debugf("Request: %s ... ", r.URL.Path)
	if e := apiHandlerFunction(w, r); e != nil {
		log.Errorf("%+v", e)
		http.Error(w, e.Message, e.Code)
	} else {
		log.Debug("... Request received.")
	}
}

//	setServerHeader sets the response header.
func setServerHeader(w http.ResponseWriter) {
	w.Header().Set("Server", serverHeader)
}

//	mapAPIFunctions maps API handler functions to the API path, as specified in a mapping.
func mapAPIFunctions(apiMap map[string]func(http.ResponseWriter, *http.Request) *apiError) {
	router := http.NewServeMux()
	log.Debug("Mapping API functions ... ")
	for path, apiFunction := range apiMap {
		router.Handle(path, apiHandler(apiFunction))
	}
	log.Debug(" ... API functions mapped.")
}

//	startServer defines logic that occurs when server starts.
func startServer(port int, apiMap map[string]func(http.ResponseWriter, *http.Request) *apiError) {
	portString := ":" + strconv.Itoa(port)
	log.Infof("Starting server (port %s)... ", portString)
	mapAPIFunctions(apiMap)
	err := http.ListenAndServe(portString, nil)
	if err != nil {
		log.Critical("%v", err)
	}
	log.Info(" ... server started.")
}

//
//	Server main entry point
//
func main() {
	initializeLogger()
	log.Debug("DEBUG")
	startServer(port, map[string]func(http.ResponseWriter, *http.Request) *apiError{
		"/hello":  hello,
		"/health": health,
	})
}
