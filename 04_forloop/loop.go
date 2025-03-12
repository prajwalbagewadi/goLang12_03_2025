package main

import (
	"fmt"
)

func main(){
	//loops

	//for 
	for i:=0;i<5;i++{
		fmt.Println("i=",i)
	}

	//while 
	i:=5
	for i >=1 {
		fmt.Println("i=",i)
		i--
	}

	//do while
	i=0
	for{ // for goes to infinty acting as do.
		fmt.Println("i=",i)
		i++
		if i > 5 {
			break
		}
	}

	
}