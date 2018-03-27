package main

import "fmt"

/*
	转义字符：
			\a	响铃
			\b	退格
			\f	换页
			|r 	回车
			\t	制表符
			\v	垂直制表符
			\'	单引号
			\"	双引号
			\\	反斜杠

	Unicode：Unicode码点对应Go语言中的rune整数类型（rune是int32等价类型）


 */

func main()  {
	s := "hello"
	fmt.Println( len(s) )			//len()函数获取字符串的长度

	fmt.Println(s[0])				//104 取第几个字符的ASCII码表

	fmt.Println(s[0:2])				//字符串截取：0作为起始位置可省略 s[:2]，len(s)作为末尾位置可省略 s[2:]
	fmt.Println("123" + s[0:2])		//字符串的拼接：用+号

	//使用反引号
	sss := `Go is a tool for managing Go source code
Usage:
		go command [arguments]
...
`
	fmt.Println(sss)

	//判断前缀后缀
	HasPrefix("2333332","2")
	HasSuffix("askdalsdaaaa","aaa")
	Contains("1234512345","34")
}

//测试一个字符串是否是另一个的前缀
func HasPrefix(s,prefix string) bool {
	b := len(s) > len(prefix) && s[:len(prefix)] == prefix
	fmt.Println(b)
	return b
}
//测试一个字符串是另一个的后缀
func HasSuffix(s, suffix string) bool {
	b := len(s) > len(suffix) && s[len(s)-len(suffix) : ] == suffix
	fmt.Println(b)
	return b
}
//测试一个字符串包含另一个字符串
func Contains(s,subStr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:],subStr) {
			fmt.Println("包含")
			return true
		}
	}
	fmt.Println("不包含")
	return false
}