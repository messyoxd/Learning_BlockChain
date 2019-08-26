package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	base   float64
	altura float64
}

func (rect Rectangle) area() float64 {
	return rect.base * rect.altura
}

type Circle struct {
	radius float64
}

func (circ Circle) area() float64 {
	return math.Pi * math.Pow(circ.radius, 2)
}

type Shape interface {
	area() float64
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	rect1 := Rectangle{10, 10}
	fmt.Printf("A area do retangulo e %f\n", getArea(rect1))
	circle := Circle{10}
	fmt.Printf("A area do circulo e %f\n", getArea(circle))
}
