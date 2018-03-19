package main

import "fmt"

/*
	------------------变量----------------------------------------------

	var 变量名字 类型 = 表达式

	可以省略 "类型" 或 "= 表达式"
		1. 省略 "类型" ： 将根据表达式来推到变量的类型信息
		2. 省略 "= 表达式" ： 将用零值初始化该变量。零值：0 false "" nil

	var i, k, k int
	var b, f, s = true, 2.3, "four"

	声明变量 ->   t := 0.2     i, j := 1, 2

	" := " 这是一个变量语句，而" = "是一个变量赋值操作

	//下面这段代码第一个语句声明了in和err两个变量，第二个语句只声明了out一个变量，然后对已经声明的err进行了赋值操作
	in, err := os.Open(infile)
	out, err := os.Close(outfile)
		简短变量声明语句中必须至少要声明一个新的变量，而不能两个都一样。



  	------------------指针----------------------------------------------
	任何类型的指针的零值都是nil
	var x, y int
	fmt,Println( &x == &x , &x == &y , &x == nil) // true false false




*/

func main() {

	// new函数 ------------------------------------------- 另一个创建变量的方法

	p := new(int)	//0xc420014048  ->  p是一个变量的地址
	fmt.Println(*p) //0 			->  *p才是值

	q := new(int)
	fmt.Println(q == p)		//false	->  new()每次都会得到新的变量的地址


	//  元组赋值  ----------------------------------
	x,y := 1,3
	x, y = y, x		//用来交换两个变量的值
	fmt.Println("x y分别为",x,y)

	fmt.Println("最大公约数为：", gcd(24,17) )
	fmt.Println("斐波那契数列第9个数为：", fib(9))

}

//计算两个数值的最大公约数（GCD）
func gcd(x, y int) int {
	for y != 0 {
		fmt.Println(x,y)
		x, y = y, x%y
		fmt.Println("---> ",x,y)
	}
	return x
}


//计算斐波那契数列 的第n个值
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x + y
	}
	return x
}