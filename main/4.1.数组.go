package main

import (
	"fmt"
	"crypto/sha256"
	"reflect"
)

/*
	数组：

 */

func main()  {
	func1()		//数组的创建 + 遍历
	func2()		//...创建数组，并指定数组编号
	func3()		//数组[...]int{3:-1} 的定义，数组的比较
	func4()
}



func func1()  {
	//默认情况下，数组的每个元素都被初始化为元素类型对应的零值。
	var a [3]int
	fmt.Println(a[0])
	fmt.Println( a[len(a) - 1] )

	//i 是元素的序号，v是数组下标对应的值
	for i, v := range a{
		fmt.Printf("第%d个 value=%d\n",i,v)
	}

	var r [3]int = [3]int{4,5}
	fmt.Println("",r[2])

	//在数组字面值中，如果在数组的长度位置出现 "..."省略号，则表示数组的长度是根据初始化值的个数来计算的。
	q := [...]int{1,2,3}
	fmt.Printf("%T \n",q)	//打印q的类型

	fmt.Println("----------func1到此结束----------")
}

func func2() {

	type currency int
	const (
		USD currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD:"$",EUR:"€",GBP:"£",RMB:"¥"}
	fmt.Println(symbol,symbol[RMB])

	fmt.Println("----------func2到此结束----------")
}

func func3() {
	r := [...]int{3:-1}		//定义了含有 3+1 个元素的数组，其中最后一个元素被初始化为-1，其他元素都用0初始化
	fmt.Println(r)

	//数组的比较
	a := [2]int{1,2}
	b := [...]int{1,2}
	c := [2]int{1,3}
	fmt.Println(a == b, a == c, b == c)

	//d := [3]int{1,2}
	//fmt.Println(a == d)	//报错：说类型不同，不能比较

	fmt.Println("----------func3到此结束----------")
}

func func4()  {
	// crypto/sha256 包的 Sum256 函数对一个任意字节slice类型的数据生成一个对应的消息摘要，消息摘要有256bit大小，因此对应 [32]byte数组类型。
	// 如果两个消息摘要相同，那么可以认为这两个消息本身相同（理论上HASH码碰撞的情况但实际应用可以基本忽略）

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n",c1)
	fmt.Println(c2)
	fmt.Println(c1 == c2)
	fmt.Println(reflect.TypeOf(c1))

	fmt.Println("----------func4到此结束----------")
}