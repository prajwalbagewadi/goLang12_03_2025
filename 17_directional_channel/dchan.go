package main

import (
	"fmt"
)

func sendOnlyChannel(ch chan <- int){ // Send-Only Channel
	ch <- 100 
	//fmt.Println(<-ch)
	fmt.Println("sendOnlyChannel started.")
	close(ch) //Close channel
}

func receiveOnlyChannel(ch <- chan int) { //Receive-Only Channel
	//ch <- 200
	fmt.Println("receiveOnlyChannel channel data:",<-ch)
}

func main(){
	ch := make(chan int) //bidirectional channel
	go func(){
		ch<-120 //send data
	}()
	val1:= <-ch //receive data
	fmt.Println("bidirectional channel=",val1)

	//send only and receive only data channel
	// eg :Practical Use Case: Pipeline Processing
	channel := make(chan int)
	//go func(){channel<-100}() 
	go sendOnlyChannel(channel);
	receiveOnlyChannel(channel)
}