// Chat with username.
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
	go clientTimeout(conn, ch)

	who := getName(conn)
	client := client{who, ch}
	client.ch <- "You " + who
	messages <- who + " connected"
	entering <- client

	var timer = time.AfterFunc(timeout, func() {
		ch <- "timeout"
		time.Sleep(100 * time.Millisecond)
		conn.Close()
	})
	defer timer.Stop()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		if input.Text() != "" {
			messages <- who + ": " + input.Text()
		}
		timer.Reset(timeout)
	} // ignore input.Err()

	leaving <- client
	messages <- who + " disconnected"
	conn.Close()
}

func getName(conn net.Conn) (name string) {
	input := bufio.NewScanner(conn)
	conn.Write([]byte("Enter your name:"))
	for input.Scan() {
		name = input.Text()
		if name == "" {
			conn.Write([]byte("name is empty, repeat:"))
			continue
		}
		return
	} // ignore input.Err()
	return
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func clientTimeout(conn net.Conn, ch chan<- string) {
	var timer *time.Timer
	timer = time.AfterFunc(5*time.Minute, func() {
		ch <- "timeout"
		time.Sleep(100 * time.Millisecond)
		conn.Close()
		timer.Stop()
	})
}
