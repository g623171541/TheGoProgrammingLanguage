package main

import (
	"time"
	"fmt"
)

/*
	goroutines：go语言中每一个并发的执行单元。可以类比作一个线程。

		当一个程序启动时，其主函数即在一个单独的 Goroutine 中运行，称为 main goroutine
		新的Goroutine 会用go语句来创建。 如
*/

func fa()  {

}

func main()  { 		// main goroutine 主线程

	fa()		// call f(); wait for it to return
	// 新创建的线程 执行 fa()
	go fa()		// cerate a new goroutine that calls f(); don't wait

//----------------------------------计算斐波那契数列的第45个元素值

	go spinner(300 * time.Millisecond)
	const n = 45
	fibN := fibonacci(n)
	fmt.Printf("\rFibonacci(%d)",n,fibN)

}

// 转圈的等待动画
func spinner( delay time.Duration )  {
	for  {
		for _,r := range `-\|/` {
			fmt.Printf("\r%c",r)
			time.Sleep(delay)
		}
	}
}

func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}