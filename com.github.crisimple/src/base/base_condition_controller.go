package main

import "fmt"

func conditionIfElse()  {
	score := 90
	if score >=90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// if条件判断还有一种特殊的写法，在 if 表达式之前添加一个执行语句，再根据变量值进行判断
	if age := 25; age > 20{
		fmt.Println("A1")
	} else if age > 18{
		fmt.Println("B2")
	} else {
		fmt.Println("C3")
	}

}

func conditionFor() {
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}

	j := 10
	for j > 0 {
		fmt.Println(j)
		j--
	}

	y := 0
	for ; y < 10; y++{
		fmt.Println(y)
	}
}


func main() {
	//conditionIfElse()
	conditionFor()
}