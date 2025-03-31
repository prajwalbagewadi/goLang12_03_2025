package main

import "fmt"

func fibo(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		
		ch <- a //send data to channel
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	n := 5
	ch := make(chan int) // create channel
	go fibo(n, ch) // start goroutine send channel into fibo
	for num := range ch { // receive data from channel
		fmt.Println(num) 
	}
}