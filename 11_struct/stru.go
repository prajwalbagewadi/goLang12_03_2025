package main

import (
	"fmt"
)

type Employee struct{
	ename string
	desg string
}

type Person struct{
	name string
	age int
}

func updateAge1(p Person){
	//pass by value (struct as a para)
	p.age=30
}

func updateAge2(p *Person){
	//pass by Reference (struct as a para)
	p.age=30
}

func main(){
	var e1 = Employee{ename:"prajwal",desg:"SDE1"}
	fmt.Println("E1=",e1)

	var e2 Employee
	fmt.Println("Enter ename=")
	fmt.Scan(&e2.ename)
	fmt.Println("Enter desg=")
	fmt.Scan(&e2.desg)
	fmt.Println("Ename=",e2.ename," desg=",e2.desg)
	var temp string
	fmt.Scan(&temp)

	var p1 = Person{name:"abc",age:25}
	fmt.Println("P1=",p1)

	var p2 Person
	fmt.Println("enter name and age:")
	fmt.Scan(&p2.name)
	fmt.Scan(&p2.age)
	fmt.Println(" name:",p2.name," age:",p2.age)

	// Passing a Structure as para
	p3 := Person{"alice",25}
	updateAge1(p3)
	fmt.Println("call by val p3=",p3)
	updateAge2(&p3)
	fmt.Println("call by ref p3=",p3)
}