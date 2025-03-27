//Write a program in GO language to print the Fibonacci series of n terms.
/*
We start with the first two Fibonacci numbers:

a = 0

b = 1

We iterate from 0 to n-1 and print a in each step.

We update values using swapping technique:

a = b (move to the next Fibonacci number)

b = a + b (compute the next Fibonacci number)

This avoids recursion and repeated calculations, making it fast and efficien
*/

package main

import "fmt"

func fibo(n int) {
	var a, b int = 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}

func main() {
	var n int
	fmt.Println("Enter the val for n to find Fibo")
	fmt.Scan(&n)
	fibo(n)
}
