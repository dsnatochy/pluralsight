package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func(ch1 chan int, ch2 chan string, wg *sync.WaitGroup) {
		select {
		case i := <-ch1: // able to receive from channel is assign the value to "i"
			fmt.Println("received i:", i)
		case ch2 <- "hello": // there is a listener on the other end and we can send the value
			fmt.Println("sent 'hello' to ch2")
		default:
			print("nothing is ready in select")
		}
		time.Sleep(500 * time.Millisecond)
		wg.Done()
	}(ch1, ch2, wg)

	go func(ch1 chan int, ch2 chan string, wg *sync.WaitGroup) {
		//fmt.Println(<-ch2)
		//ch1 <- 5
		wg.Done()
	}(ch1, ch2, wg)

	wg.Wait()
}
