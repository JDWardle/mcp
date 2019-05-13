package main

import (
	"fmt"
	"log"
	"net"

	"github.com/JDWardle/gocraft/server"
)

func main() {
	l, err := net.Listen("tcp", ":25565")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	id := 0
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		id++

		// TODO Keep track of all clients connected to the server.
		fmt.Printf("new connection from %s\n", conn.RemoteAddr())
		c := server.NewClient(id, conn)
		go c.HandleMessages()
	}
}
