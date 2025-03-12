/*
Datatypes
Basic type: Numbers, strings, and booleans come under this category.
Aggregate type: Array and structs come under this category.
Reference type: Pointers, slices, maps, functions, and channels come under this category.
Interface type
*/

/*
Basic Datatypes:
Numbers
Booleans
Strings

numbers:
Data Type   Description

int8	8-bit signed integer
int16	16-bit signed integer
int32	32-bit signed integer
int64	64-bit signed integer
uint8	8-bit unsigned integer
uint16	16-bit unsigned integer
uint32	32-bit unsigned integer
uint64	64-bit unsigned integer
int		Both int and uint contain same size, either 32 or 64 bit.
uint	Both int and uint contain same size, either 32 or 64 bit.
rune	It is a synonym of int32 and also represent Unicode code points.
byte	It is a synonym of uint8.
uintptr	It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.


Floating-Point Numbers: In Go language, floating-point numbers are divided into two categories as shown in the below table.
Possible arithmetic operations : Addition, subtraction, multiplication, division.
Three literal styles are available :
decimal (3.15)
exponential ( 12e18 or 3E10)
mixed (13.16e12)
Data Type	Description
float32	32-bit IEEE 754 floating-point number
float64	64-bit IEEE 754 floating-point number


Complex Numbers: The complex numbers are divided into two parts are shown in the below table. float32 and float64 are also part of these complex numbers. The in-built function creates a complex number from its imaginary and real part and in-built imaginary and real function extract those parts.
There are few built-in functions in complex numbers:
complex – make complex numbers from two floats.
real() – get real part of the input complex number as a float number.
imag() – get imaginary of the input complex number part as float number
Data Type	Description
complex64	Complex numbers which contain float32 as a real and imaginary component.
complex128	Complex numbers which contain float64 as a real and imaginary component.

Booleans
The boolean data type represents only one bit of information either true or false.
The values of type boolean are not converted implicitly or explicitly to any other type.


Strings
The string data type represents a sequence of Unicode code points. Or in other words, we can say a string is a sequence of immutable bytes, means once a string is created you cannot change that string. A string may contain arbitrary data, including bytes with zero value in the human-readable form. Strings can be concatenated using plus(+) operator.
*/

package main

import (
	"fmt"
)

func main(){

/*
A variable in Go is a named storage location that holds a value of a specific type. 
Variables can be declared in different ways depending on their scope and usage.
*/

//Using var (Explicit Declaration)
var name string = "prajwal"
fmt.Println("name=",name)

var x int = 10 
fmt.Print("\nx=",x)

//Using := (Short Declaration - Implicit Type)
lang := "go lang"
fmt.Print("\nlang=",lang,"\n")

//multiple var Declaration
var a,b,c int = 10,20,30
fmt.Println("a=",a)
fmt.Println("b=",b)
fmt.Println("c=",c)

//block declaration
var (
	sname string = "praj"
	sage int = 26
)
fmt.Printf("sname= %s,sage=%d\n",sname,sage)

//int 
var v1 int = 10
fmt.Println("int=",v1)

//float
var v2 float32 = 23.12345
fmt.Println("float32=",v2)

//complex 
var v3 complex64 = 13 + 33i
fmt.Println("v3=",v3)
fmt.Println("v3 real=",real(v3))
fmt.Println("v3 imag=",imag(v3))

//boolean
var b1 bool = true
fmt.Println("bool=",b1)

//string
s1 := "hello"
var s2 string = "world"
fmt.Println("concat s1 and s2=",s1+s2)

//constant
const PI float32 = 3.142
fmt.Println("PI=",PI)
}


