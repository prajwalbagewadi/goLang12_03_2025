// methods in go
package main

import (
	"fmt"
)

type Emp struct{
	eid int
	ename string
	edesg string
	esal float64
}

type Emp2 struct{
	ename string
	sal float64
}

//method
func (e Emp) printEmp(){
	fmt.Println("Emp Details:")
	fmt.Println("id=",e.eid,"name=",e.ename,"edesg=",e.edesg,"esal=",e.esal)
}

//value as Receiver (pass by val)
func (e Emp2) updateSal(newSal float64){
	e.sal=newSal
}

//Pointer as Receiver (pass by ref)
 func (e *Emp2) updateSal2(newSal float64){
	e.sal=newSal
}


func indentifyType(i interface{}){
	//type switch
	switch v := i.(type){
	case string :
		fmt.Println("this is a string",v)
	case int :
		fmt.Println("this is a integer",v)
	case bool :
		fmt.Println("this is a bool",v)
	default :
		fmt.Println("unknown type")	
	}
}

//interface 
type Speaker interface{
	Speak() string
}

//struct
type Person struct {
	Name string
}

//interface method overriden
func (p Person) Speak() string {
	return "hello im "+p.Name
}

func main(){
	fmt.Println("methods.")
	var e1 = Emp{1,"abc","sde1",50000.00}
	fmt.Println(e1)
	e1.printEmp()

	var e2 = Emp2{"alice",50000}
	fmt.Println("e2=",e2)
	e2.updateSal(60000) //value as Receiver (pass by val)
	fmt.Println("e2=",e2)

	fmt.Println("e2=",e2)
	e2.updateSal2(60000) //pointer as Receiver (pass by ref)
	fmt.Println("e2=",e2)

	//interface types 
	//Type Assertion
	var i interface{} = "helloworld"
	str,ok := i.(string)

	if ok {
		fmt.Println("this is a string:",str)
	} else {
		fmt.Println("not a string.")
	}

	var i2 interface{} = 4
	val,ok := i2.(string)
	if ok {
		fmt.Println("this is a string:",val)
	} else {
		fmt.Println("not a string")
	}

	//type switch call
	indentifyType(20)
	indentifyType("go")
	indentifyType(true)

	//interface values
	var s Speaker //interface var
	s = Person{"alice"} //assigning value to interface

	fmt.Printf("Dynamic Type:%T\n",s)//prints concrete type
	fmt.Printf("Dynamic value:",s.Speak())//func call
}