package main

import (
	"fmt"
	"sync"
	"time"
	"context"
)

//The orchestra can close sooner before all singer can finish their song, singers may stop early if the orchestra signals shutdow. This is where context comes into play
//Once a singer starts, no God can tell them to stop. You need context to tell them to shut up =))
func singers(name string, id int) {
	fmt.Println(name, id, "started singing: LALALALALA")
}

func main() {

	fmt.Println("Orchestra started...")
	/*
	// ctx (data type: context.Context) - The value stored inside ctx is a *cancelCtx!!!!
	cancelCtx
	├── done      (chan struct{}, maybe nil now. Being created lazily — only when someone calls ctx.Done().)
	├── err       (error)
	├── children  (other contexts)
	└── mutex
	*/

	var ctx context.Context
	var cancel context.CancelFunc

	// this function returns a Context AND a CancelFunc
	// cancel is var that stores a closure function, this closure captures c, uses c becoz c is declared in outer scope
	// c is the *cancelCtx created inside WithCancel
	ctx, cancel = context.WithCancel(context.Background()) //returns an empty, root context.

	fmt.Printf("Data type ctx: %T\n", ctx)    // *context.cancelCtx
	fmt.Printf("Data type cancel: %T\n", cancel) // context.CancelFunc


	var wg sync.WaitGroup
	amountSingers := 2

	for i := 1; i <= amountSingers; i++ {
		wg.Add(1)

		// about safe loop
		// safe 1.25 version fixed loop Variable Scoping https://go.googlesource.com/proposal/+/master/design/60078-loopvar.md
		//from one-instance-per-loop to one-instance-per-iteration. 
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			//fmt.Println("address of i: ", &i)
			singers("Singer", i)
		
			//must listen to ctx.Done() via select, if not they will ignore the shut up yelling from main


		
		}()
	}
	
	//wait 5ms then yell "SHUT UP!"
	time.Sleep(5 * time.Millisecond)
	fmt.Println("Conductor: SHUT UP!")

	cancel() //ctx.Done() returns a channel, and cancel() closes that channel. → all goroutines receiving from it unblock immediately
	wg.Wait()
	fmt.Println("Orchestra closed. See you in the next concert!")
}