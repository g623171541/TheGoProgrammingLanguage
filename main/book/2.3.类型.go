package main

import "fmt"

/*
	一个类型声明语句创建了一个新的类型名称，和现有的类型具有相同的底层结构
	新命名的类型提供了一个方法，用来分割不同概念的类型，这样即使他们底层类型相同也不是兼容的
	eg:	type 类型名字 底层类型
 */
 
type Celsius float64		//摄氏温度
type Farenheit float64		//华氏温度

const (
	AbsouluteZeroC	Celsius = -273.15		//绝对零度
	FreezingC		Celsius = 0				//结冰点温度
	BoilingC		Celsius = 100			//沸点温度
)

func main()  {
	//温度相互转换
	fmt.Println( CToF(100) )

	var c Celsius = 100
	fmt.Println(c.String())

}

func CToF(c Celsius) Farenheit {
	return Farenheit( c * 9 / 5 + 32 )
	//类型转换不会改变值本身，但是会使他们的语义发生变化
}

func FToC(f Farenheit) Celsius {
	return Celsius( (f - 32) * 5 / 9 )
}


//
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C",c)
}

func init()  {

}