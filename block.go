package main

import (
	"fmt"
	"sync"
	"time"
)

/*
The goroutine that owns a channel should:
1. Instantiate the channel.
2. Perform writes, or pass ownership to another goroutine.
3. Close the channel.
4. Ecapsulate the previous three things in this list and expose them via a reader
channel.

As a consumer of a channel, I only have to worry about two things:
• Knowing when a channel is closed.
• Responsibly handling blocking for any reason.
*/
func Main() {
	var dataStream chan int
	var wg sync.WaitGroup

	dataStream = make(chan int)
	a := func(wg *sync.WaitGroup) {
		defer wg.Done()
		dataStream <- 5
		fmt.Println("Sending")
	}
	b := func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		fmt.Println(<-dataStream)
	}
	wg.Add(2)
	go a(&wg)
	go b(&wg)
	wg.Wait()
}
