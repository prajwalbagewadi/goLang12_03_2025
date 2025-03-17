package main

import (
	"fmt"
	"time"
)

func main() {
	//anonymous function (lambda)
	/*
	An anonymous function (also called a lambda function) is a function without a name. 
	It can be defined inline and used immediately or assigned to a variable for later use.
	*/
	result1 := func(a, b int) int {
		return a + b
	}(5, 5) //Actual parameters

	go func(a, b int) {
		fmt.Println("go anonymous a+b=",a+b)
	}(3, 3) //Actual parameters
	time.Sleep(1*time.Second)

	fmt.Println("result1:",result1)
	fmt.Println("main func finished exec.")
}

