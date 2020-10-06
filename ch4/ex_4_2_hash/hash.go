// Hash calculates SHA digest.
// go run hash.go example // 50d858e0985ecc7f60418aaf0cc5ab587f42c2570a884095a9e8ccacd0f6545c
// go run hash.go -sha384 example // 4a55b0b4e06f9567b89623a1875d454a99e5ec9c930439f51ffc95dd43da14e433383f09f99463a9de680f16a3db796b
// go run hash.go -sha512 example // 8533a0342a78d5e0f1284e1d7ce4e130464cebe4b7b298d7fc2f00fecfb090813aea9b12cb437052ade5115a657904543cd6cebedc021a4b3aec4fd8c0730390
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var f384 = flag.Bool("sha384", false, "calculate 384 sum")
var f512 = flag.Bool("sha512", false, "calculate 512 sum")

func main() {
	flag.Parse()
	switch {
	case len(os.Args) == 1:
		fmt.Fprintln(os.Stderr, "Input one argument")
		os.Exit(1)
	case *f384:
		fmt.Printf("%x\n", sha512.Sum384([]byte(os.Args[1])))
	case *f512:
		fmt.Printf("%x\n", sha512.Sum512([]byte(os.Args[1])))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(os.Args[1])))
	}
}
