package main

import "fmt"

func main() {
	fmt.Println("Hello " + "world")
	var a int = 100
	var b int = 200
	var str = "hello=%d"
	fmt.Printf(str, a)
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Println(a)

}