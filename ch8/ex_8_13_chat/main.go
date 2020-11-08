// Chat with 5min timeout.
// go run main.go
// telnet localhost 8000
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const timeout = 5 * time.Minute

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
	ch   chan<- string // outgoing channel
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
				cli.ch <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			cli.ch <- "all users:"
			for c := range clients {
				cli.ch <- c.name
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
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

	var timer = time.AfterFunc(timeout, func() {
		ch <- "timeout"
		time.Sleep(100 * time.Millisecond)
		conn.Close()
	})

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(timeout)
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
