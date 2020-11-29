package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Panic(err)
		}

		io.WriteString(os.Stdout, "Connection accepted")
		io.WriteString(conn, "Hello \n")
		io.WriteString(conn, "Hi! \n")

		conn.Close()
	}
}
