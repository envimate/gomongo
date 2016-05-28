package monitoring

import (
	"fmt"
	"net"

	"log"

	"github.com/envimate/nrpe"
)

var mux map[string]commandHandler

type commandHandler func(args []string) *nrpe.CommandResult

func requestCountHandler(args []string) *nrpe.CommandResult {
	result := &nrpe.CommandResult{
		StatusLine: "0",
		StatusCode: nrpe.StatusUnknown,
	}

	return result
}

func nrpeHandler(c nrpe.Command) (*nrpe.CommandResult, error) {
	// handle nrpe command here
	var handler commandHandler
	var ok bool
	if handler, ok = mux[c.Name]; !ok {
		log.Printf("Got unknown command %s", c.Name)
		return nil, fmt.Errorf("Unknown command received %+v", c)
	}

	return handler(c.Args), nil
}

func connectionHandler(conn net.Conn) {
	defer conn.Close()
	nrpe.ServeOne(conn, nrpeHandler, true, 0)
}

type NRPEServer struct {
	Port int
}

func init() {
	mux = make(map[string]commandHandler)

	mux["requestCount"] = requestCountHandler
}

func (nrpeServer *NRPEServer) StartServer() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", nrpeServer.Port))

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		go connectionHandler(conn)
	}
}
