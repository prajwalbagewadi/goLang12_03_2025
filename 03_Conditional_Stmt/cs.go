package main

import (
	"fmt"
)

func main(){
	num := 10
	
	// if-else 
	if num > 0 {
		fmt.Println("positive number")
	} else {
		fmt.Println("negative number")
	}

	// switch 
	day := "mon" 
	switch(day){
		case "mon","tue","wed","thus","fri":
			fmt.Println("Working day")
		case "sat","sun":
			fmt.Println("holiday")
		default:
			fmt.Println("invalid")

	}

	//nested
}