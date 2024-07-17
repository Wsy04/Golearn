package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	var m = make(map[string]int)
	strs := strings.Fields(s)
	for i := 0; i < len(strs); i++ {
		m[strs[i]]++
	}
	return m
}
func main() {
	var m map[string]int
	m = make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println(m)
	v, ok := m["b"]
	fmt.Println(v, ok)
	delete(m, "b")
	v, ok = m["b"]
	fmt.Println(v, ok)
	m2 := wordCount("I am learning Go!")
	fmt.Println(m2)
}
