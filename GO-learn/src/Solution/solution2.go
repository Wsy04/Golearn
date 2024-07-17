package main

import (
	"fmt"
	"math"
)

type myFloat64 float64

func (f myFloat64) Abs() float64 {
	if f < 0 {
		f = -f
	}
	return float64(f)
}

func (f *myFloat64) Abs2() float64 {
	if *f < 0 {
		*f = -*f
	}
	return float64(*f)
}
func main() {
	f := myFloat64(-math.Sqrt2)
	fmt.Println(f)
	fmt.Println(f.Abs()) //值传递
	fmt.Println(f)
	fmt.Println(f.Abs2()) //地址传递
	fmt.Println(f)
}
