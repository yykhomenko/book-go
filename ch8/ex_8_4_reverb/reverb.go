package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	wg := &sync.WaitGroup{}
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go echo(c, wg, input.Text(), 1*time.Second)
	}
	wg.Wait()
	if cw, ok := c.(*net.TCPConn); ok {
		cw.CloseWrite()
	} else {
		c.Close()
	}
}

func echo(c net.Conn, wg *sync.WaitGroup, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done()
}
