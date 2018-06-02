package main

import (
	"fmt"
	"time"
)

/*
	一个对象的变量或者方法如果对调用方是不可见的话，一般就被定义为"封装"。

	Go语言只有一种口直可见性的手段：大写首字母的标识符会从定义它们的包中被导出，小写字母则不会。
	要想封装一个对象，就必须将其定义为一个struct

 */

func main()  {

	funcCounter()
	funcTime()
}

func funcTime()  {
	const day = 24 * time.Hour
	fmt.Println("一天有多少秒：",day.Seconds())
}

// 为结构体封装了几个方法
func funcCounter()  {
	var c = Counter{n:1}
	fmt.Println(c.N())

	c.Reset()
	fmt.Println(c.n)
}

type Counter struct {
	n	int
}
func (c *Counter) N() int {
	return  c.n
}
func (c *Counter) Increment() {
	c.n++
}
func (c *Counter) Reset()  {
	c.n = 0
}

