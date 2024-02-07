package main

import (
	"fmt"
)

func Select3() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)
	c3 := make(chan interface{})
	close(c3)

	go func() {
		var c1Count, c2Count, c3Count int
		defer wg.Done()
		for i := 1000; i >= 0; i-- {
			select {
			case <-c1:
				c1Count++
			case <-c2:
				c2Count++
			case <-c3:
				c3Count++
			}
		}
		fmt.Printf("c1Count: %d\nc2Count: %d\n c3: %d", c1Count, c2Count, c3Count)
	}()

}
