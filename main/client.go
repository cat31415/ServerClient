package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("incorrect usage")
		os.Exit(1)
	}
	server := os.Args[1]
	conn, err := net.Dial("tcp", server)

	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}

	fmt.Println("Connected to", server)
	defer conn.Close()
	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from server:", err)
	}
	res := string(buf[:n])
	if res == "OK" {
		fmt.Println(string(buf))
	} else {
		fmt.Println("NOT OK")
	}

}
