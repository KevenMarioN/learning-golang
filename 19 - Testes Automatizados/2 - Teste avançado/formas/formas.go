package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

func WriteArea(f Forma) {
	fmt.Printf("A área da forma é %0.2f \n", f.Area())
}

type Rectangle struct {
	Altura  float64
	Largura float64
}

type Circle struct {
	Raio float64
}

func (r Rectangle) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Raio * c.Raio)
}
