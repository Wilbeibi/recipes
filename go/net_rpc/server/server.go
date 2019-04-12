package main

import (
	"errors"
	"flag"
	"log"
	"net"
	"net/rpc"
	"strconv"
	"time"
)

var (
	port = flag.Uint("port", 3000, "port to listen for rpc calls")
)

type Server struct {
	listener net.Listener
	port     uint
	sleep    time.Duration
}

func NewServer(port uint) (*Server, error) {
	server := new(Server)
	var err error
	server.listener, err = net.Listen("tcp", ":"+strconv.Itoa(int(port)))
	if err != nil {
		return nil, err
	}
	server.port = port
	server.sleep = 2 * time.Second
	return server, nil
}

type Handler struct {
	Sleep time.Duration
}

type Response struct {
	Message string
	Ok      bool
}

type Request struct {
	Name string
}

func (h *Handler) Do(req Request, res *Response) (err error) {
	if req.Name == "" {
		err = errors.New("A name must be specified")
		return
	}
	log.Println("Accepting request")
	if h.Sleep != 0 {
		time.Sleep(h.Sleep)
	}

	res.Ok = true
	res.Message = "Hello " + req.Name
	log.Println("Finish request")
	return
}

func (server *Server) Serve() {
	rpc.Register(&Handler{
		Sleep: server.sleep,
	})
	rpc.Accept(server.listener)
}

func main() {
	flag.Parse()
	log.Printf("Server is listening on port %d ...\n", *port)
	server, err := NewServer(*port)
	if err != nil {
		log.Fatal(err)
	}
	server.Serve()
}
