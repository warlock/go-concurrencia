package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go waitText("Hello", 3)
	go waitText("World", 2)
	wg.Wait()
}

func waitText(text string, temps time.Duration) {
	for i := 1; i <= 10; i++ {
		time.Sleep(temps * time.Millisecond)
		fmt.Println(text, i)
	}
	wg.Done()
}
