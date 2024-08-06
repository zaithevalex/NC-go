package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	perimeter() float32
	square() float32
}

type Circle struct {
	radius float32
}

type Rectangle struct {
	height, width float32
}

func (circle Circle) perimeter() float32 {
	return math.Pi * circle.radius * 2
}

func (circle Circle) square() float32 {
	return circle.radius * circle.radius * math.Pi
}

func (rectangle Rectangle) perimeter() float32 {
	return 2*rectangle.height + 2*rectangle.width
}

func (rectangle Rectangle) square() float32 {
	return rectangle.height * rectangle.width
}

func result(geometry Geometry) {
	fmt.Println("square:", geometry.square())
	fmt.Println("perimeter:", geometry.perimeter())
}

func main() {
	rectangle1 := Rectangle{height: 10, width: 5}
	circle1 := Circle{radius: 10}
	fmt.Println("rectangle1:", rectangle1)
	fmt.Println("circle1:", circle1)

	result(circle1)
	result(rectangle1)

	//var geometry Geometry = rectangle1
	//fmt.Println("geometry:", geometry)
}
