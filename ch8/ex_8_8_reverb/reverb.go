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
	wg := &sync.WaitGroup{}
	defer func() {
		wg.Wait()
		c.Close()
	}()

	lines := make(chan string)
	go scan(c, lines)

	const timeout = 10 * time.Second
	timer := time.NewTimer(timeout)
	for {
		select {
		case line := <-lines:
			timer.Reset(timeout)
			wg.Add(1)
			go echo(c, wg, line, 1*time.Second)
		case <-timer.C:
			return
		}
	}
}

func scan(c net.Conn, lines chan string) {
	func() {
		defer close(lines)
		input := bufio.NewScanner(c)
		for input.Scan() {
			lines <- input.Text()
		}
		if input.Err() != nil {
			log.Println(input.Err())
		}
	}()
}

func echo(c net.Conn, wg *sync.WaitGroup, shout string, delay time.Duration) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
