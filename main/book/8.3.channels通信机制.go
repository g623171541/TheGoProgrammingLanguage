package main

import (
	"fmt"
	"reflect"
	"net"
	"log"
	"io"
	"os"
)

/*
	Channels：是goroutine(线程) 之间的通信机制，一个channels是一个通信机制。
	它可以让一个goroutine通过它给另一个goroutine发送值信息。每个channel都有一个特殊的类型，也就是channels可发送数据的类型。


 */

func main()  {
	// 一个可以发送int类型数据的channel一般写为chan int
	ch := make(chan int,1)	// 后面的 1 代表channel的容量，如果容量大于零，则channel就是带缓冲的channel

	// channel的零值为nil，可以用==进行比较

	/*
		一个channel有发送和接受两个主要的操作，都是通信行为。
		一个发送语句将一个值从一个 goroutine通过channel发送到另一个执行接受操作的goroutine。
		发送和接收两个操作都是用 <- 运算符。
			在发送语句中，<- 运算符分割channel和要发送的值
			在接收语句中，<- 运算符写在channel对象之前。一个不使用接收结果的接收操作也是合法的。
	*/

	ch <- 10		// 写入
	v := <- ch		// 读取
	fmt.Println(v)
	fmt.Println(reflect.TypeOf(ch))

	/* 	channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic异常。
	*/
	close(ch)



	// 不带缓存的Channels
	channel1()

}

// 不带缓存的Channels
func channel1()  {
	/* 一个基于无缓存channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的channels上执行接收操作，
	当发送的值通过channels成功传输后，两个goroutine可以继续执行后面的语句。
		反之，如果接收操作先发生，那么接收者goroutine也将阻塞，直到有另一个goroutine在相同的channels上执行发送操作。

	基于无缓存channels的发送和接收操作将导致两个goroutine 做一次同步操作。因为这个原因，无缓存channels有时候也被称为同步channels。

	当通过一个无缓存channel发送数据时，接受者收到数据发生在唤醒发送者goroutine之前（happens before，这是go语言并发内存模型的一个关键术语。）

	 */

	 // 在主goroutine中，将标准输入复制到server，因此当客户端程序关闭标准输入时，后台goroutine可能依然在工作。
	 // 我们需要让主goroutine等待后台goroutine 等待后台goroutine完成工作后再退出，我们使用了一个channel来同步两个goroutine
	conn,err := net.Dial("tcp","localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout,conn)
	}()

	mustCopy(conn,os.Stdin)
	conn.Close()
	<- done

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}