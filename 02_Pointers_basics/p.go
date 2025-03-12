package main

import (
	"fmt"
)

func passbyval(val int){
	val=50
}

func passbyref(ptr *int){
	*ptr=45
}

func main(){
	fmt.Println("Pointers")
	//declaration of ptr
	var ptr *int

	//assigning of ptr
	x := 42
	ptr = &x
	fmt.Println("val in ptr=",ptr)
	fmt.Println("val stored at address in ptr=",*ptr)
	/*
		& (Address Operator) → Gets the memory address.
		* (Dereference Operator) → Accesses the value stored at the address.
	*/

	//modify val using ptr
	fmt.Println("x before mod=",x)
	*ptr=32
	fmt.Println("x after mod=",x)


	//passbyval
	fmt.Println("x before passbyval=",x)
	passbyval(x)
	fmt.Println("x after passbyval=",x) //no change

	//passbyref
	fmt.Println("x before passbyref=",x)
	passbyref(&x) //&var
	fmt.Println("x after passbyref=",x) 


}