package main

import (
	"fmt"
	"time"
)

var MaxOutstanding = 10

type Request struct {
	message string
}

var sem = make(chan int, MaxOutstanding)

func main() {
	queue := make(chan *Request)
	defer close(queue)

	go Serve(queue)

	for i := 1; i <= 25; i++ {
		request := &Request{message: fmt.Sprintf("req-%d", i)}
		fmt.Printf("Send request %+v\n", request)
		queue <- request
		time.Sleep(200 * time.Millisecond)
	}
}

func process(r *Request) {
	fmt.Printf("Processing... %v\n", r.message)
	time.Sleep(5 * time.Second)
	fmt.Printf("Done! %v\n", r.message)
}

func Serve(queue chan *Request) {
	for req := range queue {
		req := req // Create new instance of req for the goroutine.

		sem <- 1
		go func() {
			process(req)
			<-sem
		}()
	}
}
