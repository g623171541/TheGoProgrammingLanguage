package main

import "fmt"

/*
	
 */

func main()  {
	s := "hello"
	fmt.Println( len(s) )			//len()函数获取字符串的长度

	fmt.Println(s[0])				//104 取第几个字符的ASCII码表

	fmt.Println(s[0:2])				//字符串截取：0作为起始位置可省略 s[:2]，len(s)作为末尾位置可省略 s[2:]
	fmt.Println("123" + s[0:2])		//字符串的拼接：用+号

	101页

}