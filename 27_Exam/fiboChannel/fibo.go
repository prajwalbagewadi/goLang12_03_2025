package main

import "fmt"

func fibo(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	n := 5
	ch := make(chan int)
	go fibo(n, ch)
	for num := range ch {
		fmt.Println(num)
	}
}