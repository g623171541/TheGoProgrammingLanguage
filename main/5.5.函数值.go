package main

import "fmt"

/*
	1.在Go中，函数也像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。
	2.函数类型的零值是nil，调用值为nil的函数值会引起panic错误.
	3.函数值可以与 nil 比较

 */

func main()  {
	function1()
}

// 1.在Go中，函数也像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。
func function1()  {
	f := square
	fmt.Println(f(3))

	//f = product		//类型错误，can't assign func(int,int) int to func(int) int

	//3.函数值可以与 nil 比较
	fmt.Println( f == nil )
}










func square(n int) int {
	return n * n
}
func product(m,n int) int {
	return m * n
}