package main

import (
	"fmt"
)

func main(){
	for i:=0;i<5;i++{
		fmt.Println("ilteration=",i)
	}

	//ilterating over arrays
	arr := []int{10,20,30,40,50}
	/*
	In Golang Range keyword is used in different kinds of data structures in order to iterates over elements. 
	The range keyword is mainly used in for loops in order to iterate over all the elements of a map, slice, channel, or an array. 
	*/
	for ind,val := range arr {
		fmt.Println("ind=",ind," val=",val)
	}

	//ilterating over Slice
	slice := arr[1:4]
	for ind,val := range slice {
		fmt.Println("ind=",ind," val=",val)
	}

	//Iterating Over Maps
	MAP := map[string]int{"orange":10,"cherry":20,"pineapple":5}
	for key,val := range MAP{
		fmt.Println("Key=",key," val=",val)
	}
}