package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var count int

//go run -race racecondition.go

func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go interncicle()
	}
	wg.Wait()
	fmt.Println("Count: ", count) // Resultat problematic
}

func interncicle() {
	for i := 1; i <= 10; i++ {
		x := count
		x++
		count = x
	}
	wg.Done()
}
