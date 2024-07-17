package main

import (
	"math"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		temp := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			temp[j] = uint8(float64(i) * math.Log(float64(j)))
		}
		pic[i] = temp
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
