package main

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalln(errors.New("must be at least one param"))
	}

	for _, arg := range os.Args[1:] {
		kv := strings.Split(arg, "=")
		name, addr := kv[0], kv[1]
		clock(name, addr)
	}
}

func clock(name, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	mustCopy(os.Stdout, conn)
}

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}
