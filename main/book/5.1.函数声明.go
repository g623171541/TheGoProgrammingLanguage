package main

import "fmt"

/*
	函数的声明：
		包括函数名、形参列表、返回值列表（可省略）、以及函数体
		func name(parameter-list) (result-list) {
			body
		}
 */

func main()  {
	// 打印类型
	fmt.Printf("%T\n",add)	// func(int, int) int
	fmt.Printf("%T\n",sub)	// func(int, int) int
	fmt.Printf("%T\n",first)	// func(int, int) int
	fmt.Printf("%T\n",zero)	// func(int, int) int

	// 函数的形参作为局部变量，被初始化为调用者提供的值
	fmt.Println(add(2,1))
	// 实参通过值的方式传递，因此函数的形参是实参的拷贝。
	// 对形参进行修改不影响实参，但是如果实参包括引用类型，如指针 slice map function channel等类型，实参可能会由于实参的间接引用被修改。


	// 如果遇到没有函数体的函数声明，表示该函数不是Go实现的，这样的声明定义了函数标识符：
	// 如:
	// package math
	// func Sin(x float64) float


}

func add(x int, y int) int {
	return x + y
}

func sub(x,y int) (z int) {
	z = x - y
	return
}

func first(x int, _ int) int {		// 可以强调某个参数未被使用
	return x
}

func zero(int, int) int {
	return 0
}