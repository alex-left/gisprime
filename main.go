package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// Global config
var servePort = "8000"

func processFlags() {
	port := flag.String("p", servePort, "Listen port")
	flag.Parse()
	if *port != servePort {
		servePort = *port
	}
}

func processEnvVars() {
	var port = ""
	if port = os.Getenv("HTTP_PORT"); port != "" {
		servePort = port
	}
}

func processGlobalConfig() {
	processFlags()
	processEnvVars()
}

func checkPort() bool {
	port, err := strconv.Atoi(servePort)
	if err != nil {
		return false
	}
	return port < 65536 && port > 0
}

func setRouter() *mux.Router {
	var router = mux.NewRouter()
	router.HandleFunc("/{number:[0-9]+}", numHandler).Methods("GET")
	router.HandleFunc("/health", healthCheckHandler).Methods("GET")
	return router
}

func main() {
	processGlobalConfig()
	if !checkPort() {
		log.Fatal("Invalid port: ", servePort)
	}
	http.Handle("/", setRouter())
	log.Printf("Starting server at port %s\n", servePort)
	log.Fatal(http.ListenAndServe(":"+servePort, nil))
}
