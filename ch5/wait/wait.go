// Wait waits for server available until timeout comes.
// go run wait.go http://golang.or
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := WaitForServer(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retry...", err)
		time.Sleep(time.Second << uint(tries))
	}

	return fmt.Errorf("server %s not responding; time %s", url, timeout)
}
