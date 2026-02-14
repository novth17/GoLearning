package main

import (
	"fmt"
	"sync"
)

//"time"

/*
You have many tasks.
You have limited workers.
Workers pull tasks from a shared queue.

10 pizza orders arrive.
3 pizza makers.
There is one order counter (channel)

Real Backend Mapping

Pizza factory = HTTP server
Pizza Orders = incoming requests
Workers = goroutines handling jobs
Channel of Orders = request queue
WaitGroup = graceful shutdown
*/


func worker(id int, orders <-chan int) {

	//keep receving order until channel is closed. if channel dont close, stuck here!!!
	for order := range orders {
		fmt.Println("Worker", id, "is making pizza", order)
	}
}

func main() {

	//main take customer orders
	//put them on counter
	
	//Workers (goroutines): Stand by the counter waiting.

	//They just take the next order available.
	//If no pizzas are on the counter:
	//Workers wait.

	orders := make(chan int)
	
	var wg sync.WaitGroup
	for workerId := 1; workerId <= 5; workerId++ {
		wg.Add(1)
		go func(workerId int) { 
			defer wg.Done()
			worker(workerId, orders)
		}(workerId)
	}

	for order := 1; order <= 5 ; order++ {
		orders <- order
	}
	close(orders) //if dont close, hang!range receiving orders until the channel is closed. they block forever waiting for more pizza order.
	wg.Wait()
}