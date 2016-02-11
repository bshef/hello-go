// Package server provides logic to listen and serve content and faciliate
// a RESTful API.
package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"text/template"

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

//	templates is a data structure containing parsed HTML template files.
var templates = template.Must(template.ParseFiles("./view/view.html"))

//	validPath is a regex pattern used to mitigate potential XSS risks.
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

//
//	Type definitions and constructors
//

// apiHandler is the standard function for handling API requests.
type apiHandler func(http.ResponseWriter, *http.Request) error

// Page is a struct containing information relevant to each page.
type Page struct {
	Title string
	Body  []byte
}

//
//	API Handler Functions
//

//	handleRoute not only maps a function to an API route,
//	but also defines standard behavior for every API request.
func handleRoute(route string, handler apiHandler) {
	handlerFunction := func(w http.ResponseWriter, r *http.Request) {
		log.Infof("API Request:\t%s", r.URL.String())
		setServerHeader(w)
		if err := handler(w, r); err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.HandleFunc(route, handlerFunction)
}

//	hello responds with "Hello, world!"
func hello(w http.ResponseWriter, r *http.Request) error {
	setServerHeader(w)
	io.WriteString(w, "Hello, world!")
	return nil
}

//	health responds with 200 OK
func health(w http.ResponseWriter, r *http.Request) error {
	setServerHeader(w)
	w.WriteHeader(http.StatusOK)
	return nil
}

//
//	HTML templating functions
//

//	save takes as its receiver a pointer to Page, and saves Page.Body to a text file.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//	loadPage reads the file's contents and returns a constructed Page literal
func loadPage(path string, title string) (*Page, error) {
	filename := path + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

//	renderTemplate renders an HTML template file with the specified Page data.
func renderTemplate(w http.ResponseWriter, templateFile string, page *Page) {
	err := templates.ExecuteTemplate(w, templateFile+".html", page)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//	viewHandler handles URLs prefixed with "/view/"
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := loadPage("./view/"+title, title)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		renderTemplate(w, "view", page)
	}
}

//	makeHandler returns a handler function if a valid title is requested.
func makeHandler(handlerFn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			err := "Invalid Page Title"
			log.Error(err)
			http.Error(w, err, http.StatusInternalServerError)
			return
		}
		handlerFn(w, r, m[2])
	}
}

//	serveTemplatedContent serves templated content from specified directories.
func serveTemplatedContent() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
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
	//	log.Debug("debug"), log.Info("info"), log.Notice("notice")
	//	log.Warning("warning"), log.Error("err"), log.Critical("crit")
	return
}

//	setServerHeader sets the response header.
func setServerHeader(w http.ResponseWriter) {
	w.Header().Set("Server", serverHeader)
	return
}

//	serveStaticContent serves static content based on the requested path.
func serveStaticContent() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
}

//	setupRoutes sets up routes and paths handled by the server.
func setupRoutes() {
	//	Serve static content from /public directory
	serveStaticContent()

	//	Serve templated content directory
	serveTemplatedContent()

	//	Handle basic API routes
	handleRoute("/hello", hello)
	handleRoute("/health", health)
	return
}

//	startServer defines logic that occurs when server starts.
func startServer(port int) {
	portString := ":" + strconv.Itoa(port)
	log.Infof("Starting server (port %s) ... ", portString)

	//	Actually start the server
	err := http.ListenAndServe(portString, nil)
	if err != nil {
		log.Critical("%v", err)
	}
	return
}

//
//	Server main entry point
//
func main() {
	initializeLogger()
	setupRoutes()
	startServer(port)
}
