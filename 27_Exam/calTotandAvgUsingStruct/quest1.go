package main

import (
	"fmt"
)

type Student struct {
	Name string
	Marks []int
	Total float64
	Average float64
}


func main(){
	fmt.Println("Student Marks Calculator.")
	 var n int=5
	// fmt.Println("Enter no of students")
	//fmt.Scanln(&n)
	Students := make([]Student,n)
		Students[0]= Student{"a",[]int{20,20,20,20,20},0,0}
		Students[1]= Student{"b",[]int{20,15,15,10,20},0,0}
		Students[2]= Student{"c",[]int{20,18,13,14,20},0,0}
		Students[3]= Student{"d",[]int{17,19,14,14,10},0,0}	
		Students[4]= Student{"e",[]int{20,19,19,18,18},0,0}	
	// for i:=0;i<5;i++{
		
	// }
	for i:=0;i<5;i++{
		for j:=0;j<len(Students[i].Marks);j++{
			Students[i].Total+=float64(Students[i].Marks[j])
		}
	}
	for i:=0;i<5;i++{
		Students[i].Average=float64(Students[i].Total)/float64(len(Students[i].Marks))
	}
	for i:=0;i<5;i++{
		fmt.Println("Student Name:",Students[i].Name)
		fmt.Println("Student Marks:",Students[i].Total)
		fmt.Println("Student Average:",Students[i].Average)
	}
}