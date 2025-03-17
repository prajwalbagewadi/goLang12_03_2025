package main

import (
	"fmt"
	"time"
)

func main(){
	
	unbuff := make(chan int) //unbuffered channel
	
	go func(){
		unbuff <- 10
	}()
	time.Sleep(1*time.Second)

	unbuff_val := <-unbuff
	fmt.Println("unbuffered value:",unbuff_val)

	buff := make(chan int,2) //buffered channel with size 2
	buff <- 20
	buff <- 30
	fmt.Println("buffered Channel val:",<-buff)
	fmt.Println("buffered Channel val:",<-buff)
}