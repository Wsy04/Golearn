package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs1(V Vertex) float64 {
	return math.Sqrt(V.X*V.X + V.Y*V.Y)
}
func main() {
	v := Vertex{3, 4}
	v1 := Vertex{5, 12}
	fmt.Println(v.Abs())
	fmt.Println(Abs1(v1))
}
