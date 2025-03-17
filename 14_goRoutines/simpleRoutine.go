package main

import (
	"fmt"
	"time"
)

func greet(name string){
	fmt.Println("hello. ",name)
}
func main(){
	go greet("GO Developer")
	time.Sleep(1*time.Second)
	/*
	By default, the main function does not wait for goroutines, so time.Sleep() gives them time to execute.
	*/
	fmt.Println("main function finished.")
}