package main

import (
	"fmt"
	"math"
)

type square struct {
	L float64
}
type circle struct {
	R int
}
type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(float64(c.R), 2)
}
func (s square) area() float64 {
	return s.L * s.L
}

func declarations() {
	sq := square{
		L: 10,
	}
	circ := circle{
		R: 50,
	}

	info(sq)
	info(circ)
}

func info(x shape) {
	fmt.Println(x.area())

}
