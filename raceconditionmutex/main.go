package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var protect sync.Mutex
var count int

//go run -race racecondition.go

func main() {
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go interncicle()
	}
	wg.Wait()
	fmt.Println("Count: ", count) // Resultat problematic
}

func interncicle() {
	for i := 1; i <= 10; i++ {
		protect.Lock()
		x := count
		x++
		count = x
		protect.Unlock()
	}
	wg.Done()
}
