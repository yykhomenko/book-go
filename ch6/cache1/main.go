package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex
	mapping = make(map[string]string)
)

func LookUp(key string) string {
	mu.Lock()
	defer mu.Unlock()
	return mapping[key]
}

func main() {
	fmt.Println(LookUp(""))
}
