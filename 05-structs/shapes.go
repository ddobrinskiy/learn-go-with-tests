package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (tr Triangle) Area() float64 {
	return tr.Base * tr.Height * 0.5
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return rectangle.Height*2 + rectangle.Width*2
}
