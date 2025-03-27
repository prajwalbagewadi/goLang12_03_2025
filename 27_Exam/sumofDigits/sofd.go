package main

import "fmt"

func sumofDigits(n int) int {
	if n == 0 {
		return 0
	}
	return (n % 10) + sumofDigits(n/10)
}
func main() {
	n := 123
	res := sumofDigits(n)
	fmt.Println("sum=",res)
}