package main

import (
	"fmt"
	"time"
)

func dain() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)
	select {
	case <-c1:
		fmt.Println("hello top")
	case <-c2:
		fmt.Println("Hello bottom")
	default:
		time.Sleep(time.Second * 2)
	}

}
