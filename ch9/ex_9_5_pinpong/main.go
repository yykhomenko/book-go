package main

func main() {
	pinc := make(chan string)
	pongc := make(chan string)

	go func() {
		pinc <- "pin"
	}()

	go pin(pinc, pongc)
	pong(pongc, pinc)
}

func pin(out chan<- string, in <-chan string) {
	for range in {
		out <- "pin"
	}
}

func pong(out chan<- string, in <-chan string) {
	for range in {
		out <- "pong"
	}
}
