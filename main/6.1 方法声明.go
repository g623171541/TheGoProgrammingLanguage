package main

import (
	"math"
	"fmt"
)

/*
	方法声明：在函数声明时，在其名字之前放一个变量，就是一个方法。
 */

func main()  {
	p := Point6{1,2}
	q := Point6{4,6}
	fmt.Println( Distance(p,q) ) 	// function  实际上用的是包级别的函数 main.Distance
	fmt.Println( p.Distance(q) )	// method  使用刚刚声明的Point6，调用的是Point6类下声明的 Point6.Distance方法


	// 计算三角形的周长
	perim := Path{
		{1,1,},
		{5,1},
		{5,4},
		{1,1},
	}
	fmt.Println(perim.Distance())


}

type Point6 struct {
	x, y float64
}

func Distance(p,q Point6) float64 {
	return math.Hypot(q.x-p.x , q.y-p.y)
}



/*
	p：叫做方法的接收器，早期的面向对象语言留下的遗产将调用一个方法成为"向一个对象发送消息"。
	建议保持接收器在方法间传递时的一致性和简短性，可以使用其类型的第一个字母：p

	在方法调用过程中，接收器参数一般会在方法名之前出现。

	-->  相当于给某种类型（这里是Point6）添加一种方法。
*/
func (p Point6) Distance(q Point6) float64  {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

type Path []Point6
// 计算几个点之间的距离
func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p{
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}