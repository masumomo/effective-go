package main

import (
	"fmt"
	"time"
)

var MaxOutstanding = 10

type Request struct {
	message string
}

var quit = make(chan bool)

func main() {
	queue := make(chan *Request)
	defer close(queue)

	go Serve(queue, quit)

	for i := 1; i <= 25; i++ {
		request := &Request{message: fmt.Sprintf("req-%d", i)}
		fmt.Printf("Send request %+v\n", request)
		queue <- request
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("Sending a quit signal")
	quit <- true
}

func process(r *Request) {
	fmt.Printf("Processing... %v\n", r.message)
	time.Sleep(5 * time.Second)
	fmt.Printf("Done! %v\n", r.message)
}

func handle(queue chan *Request) {
	for r := range queue {
		process(r)
	}
}

func Serve(clientRequests chan *Request, quit chan bool) {
	// Start handlers
	for i := 0; i < MaxOutstanding; i++ {
		go handle(clientRequests)
	}
	<-quit // Wait to be told to exit.
}
