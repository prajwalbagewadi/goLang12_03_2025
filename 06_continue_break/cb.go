package main

import (
	"fmt"
)

func main(){
	//break
	for i:=0;i<15;i++{
		if i == 8 {
			fmt.Println("loop stopped:",i)
			break
		}
		fmt.Println(i)
	}

	//continue
	for i:=0;i<15;i++{
		if i == 6 {
			fmt.Println("skipped:",i)
			continue
		}
		fmt.Println(i)
	}
}