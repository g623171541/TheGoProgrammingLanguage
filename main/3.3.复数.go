package main

import "fmt"

/*
	两种精度的复数类型：complex64 和 complex128
				分别对应float32 和 float64 两种浮点数精度
 */

func main()  {
	//x := 1 + 2i		//简写
	var x complex128 = complex(1,2)

	//y := 3 + 4i
	var y complex128 = complex(3,4)
	fmt.Println( x * y )
	fmt.Println(real(x*y))		//real()函数返回复数的实部
	fmt.Println(imag(x*y))		//imag()函数返回复数的虚部
}