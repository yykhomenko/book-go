// Chat with nonblocking broadcast writes.
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

type Client struct {
	name string
	ch   chan string
}

func NewClient(name string) Client {
	return Client{name, make(chan string, 10)}
}

var (
	entering = make(chan Client)
	leaving  = make(chan Client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[Client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				select {
				case cli.ch <- msg:
				default:
				}
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
	client := NewClient(getName(conn))
	go clientWriter(conn, client.ch)
	go clientTimeout(conn, client.ch)
	client.ch <- "You " + client.name
	messages <- client.name + " connected"
	entering <- client

	var timer = time.AfterFunc(timeout, func() {
		client.ch <- "timeout"
		time.Sleep(100 * time.Millisecond)
		conn.Close()
	})
	defer timer.Stop()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		if input.Text() != "" {
			messages <- client.name + ": " + input.Text()
			timer.Reset(timeout)
		}
	}
	if input.Err() != nil {
		log.Print(input.Err())
	}

	leaving <- client
	messages <- client.name + " disconnected"
	conn.Close()
}

func getName(conn net.Conn) (name string) {
	conn.Write([]byte("Enter your name:"))
	input := bufio.NewScanner(conn)
	for input.Scan() {
		name = input.Text()
		if name == "" {
			conn.Write([]byte("name is empty, repeat:"))
			continue
		}
		return
	}
	if input.Err() != nil {
		log.Print(input.Err())
	}

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
