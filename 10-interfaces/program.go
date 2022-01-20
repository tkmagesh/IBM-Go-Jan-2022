package main

import (
	"fmt"
	"math"
)

//Sprint-1
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

//Sprint-2
type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

type ShapeWithArea interface {
	Area() float64
}

func PrintArea(x ShapeWithArea) {
	fmt.Println("Area = ", x.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float64
}

func PrintPerimeter(x ShapeWithPerimeter) {
	fmt.Println("Perimeter = ", x.Perimeter())
}

func main() {
	c := Circle{Radius: 12}
	//fmt.Println("Area = ", c.Area())
	PrintArea(c)
	PrintPerimeter(c)

	r := Rectangle{Height: 12, Width: 10}
	//fmt.Println("Area = ", r.Area())
	PrintArea(r)
	PrintPerimeter(r)
}
