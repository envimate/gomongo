package main

import (
	"flag"
	"fmt"
	"gomongo/monitoring"
	"io"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2"
)

var port int
var nrpePort int
var mongourl string

func init() {
	flag.IntVar(&port, "port", 8080, "Port on which to listen")
	flag.StringVar(&mongourl, "mongourl", "mongodb://localhost:27017", "the mongodb connection url")
	flag.IntVar(&nrpePort, "nrpePort", 5667, "Port to listen for NRPE requests")
	flag.Parse()
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func defaultHandler(rw http.ResponseWriter, request *http.Request) {
	io.WriteString(rw, "handler called for URI \n")
	io.WriteString(rw, request.RequestURI)
}

func (mhandler *MongoHandler) cityIdHandler(rw http.ResponseWriter, request *http.Request) {
	cityColl := mhandler.db.C("city")
	query := request.URL.Query()
	cityId := query.Get("id")
	io.WriteString(rw, cityId+"\n")

	result := City{}

	id, _ := strconv.ParseInt(cityId, 10, 64)

	err := cityColl.FindId(id).One(&result)
	fmt.Printf("%+v", cityColl.FindId(id))
	if err != nil {
		io.WriteString(rw, "City with id "+cityId+" not found\n")
		log.Println(err)
		return
	}

	io.WriteString(rw, "Found city "+result.Name+"\n")
}

type City struct {
	Id   int64
	Name string
}

func main() {
	log.Println("Starting server on port", port)

	session, err := mgo.Dial(mongourl)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	db := session.DB("geo")

	mh := &MongoHandler{db}
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mh,
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))

	mux["/"] = defaultHandler
	mux["/city"] = mh.cityIdHandler

	nrpeServer := &monitoring.NRPEServer{
		Port: nrpePort,
	}

	go nrpeServer.StartServer()

	log.Fatal(s.ListenAndServe())
}

type MongoHandler struct {
	db *mgo.Database
}

func (mhandler *MongoHandler) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	if handler, ok := mux[request.URL.Path]; ok {
		handler(rw, request)
		return
	}
	io.WriteString(rw, "got request "+request.URL.String()+"\n")
}
