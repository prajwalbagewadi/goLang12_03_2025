package main

import (
	"fmt"
)

func main(){
	var name string
	var age int
	// fmt.Scan : Used for simple inputs like integers, strings, or floats.
	fmt.Println("Enter name:")
	fmt.Scan(&name)
	fmt.Println("Enter age:")
	fmt.Scan(&age)
	fmt.Println("name=",name)
	fmt.Println("age=",age)

	// fmt.Scanf : Takes formatted input, similar to scanf in C.
	var temp string
	fmt.Scanf("%s",&temp)
	fmt.Println("Enter name and age:")
	fmt.Scanf("%s %d",&name,&age)
	fmt.Println("name=",name)
	fmt.Println("age=",age)
}