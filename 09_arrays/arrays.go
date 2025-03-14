// arrays in go
package main

import (
	"fmt"
)

func processbyval(arr [5]int){
	//array as parameter
	for ind,val := range arr{
		fmt.Print(ind,":",val," ")
	}
}

func processbyRef(arr *[5]int){
	//passing array pointer as parameter
	fmt.Println("original arr=",*arr)
	arr[1]=99
	fmt.Println("modified arr=",*arr)
}

func main(){
	/*
	syntax:
		var arr_name [size]type
		or
		var arr_name=[size]type{val1,val2...}
		or
		arr_name:=[size]type{val1,val2...}
	*/
	//simple array
	var arr1 [5]int
	for ind,_:=range arr1{
		fmt.Println("enter arr1 val:")
		fmt.Scan(&arr1[ind])
	}
	for ind,val:=range arr1{
		fmt.Print(ind,":",val," ")
	}
	fmt.Println()
	//Array Literal with Auto Size (...)
	// var arr2 =[...]int{1}
	// for ind,_:=range arr2{
	// 	fmt.Println("enter arr2 val:")
	// 	fmt.Scan(&arr2[ind])
	// }
	// for ind,val:=range arr2{
	// 	fmt.Print(ind,":",val," ")
	// }
	fruits := [...]string{"apple","cherry","banana","pineapple"}
	fmt.Println("fruits=",fruits)

	//access and modify array elements
	cart := [...]string{"bread","milk","eggs"}
	fmt.Println("original cart=",cart)
	cart[1]="butter"
	fmt.Println("modified cart=",cart)

	//multidimensional array
	/*
	syntax:
		var arr_name [rows][cols]type
		or
		arr_name:=[rows][cols]type{{v1,v2,...},{v1,v2,...}} 
	*/
	var matrix = [2][3]int{{1,2,3},{1,2,3}}
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
	for _,row := range matrix{  //ind and val as row(arr) := where range 2d array
		for _,val:= range row { //ind and val as single element := where range is the single row(array)
			fmt.Print(val)
		}
		fmt.Println()
	}

	//array as function para
	arr2:=[5]int{5,4,3,2,1}
	processbyval(arr2) //call
	fmt.Println("original arr2=",arr2)
	processbyRef(&arr2) //call
	fmt.Println("modified arr2=",arr2)
}