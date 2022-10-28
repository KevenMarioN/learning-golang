package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func writeArea(f forma) {
	fmt.Printf("A área da forma é %0.2f \n", f.area())
}

type rectangle struct {
	altura  float64
	largura float64
}

func (r rectangle) area() float64 {
	return r.altura * r.largura
}

func (c circle) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

type circle struct {
	raio float64
}

func main() {
	r := rectangle{10, 15}
	writeArea(r)

	c := circle{14.5}
	writeArea(c)
}
