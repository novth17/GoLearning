package main

import (
	"fmt"
	"time"
)

func countA() {
	for i := 1; i <= 10000000 ; i++ {
		fmt.Printf("Thread A, number: %d\n", i)
	}
}

func countB() {
	for i := 1; i <= 10000000 ; i++ {
		fmt.Printf("Thread B, number: %d\n", i)
	}
}

func main () {
	go countA();
	go countB();
	time.Sleep(0 * time.Millisecond)
}