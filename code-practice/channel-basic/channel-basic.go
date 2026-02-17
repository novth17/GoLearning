package main

import (
	"fmt"
	//"sync"
)

func main() {

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

		close(ch)

	x := <-ch //can still receieve after channel is closed
	fmt.Println(x)

	//ch <- 99 //The program stops immediately here && panic, send to closed channel
	y := <-ch 
	fmt.Println(y)

	z := <-ch
	fmt.Println(z) //0: wtf???

	i := <-ch
	fmt.Println(i)

	/*
	Lesson: Receiving from a closed channel NEVER blocks. It returns the zero value of the type. And it will keep doing that forever.
	https://github.com/golang/go/blob/master/src/runtime/chan.go starting from line 521 
	When a channel is closed, it clear the memory at ep to the zero value of the channel’s element type.
	So:
		chan int → 0
		chan string → ""
		chan bool → false
		chan struct{} → empty struct
	*/
}