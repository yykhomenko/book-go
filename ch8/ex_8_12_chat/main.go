package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	name string
	c    chan<- string // outgoing channel
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all client messages
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages: // broadcast message
			for cli := range clients {
				cli.c <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			cli.c <- "all users:"
			for c := range clients {
				cli.c <- c.name
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.c)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You " + who
	messages <- who + " connected"
	entering <- client{who, ch}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	} // ignore input.Err()

	leaving <- client{who, ch}
	messages <- who + " disconnected"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
