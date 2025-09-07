package main

import (
	"fmt"
	"httpfromtcp/internal/request"
	"log"
	"net"
)
func main() {
	listener, err := net.Listen("tcp", ":42069")

	if err != nil {
		log.Fatal("error", "error", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}
		 r, err := request.RequestFromReader(conn)
		 if err != nil {
			log.Fatal("error", "error", err)
		}

		fmt.Printf("Request Line:\n")
		fmt.Printf("- Method: %s\n", r.RequestLine.Method)
		fmt.Printf("- Target: %s\n", r.RequestLine.RequestTarget)
		fmt.Printf("- Version: %s\n", r.RequestLine.HttpVersion)
	}
}
