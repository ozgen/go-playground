package main

import "fmt"

type Triange struct {
	base   float64
	height float64
}

type Square struct {
	sideLength float64
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (t Triange) getArea() float64 {
	return 0.5 * t.base * t.height
}

type Shape interface {
	getArea() float64
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := Triange{base: 3, height: 4}
	s := Square{sideLength: 5}

	printArea(t)
	printArea(s)
}
