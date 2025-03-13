package main

import (
	"fmt"
)

//            parameters  return-type
func addition(x int,y int)int{
	//basic function with parameters and return type
	return x+y
}

func changeVal1(x int){
	//call by value passes copy of a variable 
	x=1
}

func changeVal2(ptr *int){
	//call by reference passes address of a variable 
	*ptr=2
}

func addAndMulti(x int,y int)(sum int,product int){
	//Named Return Variables
	sum=x+y //assign values directly to named variables
	product=x*y
	return //no need to explictly mention sum or product
}

func variadic(args ...int) int {
	total := 0
	for _,val := range args{
		total+=val
	}
	return total
}


func deferfunc(){
	//defer function (works for inbuilt functions)
	var x,y int = 10,20
	defer fmt.Println("defer func result=",x+y)
 	// fmt.Println("enter x and y:")
	// fmt.Scanf("%d %d",&x,&y)
	fmt.Println("x=",x)
	fmt.Println("y=",y)
}

func factorial(n int)int{
	//recursive function
	if n==0{ //base condition
		return 1
	}
	return n * factorial(n-1) //recursive call
}

func main(){
	fmt.Println("functions.")
	fmt.Println("keyword: func used.")

	result := addition(5,6)
	fmt.Println("addition call=",result)	

	//call by value passes copy of a variable 
	v1 := 20
	changeVal1(v1)
	fmt.Println("call by val=",v1) //val doesn't change

	//call by reference passes address of a variable 
	v2 := 20
	changeVal2(&v2)
	fmt.Println("call by Ref=",v2) //val changes

	//Named Return Variables
	/*
	How Named Return Variables Work (Simplified)
		1)Declare return variables with types in the function signature.
		2)Assign values directly to these return variables inside the function.
		3)Use return without specifying values—Go returns them automatically.
	*/
	s,p := addAndMulti(5,5)
	fmt.Println("sum=",s)
	fmt.Println("multiply=",p)

	//Variadic (Variable Argument) Functions in Go
	/*
		Syntax :
			func fun_name(param ...type) returntype {
				//body
			}
			...type (three dots ... before type)means the function can accept multiple args
	*/
	fmt.Println("variadic func call 1:",variadic(1,2,3))
	fmt.Println("variadic func call 2:",variadic(1,2,3,4,5))

	/*
	Defer Statement in Go
	The defer statement in Go delays the execution of a function until the surrounding function exits. 
	It is commonly used for resource cleanup (closing files, releasing locks, etc.).
	*/

	//simple defer statement 
	defer fmt.Println("hello")
	fmt.Println("world")

	//defer function
	deferfunc() //defer func call

	//Multiple defer Calls (LIFO Order)
	defer fmt.Println("stmt 1:",1)
	defer fmt.Println("stmt 2:",2)
	defer fmt.Println("stmt 3:",3)

	//recursive function:
	/*
	A recursive function is a function that calls itself to solve a problem. 
	It breaks the problem into smaller subproblems until a base condition is met.
	*/
	/*
	syntax:
	func fun_name(para type)return-type{
		if baseCondition{
			return result
		}
		return fun_name(modified parameter)//recursive function call
	}
	*/
	fmt.Println("factorial eg:",factorial(5))
}