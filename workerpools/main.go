package main

import (
	"math/rand"
	"fmt"
	"time"
)

func echoWorker(in, out chan int) {
	for {
		n := <- in
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		out <- n
	}
}

func produce(in chan <- int) {
	i := 0
	for {
		fmt.Printf("-> Send job: %d\n", i)
		in <- i
		i++
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)

	for i:= 0; i < 4; i++ {
		go echoWorker(in, out)
	} 
	go produce(in)

	for n := range out {
		fmt.Printf("<- Receive job: %d\n", n)
	}
}