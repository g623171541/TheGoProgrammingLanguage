package main

import (
	"fmt"
	"os"
	"log"
)

/*
	错误：
		value, ok := func()		//通常被命名为 ok
		if !ok {
			// do something, 处理错误
		}

		打印error, error 可能是 nil 或者 non-nil
			fmt.Println(err)
			fmt.Printf("%v",err)

		当读取文件发生错误时，Read函数会返回可以读取的字节数以及错误信息，对于这样的情况，正确的处理方式应该是先处理这些不完整的数据，再处理错误。

		Go使用控制流机制（如if和return）处理异常，这使得编码人员能更多的关注错误处理。

	***
		在Go中，错误处理有一套独特的编码风格，检查某个子函数是否失败后，我们通常将处理失败的逻辑代码放在处理成功的代码之前。
		如果某个错误会导致函数返回，那么成功时的逻辑代码不应该放在else语句块中，而应直接放在函数体中。

 */

func main()  {

	// 如果发生错误程序无法继续运行，就要输出错误信息并结束程序。此操作只应在main中执行。
	if err := new(error); err != nil {
		fmt.Fprintf(os.Stderr,"site is down: %v\n",err)
		os.Exit(1)
	}
}

func err1()  {
	err := new(error)
	fmt.Errorf("error is %v",err)
	// fmt.Errorf 函数使用 fmt.Sprintf 格式化错误信息并返回。
}

func logMsg()  {
	// 只需要输出错误信息就足够了，不需要中断程序的运行，通过log包提供的函数
	if err := new(error); err != nil {
		log.Printf("faild : %v",err)
		fmt.Fprintf(os.Stderr,"faild : %v",err)
	}

	// log包中的所有函数会为没有换行符的字符串增加换行符。
}