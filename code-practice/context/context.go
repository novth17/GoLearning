package main

import (
	"fmt"
	"sync"
	"time"
	"context"
)

//The orchestra can close sooner before all singer can finish their song, singers may stop early if the orchestra signals shutdow. This is where context comes into play
//Once a singer starts, no God can tell them to stop. You need context to tell them to shut up =))
func singers(ctx context.Context, name string, id int) {
	fmt.Println(name, id, "started singing")

	//select has to be in loop because cancellation is an event, not a one-time check.
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, id, "shutted up")
			
			return
		default:
			fmt.Println(name, id, "LALALALALA")
			time.Sleep(2 * time.Millisecond)
		}
	}
}

func main() {

	fmt.Println("Orchestra started...")
	// ctx (data type: context.Context) - The value stored inside ctx is a *cancelCtx!!!!
	/*
	cancelCtx
	├── done      (chan struct{}, maybe nil now. Being created lazily — only when someone calls ctx.Done().)
	├── err       (error)
	├── children  (other contexts)
	└── mutex
	*/

	var ctx context.Context
	var cancel context.CancelFunc

	// this function returns a Context AND a CancelFunc - read source code!!!
	// cancel is var that stores a closure function, this closure captures c, uses c becoz c is declared in outer scope
	// c is the *cancelCtx created inside WithCancel
	//context.Background() gives a root context, does nothing && never-canceled
	//WithCancel(parent) -> create a new child context *cancelCtx && link to parent
	// → returns: ctx → the child context
	// cancel → the closure that can cancel this child
	ctx, cancel = context.WithCancel(context.Background())

	fmt.Printf("Data type ctx: %T\n", ctx)    // *context.cancelCtx
	fmt.Printf("Data type cancel: %T\n", cancel) // context.CancelFunc


	var wg sync.WaitGroup
	amountSingers := 5

	for i := 1; i <= amountSingers; i++ {
		wg.Add(1)

		// about safe loop
		// safe 1.25 version fixed loop Variable Scoping https://go.googlesource.com/proposal/+/master/design/60078-loopvar.md
		//from one-instance-per-loop to one-instance-per-iteration. 
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			//fmt.Println("address of i: ", &i)
			singers(ctx, "Singer", i)
			//must listen to ctx.Done() via select, if not they will ignore the shut up yelling from main
		}()
	}
	
	//let them sing for 5 seconds then yell "SHUT UP!"
	time.Sleep(3 * time.Second)
	fmt.Println("Conductor: EVERYONE SHUT UP!")

	cancel() //ctx.Done() returns a channel, and cancel() closes that channel. → all goroutines receiving from it unblock immediately
	wg.Wait()
	fmt.Println("Orchestra closed. See you in the next concert!")
}