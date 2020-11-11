// GOOS=darwin GOARCH=amd64 go build
// GOOS=linux GOARCH=amd64 go build
// GOOS=windows GOARCH=amd64 go build
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
