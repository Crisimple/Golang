package main

import "fmt"

const (
	a, b= iota + 1, iota + 2 // 1, 2
	c, d                   // 2, 3
	e, f                   // 3, 4
)


func main_test()  {
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)
}