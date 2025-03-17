package main

import (
	"fmt"

	"eg.com/moduleB"
)

func main(){
	fmt.Println("moduleA running...")
	moduleB.HelperFunction()
}