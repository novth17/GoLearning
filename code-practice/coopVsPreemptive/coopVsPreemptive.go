package main

import (
	"fmt"
	"sync"
	//"time"
)

func infinityLoop() {
	for {}

}

func anotherFunc() {
	fmt.Println("pupu CHI ice cream");
}

func main () {

	var wg sync.WaitGroup

	total := 10
	
	wg.Add(1)
	go func (){
		defer wg.Done()
		infinityLoop()
	}()

	for i := 0; i < total; i++ {
		wg.Add(1)

		go func(i int) { 
			defer wg.Done()
			anotherFunc()
		}(i)
	}

	wg.Wait()
}