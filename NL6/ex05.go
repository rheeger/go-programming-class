package main

import "fmt"
import "math"

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (a circle) area() float64 {
	return a.radius * a.radius * math.Pi
}

func (a square) area() float64 {
	return a.length * a.length
}

type shape interface {
	area() float64
}

func info(x shape) {
	fmt.Println(x.area())

}

func main() {
	sq := square{
		length: 23,
	}

	cir := circle{
		radius: 23,
	}

	info(sq)
	info(cir)
}
