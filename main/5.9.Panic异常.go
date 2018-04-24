package main

import (
	"fmt"
	"regexp"
	"runtime"
	"os"
)

/*
	panic
		不是所有的panic异常都来自运行时，直接调用内置的panic函数也会引发panic异常
		panic函数接受任何值作为参数。

		由于panic会引起程序的崩溃，因此panic一般用于严重错误，如程序内部的逻辑不一致。

		任何可以预料到的错误，如不正确的输入，错误的配置或是失败的I/O操作都应该被优雅的处理，最好的方式就是使用Go的错误机制。
			考虑regexp.Compile函数，该函数将正则表达式编译成有效的可匹配格式。当输入的正则表达式不合法时，该函数会返回一个错误。
			在程序源码中，大多数正则表达式是字符串字面值，因此regexp包提供了包装regexp.MustCompile检查输入的合法性。
 */

func main()  {
	//panic1()		//panic异常
	panic2()		//正则表达式

	defer printStack()	//在Go的panic机制中，延迟函数的调用在释放堆栈信息之前。
	f(3)
}

func panic1()  {
	switch s := "0"; s {
	case "s":
	case "h":
	case "c":
	default:
		panic(fmt.Sprintf("invalid suit %q",s))
	}
}

func panic2()  {
	//正则表达式
	// 函数名中 Must前缀是一种针对此类函数的命名约定
	var httpSchemeRE = regexp.MustCompile(`^https?:`)
	fmt.Println(httpSchemeRE)
}

func f(x int)  {
	fmt.Printf("f(%d)\n",x+0/x)
	defer fmt.Printf("defer %d\n",x)
	f(x - 1)
}

func printStack()  {
	var buf [4096]byte
	n := runtime.Stack(buf[:],false)
	os.Stdout.Write(buf[:n])
}