package main

import "fmt"

// 初始化数组时可以使用初始化列表来设置数组元素的值
func initArray1()  {
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
}

// 按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度
func initArray2()  {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}

// 可以使用指定索引值的方式来初始化数组
func initArray3()  {
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}


// 遍历数组
func indexArray()  {
	// 方法一：for...循环遍历
	var a = [...]string{"北京", "上海", "广州", "深圳"}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法二：for range遍历
	for index, value := range a {
		fmt.Println(index, "-", value)
	}
}


// 多维数组
func multiArray()  {
	// 二位数组的定义
	// 多维数组只有第一层可以使用...来让编译器推导数组长度
	a := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"西安", "宝鸡"},
	}
	fmt.Println(a)
	fmt.Println(a[2][0])
	fmt.Println("---------------")

	// 二维数组的遍历
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Println(v2, "\t")
		}
		fmt.Println()
	}
}


// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值
func modifyArray1(x [3]int)  {
	x[0] = 100
}
func modifyArray2(x [3][2]int)  {
	x[2][0] = 100
}
func copyArray()  {
	a := [3]int{10, 20, 30}
	modifyArray1(a)	 // 在modify中修改的是a的副本x
	fmt.Println(a)	// [10 20 30]

	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b)		// 在modify中修改的是b的副本x
	fmt.Println(b)		// [[1 1] [1 1] [1 1]]
}


//func main()  {
//	//initArray1()
//	//initArray2()
//	//initArray3()
//
//	//indexArray()
//
//	//multiArray()
//
//	copyArray()
//}