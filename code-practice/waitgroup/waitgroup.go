package main

import (
	"fmt"
	"sync"
	 "time"
)

func countA() {
	for i := 1; i <= 3 ; i++ {
		fmt.Printf("Thread A, number: %d\n", i)
		time.Sleep(10 * time.Millisecond)

	}
}

func countB() {
	for i := 1; i <= 3 ; i++ {
		fmt.Printf("Thread B, number: %d\n", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main () {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		countA()
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		countB()
	}()

	wg.Wait()
}