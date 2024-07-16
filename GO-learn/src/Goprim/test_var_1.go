package main

import "fmt"

// 全局变量的声明
var gA int = 100
var gB = 200

// 变量的声明
func main() {
	//1.有类型没赋值,值为默认值
	var a int
	fmt.Println("a =", a)
	//2.同时设置类型和值
	var b int = 100
	fmt.Println("b =", b)
	//只赋值，自动识别类型
	var c = 100
	var cc = "hello"
	fmt.Println("c =", c)
	fmt.Printf("type of c is %T\n", c)
	fmt.Printf("type of cc is %T\n", cc)
	//直接声明,只可以用于函数体内,不可以用于全局变量
	d := "114514"
	fmt.Println("d =", d)
	fmt.Printf("type of d is %T\n", d)

	fmt.Printf("gA = %d\n", gA)
	fmt.Printf("gB = %d", gB)

	var xx, yy = 114514, 1919810
	fmt.Printf("xx = %d,yy = %d\n", xx, yy)
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, " jj = ", jj)
}
