package main

import (
	"fmt"
	"sync"
)

func Select() {
	c1 := make(chan interface{})
	//close(c1)
	c2 := make(chan interface{})

	//close(c2)
	var wg sync.WaitGroup
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		var c1Count, c2Count, c3Count int
		defer wg.Done()
		for i := 1000; i >= 0; i-- {
			select {
			case <-c1:
				c1Count++
			case <-c2:
				c2Count++
			default:
				c3Count++
			}
		}
		fmt.Printf("c1Count: %d\nc2Count: %d\n c3: %d", c1Count, c2Count, c3Count)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 1000; i >= 0; i-- {
			c1 <- 4
			//time.Sleep(time.Millisecond * 10)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 1000; i >= 0; i-- {
			c2 <- 4
			//time.Sleep(time.Millisecond * 20)
		}
	}(&wg)
	wg.Wait()
}
