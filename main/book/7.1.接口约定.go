package main

import "fmt"	// 讲的太深了，看不懂

/*
	接口类型是对其他类型行为的抽象和概况，因为接口类型不会和特定的实现细节绑定在一起，通过这种抽象的方式我们可以让我们的函数更加灵活和更加具有适应能力。

	Go语言中接口类型的独特之处在于他是满足隐式实现的：也就是说我们没有必要对于给定的具体类型定义所有满足的接口类型。简单的拥有一些必需的方法就足够了。

	接口类型：是一种抽象的类型。它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；
			它们志辉展示出他们自己的方法。
			也就是说当你有看到一个接口类型的值时，你不知道他是什么，唯一知道的就是可以通过它的方法来做什么。

		具体描述了一系列方法的集合，一个
 */

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int,error) {
	fmt.Println(*c)
	*c += ByteCounter(len(p))
	fmt.Println(*c)
	return len(p),nil
}

func main()  {
	var a ByteCounter = 1
	count,_ := a.Write([]byte("12345"))
	fmt.Println("count:",count)
	fmt.Println("a:",a)
}

//-------------------------------------接口内嵌

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface{
	Close() error
}

type ReaderWriter interface  {
	Reader
	Closer
}


