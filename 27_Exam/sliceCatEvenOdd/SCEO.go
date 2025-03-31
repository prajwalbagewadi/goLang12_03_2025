package main

import "fmt"

func checkOddEven(arr []int, odd chan int, even chan int) {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			even <- arr[i]
		} else {
			odd <- arr[i]
		}
	}
	close(odd)
	close(even)
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Use buffered channels to prevent blocking
	oddch := make(chan int,len(nums))
	evench := make(chan int,len(nums))
	go checkOddEven(nums, oddch, evench)
	fmt.Println("Even numbers:")
	for e := range evench {
		fmt.Println("Even:", e)
	}
	fmt.Println("Odd numbers:")
	for o := range oddch {
		fmt.Println("odd:", o)
	}
}