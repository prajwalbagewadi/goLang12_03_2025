package main

import (
	"fmt"
)

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length float64
	width  float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main(){
	c := circle{radius: 5}
	r := rectangle{length: 4, width: 6}
	fmt.Println("Circle Area:", c.area())
	fmt.Println("Circle Perimeter:", c.perimeter())
	fmt.Println("Rectange Area:", r.area())
	fmt.Println("Rectangle Perimeter:", r.perimeter())
}