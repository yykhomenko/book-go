package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", "8080", "server port")
	flag.Parse()
}

func main() {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Second)
	}
}
