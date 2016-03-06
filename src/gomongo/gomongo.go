package main

import (
	"log"
	"net/http"
	"flag"
	"fmt"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "Port on which to listen")
	flag.Parse()
}
func main() {
	log.Println("Starting server on port", port)
	s := &http.Server {
		Addr: fmt.Sprintf(":%d", port),
		Handler: &MongoHandler{},
	}

	log.Fatal(s.ListenAndServe())
}

type MongoHandler struct {
}

func (mhandler *MongoHandler) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	log.Println("got request", request.RequestURI)
}

