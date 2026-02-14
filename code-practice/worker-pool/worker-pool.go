package main

import "fmt"

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
Channel = request queue
WaitGroup = graceful shutdown
*/


func worker(id int, orders <-chan int) {
	
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




}