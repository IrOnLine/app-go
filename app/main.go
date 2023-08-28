package main

import (
	"app/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// Initialize logger
var log = logrus.New()

// handleHello GET /hello
func handleHello(w http.ResponseWriter, r *http.Request) {

	//	log.Println(r.Method, r.RequestURI)

	// Returns hello world! as a response
	fmt.Fprintln(w, "Hello world!")
}

func main() {
	// Load environment variables
	config.Load()

	// registers handleHello to GET /hello
	http.HandleFunc("/hello", handleHello)
	// starts the server on port 5000
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatalln(err)
	}
}
