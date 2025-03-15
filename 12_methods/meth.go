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

}