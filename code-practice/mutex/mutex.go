package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	//"time"
)

func increment(i *int32) {
	atomic.AddInt32(i, 10)
}

func main() {

	var wg sync.WaitGroup
	var mtx sync.Mutex

	var number int32
	number = 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mtx.Lock()
			increment(&number)
			fmt.Printf("Address of number in grt: %p\n", &number)
		
			fmt.Printf("Value number: %d\n", atomic.LoadInt32(&number))
			mtx.Unlock()
		}()
	}
	wg.Wait()
}