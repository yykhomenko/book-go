package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	cmd, err := net.Listen("tcp", ":21")
	if err != nil {
		log.Fatal()
	}

	for {
		conn, err := cmd.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go NewConn(conn).run()
	}
}

type conn struct {
	rw net.Conn
}

func NewConn(c net.Conn) *conn {
	return &conn{c}
}

func (c conn) run() {
	defer c.rw.Close()
	sc := bufio.NewScanner(c.rw)
	for sc.Scan() {
		if sc.Err() != nil {
			log.Printf("scan: %v", sc.Err())
			continue
		}

		cmd := NewCmd(sc.Text())
		if cmd.Name == "" {
			continue
		}
		log.Printf("CMD %s", cmd)

		out, err := cmd.Exec()
		if err != nil {
			fmt.Fprintf(conn, "%s\n500\n", err)
		} else {
			fmt.Fprintf(conn, "%s\n220 ОК\n", out)
		}
	}
}

func exec() {
}
