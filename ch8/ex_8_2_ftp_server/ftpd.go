package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	cmd, err := net.Listen("tcp4", ":21")
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
	rw     net.Conn
	cmdErr error
}

func NewConn(c net.Conn) *conn {
	return &conn{c, nil}
}

func (c *conn) run() {
	defer c.rw.Close()
	c.writeln("220 Ready.")
	s := bufio.NewScanner(c.rw)

	var cmd string
	// var args []string
	for s.Scan() {
		if s.Err() != nil {
			log.Printf("scan: %v", s.Err())
			continue
		}
		fields := strings.Split(s.Text(), " ")
		if len(fields) == 0 {
			continue
		}
		for i, field := range fields {
			fields[i] = strings.TrimSpace(field)
		}

		cmd = strings.ToUpper(fields[0])
		// args = fields[1:]

		fmt.Println(cmd)
		switch cmd {
		case "QUIT":
			c.writeln("221 Goodbye.")
			return
		case "USER":
			c.writeln("230 Login successful.")
		default:
			c.writeln(fmt.Sprintf("502 Command %q not implemented.", cmd))
		}
	}
}

func (c *conn) writeln(s ...interface{}) {
	if c.cmdErr != nil {
		return
	}
	s = append(s, "\r\n")
	_, c.cmdErr = fmt.Fprint(c.rw, s...)
}
