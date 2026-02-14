package main

import (
	"fmt"
	//"time"
)

func philo(leftFork chan bool, rightFork chan bool) {

	fmt.Println("Philo is hungry")

	//The current goroutine receives the bool, take it out
	// if there's value, take it out. If not the blocks.
	x :=<-leftFork 
	fmt.Println("Took left fork")

	y := <-rightFork
	fmt.Println("Took right fork")
	
	fmt.Println("Value left tossed to x : ", x)
	fmt.Println("Value right tossed to y : ", y)

	fmt.Println("Philo is eating with 2 forks")

	select {
		case v := <-leftFork:
			fmt.Println("leftFork was available:", v)
		default:
			fmt.Println("leftFork is currently taken")
	}

	//check if channel is empty, if withdrawal is possible then it was available, if not then taken
	select {
		case v := <-leftFork:
			fmt.Println("leftFork was available:", v)
		default:
			fmt.Println("leftFork is currently taken")
	}

	select {
		case v := <-rightFork:
			fmt.Println("rightFork was available:", v)
		default:
			fmt.Println("rightFork is currently taken")
	}
	//A channel is a reference type. printing it will print its address
	fmt.Println("Channel of left fork: ", leftFork)
	fmt.Println("Channel of right fork: ", rightFork)
	
	leftFork <- true
	rightFork <- true
}

//channel require a sender and a receiver that canb run concurrently
func main () {

	
	leftFork := make(chan bool, 1)
	rightFork:= make(chan bool, 1)

	leftFork <- true // put true, means available
	rightFork <- true

	philo(leftFork, rightFork);

}