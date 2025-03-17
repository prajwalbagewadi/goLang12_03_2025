package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int,wg *sync.WaitGroup){
	defer wg.Done() //Decrement the WaitGroup counter when goroutine completes
	fmt.Println("worker id=",id,"started...")
	time.Sleep(1*time.Second) //simulate work.
	fmt.Println("worker completed id:",id)
}

func main(){
	var swg sync.WaitGroup

	numworkers := 5
	swg.Add(numworkers) //Add workers count before starting goroutines

	for i:=0;i<numworkers;i++{
		//error deadlock
		//swg.Add(i)//Increment the WaitGroup counter for each goroutine
		go worker(i,&swg)  // Start a goroutine
	}

	swg.Wait() //Wait for all goroutines to finish
	fmt.Println("all go routines finished.")
	fmt.Println("main func finished exec.")
}