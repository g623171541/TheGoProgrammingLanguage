package main

/*
	Channels：是goroutine 之间的通信机制，一个channels是一个通信机制。
	它可以让一个goroutine通过它给另一个goroutine发送值信息。每个channel都有一个特殊的类型，也就是channels可发送数据的类型。


 */

func main()  {
	// 一个可以发送int类型数据的channel一般谢伟chan int
	ch := make(chan int)
}