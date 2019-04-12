package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Response struct {
	Message string
	Ok      bool
}

type Request struct {
	Name string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	request := Request{Name: "wilson"}
	response := new(Response)
	err = client.Call("Handler.Do", request, response)

	if err != nil {
		return
	}

	fmt.Println(response.Message)
}
