// interfaces in go also refer to methods in go
package main

import (
	"fmt"
)

type Speaker interface{
	Speak()
}

type Walker interface{
	Walk()
}

type Human interface{
	//embedded interface
	Walker
	Speaker
}

type Person struct{
	Name string
}

func (p Person) Speak(){
	fmt.Println(p.Name,"Says hello.")
}

func (p Person) Walk(){
	fmt.Println(p.Name,"is Walking.")
}


func main(){
	fmt.Println("interfaces")
	/*
	What is an Interface?
	An interface in Go is a collection of method signatures that a type must implement. It allows for polymorphism and promotes code flexibility.
	*/
	var s Speaker
	p1 := Person{"Alice"}
	s = p1
	s.Speak()

	//call for embedded interface
    var h Human 
	h = p1
	h.Speak()
	h.Walk()

	//Empty interface
	var eiface interface{}
	eiface = 25
	fmt.Printf("type eiface=",eiface)

	eiface = "hello world."
	fmt.Printf("type eiface=",eiface)
}