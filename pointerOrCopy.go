package main

import "fmt"

func main() {
	ch := make(chan int)
	fun := func(ch1 chan int) {
		fmt.Println(&ch1)
		defer close(ch1)
	}
	fmt.Println(&ch)
	fun(ch)
	ch <- 5
}
