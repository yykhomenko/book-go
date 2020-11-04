package main

import (
	"fmt"
	"log"
	"net"
	"strings"
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
	for {
		var buf []byte
		conn.Read(buf)
		cmd := NewCmd(string(buf))
		fmt.Fprintf(conn, "%v", cmd)
	}
}

type Cmd struct {
	cmd  string
	args []string
}

func NewCmd(line string) *Cmd {
	words := strings.Split(line, " ")
	cmd := strings.TrimSpace(words[0])
	args := words[1:]
	for i, arg := range args {
		args[i] = strings.TrimSpace(arg)
	}
	return &Cmd{cmd, args}
}
