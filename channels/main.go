package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//ch := make(chan int, 100) // buffered channel

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	time.Sleep(1 * time.Second)
}
