package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	defer fmt.Println(t)
	for i := 9; i > 0; i-- {
		defer fmt.Printf("%d", i)
	}

	switch {
	case t.Hour() < 12:
		fmt.Println("早上好")
	case t.Hour() < 14:
		fmt.Println("中午好")
	case t.Hour() < 17:
		fmt.Println("下午好")
	default:
		fmt.Println("晚上好")
	}
}
