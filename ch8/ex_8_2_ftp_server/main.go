package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/yykhomenko/book-gopl/ch8/ex_8_2_ftp_server/ftp"
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
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		if sc.Err() != nil {
			log.Printf("scan: %v", sc.Err())
			continue
		}

		cmd := ftp.NewCmd(sc.Text())
		if cmd.Name == "" {
			continue
		}
		log.Printf("CMD %s", cmd)

		out, err := cmd.Exec()
		if err != nil {
			fmt.Fprintf(conn, "%s\n500\n", err)
		} else {
			fmt.Fprintf(conn, "%s\n200 ОК\n", out)
		}
	}
}
