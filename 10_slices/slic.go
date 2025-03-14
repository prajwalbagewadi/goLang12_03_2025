// 3.4 Slices and Slice Parameters
package main

import "fmt"

func modSlice(s []int){
	s[0]=100
}

func main(){
	//Slicing an Array or Another Slice
	arr := [8]int{1,2,3,4,4,3,2,1}
	slice := arr[1:4] //slice from array // Slicing elements from index 1 to 3 (excludes 4)
	fmt.Println("slice from arr=",slice)

	var slice1 []int
	fmt.Println("empty slice=",slice1)

	slice2 := []int{10,20,30,40} //slice with initilized values
	fmt.Println("initilzed slice=",slice2)

	slice3 := make([]int,3,5) //make([]type,len,capacity)
	fmt.Println("using make method=",slice3," len=",len(slice3)," cap=",cap(slice3))
	slice3=append(slice3,10,20,20)
	fmt.Println("using make method with append=",slice3," len=",len(slice3)," cap=",cap(slice3))

	//slice parameter
	slice4 := []int{1,2,3,4}
	fmt.Println("orginal slice=",slice4)
	modSlice(slice4)
	fmt.Println("modified slice=",slice4)


	//multidimensional slice 
	m1s := [][]int {
		{1,2,3},
		{4,5,6},
	}
	fmt.Println("m1s=",m1s)

	//dynamic
	rows :=3 
	cols :=3
	m2s := make([][]int,rows)
	for i:=0;i<rows;i++{
		m2s[i]=make([]int,cols)
	}
	fmt.Println("m2s=",m2s)
}
