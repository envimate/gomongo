package main

import (
	"log"
	"net/http"
	"flag"
	"fmt"
	"io"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "Port on which to listen")
	flag.Parse()
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func defaultHandler(rw http.ResponseWriter, request *http.Request){
	io.WriteString(rw, "handler called for URI \n")
	io.WriteString(rw, request.RequestURI)
}

func cityIdHandler(rw http.ResponseWriter, request *http.Request){
	query := request.URL.Query()
	io.WriteString(rw, query.Get("id")+"\n")
}

func main() {
	log.Println("Starting server on port", port)
	s := &http.Server {
		Addr: fmt.Sprintf(":%d", port),
		Handler: &MongoHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))

	mux["/"] = defaultHandler
	mux["/city"] = cityIdHandler

	log.Fatal(s.ListenAndServe())
}

type MongoHandler struct {
}

func (mhandler *MongoHandler) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	if handler, ok := mux[request.URL.Path]; ok {
		handler(rw, request)
		return
	}
	io.WriteString(rw, "got request "+request.URL.String()+"\n")
}

