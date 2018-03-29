package main

import (
	"fmt"
	"time"
	"reflect"
)

/*

 */

type Weekday int
const (

	//iota常量生成器：用于生成一组相似规则的初始化常量，在第一个声明的常量所在行 iota将会被置位0，然后在每一个有常量声明的行加一
	//iota 翻译为 极少量

	Sunday Weekday = iota	//0
	Monday					//1
	Tuesday					//2
	Wednesday				//3
	Thursday				//4
	Friday					//5
	Saturday				//6
)

type Flags uint
const (
	FlagUp Flags = 1 << iota		//1
	FlagBroadcast					//2
	FlagLoopback					//4
	FlagPointToPoint				//8
	FlagMulticast					//16
)


const (
	_ = 1 << (10 * iota)
	KiB	//1024
	MiB	//1048576
	GiB	//1073741824
	TiB
	PiB
	EiB	//1152921504606846976
	ZiB
	YiB
)


func main()  {
	const noDelay time.Duration = 0					//time.Duration 是一个命名类型，底层类型是int64
	const timeout = 5 * time.Minute

	fmt.Println(reflect.TypeOf(noDelay))			//打印还是time.Duration类型
	// %T 打印变量的类型
	fmt.Printf("%T %[1]v\n",noDelay)			//time.Duration 0s
	fmt.Printf("%T %[1]v\n",timeout)			//time.Duration 5m0s
	fmt.Printf("%T %[1]v\n",time.Minute)		//time.Duration 1m0s

	fmt.Println(KiB)
	fmt.Println(MiB)
	fmt.Println(GiB)

	// 无类型常量------------------------
	// 6种：无类型的布尔值、无类型的证书、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串
	// 编译器为这些没有明确的基础类型的数字常量提供比基础类型更高精度的算术运算，至少有256bit的运算精度

	fmt.Println(YiB / ZiB) 	//1024，虽然YiB的值已经超出了任何GO语言中整数类型能表达的范围。但是依旧是合法的常量。

	var f float64 = 212
	fmt.Println( (f - 32) * 5 / 9 )		//100	(f - 32) * 5 is a float64
	fmt.Println( 5 / 9 * (f - 32) )		//0 	5/9 is an untyped integer , 0
	fmt.Println( 5.0 / 9.0 * (f - 32) )	//100	5.0 / 9.0 is an untyped float
	//只有常量可以是无类型的

	i := 0			// untyped integer;			implicit int(0)
	r := '\000'		// untyped rune;			implicit rune('\000')
	fl := 0.0		//untyped floating-point;	implicit float64(0.0)
	c := 0i			//untyped complex;			implicit complex128(0i)
	fmt.Println(i, r, fl, c )

}
