package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	cmd, err := net.Listen("tcp", ":21")
	if err != nil {
		log.Fatal()
	}

	for {
		conn, err := cmd.Accept()
		if err != nil {
			continue
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Fatal(err)
	}
}
