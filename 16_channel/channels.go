/*
Channels in Go
What Are Channels?
Channels in Go are used to communicate between goroutines.
They allow goroutines to send and receive data safely.
Channels help synchronize execution without using explicit locks or shared memory.

*/

package main

import (
	"fmt"
	"time"
)

func worker(ch chan string){
	time.Sleep(2*time.Second)
	ch <- "data send task Completed."
}

func main(){
	var ch chan int //declare channel of type int
	ch = make(chan int) //initialize the channel

	//send and receive data using lambda func
	go func(){
		ch <- 4 //sending data to channel
	}()
	time.Sleep(1*time.Second)
	data := <- ch
	fmt.Println("data recived from ch=",data)

	//channel with function
	ch1 := make (chan string)

	go worker(ch1) //start worker with go routine
	fmt.Println("Waiting for Worker to complete...")
	msg := <-ch1 //Receiving msg
	fmt.Println("msg received from worker():",msg)

	fmt.Println("main function completed.")
}