package main

import (
	"fmt"
)

var globalvar int = 50

func fun(){
	//fmt.Println("localvar=",localvar)
	fmt.Println("globalvar=",globalvar)
}

func main(){
	var localvar int = 10
	fmt.Println("localvar=",localvar)
	fmt.Println("globalvar=",globalvar)
	fun()
}