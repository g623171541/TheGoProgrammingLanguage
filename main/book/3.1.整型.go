package main

import "fmt"

/*
	整形：	①有符号  int8  int16  int32  int64
			②无符号  uint8 uint16 uint32 uint64

	在Go语言中，% 取模运算符的符号和被取模数的符号是一致的，如：
		-5%3  == -2
		-5%-3 == -2

	位运算：
		&	位运算 AND
		|	位运算 OR
		^	位运算 XOR			二元运算符：按位异或；一元运算符：按位取反
		&^	位清空（AND NOT）		按位置零
		<<	左移
		>> 	右移

 */

func main()  {
	//按位运算
	bit()
}

//按位运算
func bit()  {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("x    = %08b\n",x)
	fmt.Printf("y    = %08b\n",y)

	fmt.Printf("x&y  = %08b\n",x & y)		//与：相同位的两个数字都为1，则为1；若有一个不为1，则为0
	fmt.Printf("x|y  = %08b\n",x | y)		//或：有一个是1则为1，
	fmt.Printf("x^y  = %08b\n",x ^ y)		//异或：相同位不同则为1，相同则为0
	fmt.Printf("x&^y = %08b\n",x &^ y)		//位清空：如果y bit位上的数是0则取x上对应位置的值， 如果y bit位上为1则结果位上取0

}