package main

import (
	"fmt"
)

type shape interface {
	printArea() float64
	getIDShape() string
}

type triangle struct {
	shapeId string
	height  float64
	base    float64
}

type square struct {
	shapeId    string
	sideLenght float64
}

func main() {

	ta := triangle{

		shapeId: "triangle",
		height:  10,
		base:    5,
	}

	sa := square{

		shapeId:    "square",
		sideLenght: 10,
	}

	getArea(ta)
	getArea(sa)

}

func getArea(s shape) { // the function that takes the shape interface as an argument and prints the area with the id

	area := s.printArea()
	nameid := s.getIDShape()
	fmt.Printf("the area of %v the is: %0.2f\n", nameid, area)
}

func (ta triangle) printArea() float64 { //the function that prints the area of the triangle

	triangleArea := 0.5 * ta.base * ta.height
	return triangleArea
}

func (sa square) printArea() float64 { //the function that prints the area of the square

	squareArea := sa.sideLenght * sa.sideLenght
	return squareArea
}

func (s square) getIDShape() string { //the function that returns the id of the square
	return s.shapeId
}
func (t triangle) getIDShape() string { //the function that returns the id of the triangle
	return t.shapeId
}
