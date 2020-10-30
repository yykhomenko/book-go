package main

import (
	"fmt"
	"time"
)

type Rocket struct{}

func (r *Rocket) Lunch() {
	fmt.Println("lunch!")
}

func main() {
	r := new(Rocket)
	time.AfterFunc(2*time.Second, func() { r.Lunch() })
	time.AfterFunc(3*time.Second, r.Lunch)
	time.Sleep(4 * time.Second)
}
