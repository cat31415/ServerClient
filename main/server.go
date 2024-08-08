package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("Connected")
	conn.Write([]byte("OK"))
	defer conn.Close()
}
