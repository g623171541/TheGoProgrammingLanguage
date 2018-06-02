package main

import (
	"fmt"
	"os"
)

/*

	%d				十进制整数
	%x, %o, %b		十六进制，八进制，二进制整数
	%f, %g, %e 		浮点数：3.141593  3.141592653589793  3.131593e+00（带指数）
	%t				布尔：true 或 false
	%c				字符（rune）（Unicode码点）
	%s				字符串
	%q				带双引号的字符串"abc"或带单引号的字符串'abc'
	%v 				变量的自然形式（natural format）
	%T				变量的类型
	%%				字面上的百分号标志（无操作数）



	\t				制表符
	\n 				换行符

*/

func main()  {

	// ①：Fprintf
	// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	// Fprintf 将参数列表 a 填写到格式字符串 format 的占位符中，并将填写后的结果写入 w 中，返回写入的字节数
	l,err := fmt.Fprintf(os.Stdout,"%d",111)
	fmt.Println(l,err) // l == 3    err == nil
	// 前缀 F 表示文件（File）也表明格式化输出结果应该被写入第一个参数提供的文件中。


	//--------------------------Printf 打印格式化后的内容-------------------------------------
	// ②：Printf
	// Printf(format string, a ...interface{}) (n int, err error)
	// Printf 将参数列表 a 填写到格式字符串 format 的占位符中，并将填写后的结果写入 os.Stdout 中，返回写入的字节数
	l,err = fmt.Printf("%d",222)
	fmt.Println(l,err)

	//------------------------Sprintf 拼接字符串--------------------------------------------
	// ③：Sprintf
	// Sprintf(format string, a ...interface{}) string
	// Sprintf 将参数列表 a 填写到格式字符串 format 的占位符中，并返回填写后的结果
	s := fmt.Sprintf("Sprintf %s","abc")
	fmt.Println("s:",s)

	// ④：Fprint
	// Fprint(w io.Writer, a ...interface{}) (n int, err error)
	// Fprint 将参数列表 a 填写到格式字符串 format 的占位符中，并将填写后的结果写入 w 中，返回写入的字节数
	l,err = fmt.Fprint(os.Stdout,3333)
	fmt.Println(l,err)

	//------------------------Print 打印内容-----------------------------------------------
	// ⑤：Print
	// Print(a ...interface{}) (n int, err error)
	// Print 将参数列表 a 填写到格式字符串 format 的占位符中，并将填写后的结果写入 os.Stdout 中，返回写入的字节数
	l,err = fmt.Print(222)
	fmt.Println(l,err)

	// ⑥：Sprint
	// Sprint(a ...interface{}) string
	// 将参数列表 a 组装成一个字符串，如果相邻两个参数为数字，则用空格分割
	s = fmt.Sprint(6,11,"2")
	fmt.Println(s)

	// ⑦：Fprintln
	// Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	// 参数列表 a 写入 w 中，打印输出并返回字节数，会有一个换行
	l,err = fmt.Fprintln(os.Stdout,"777")
	fmt.Println(l,err)

	//------------------------Println 打印内容--------------------------------------------
	// ⑧：Println
	// Println(a ...interface{}) (n int, err error)
	// 打印参数列表，参数列表 a 写入 os.Stdout 中，返回写入的字节数，包括一个换行符
	l, err = fmt.Println("888")
	fmt.Println(l,err)

	// ⑨：Sprintln
	// Sprintln(a ...interface{}) string
	// 将参数列表转化为字符串，并不会打印。
	s = fmt.Sprintln(1,2,3)
	fmt.Println(s)
}