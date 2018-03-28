package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
	"reflect"
	"bytes"
)

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

	标准库中有四个包对字符串的处理尤为重要：bytes、strings、strconv、unicode
	1.strings包	提供了许多如字符串的查询、替换、比较、截断、拆分、合并等功能
	2.bytes包	提供了许多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。因为字符串是制度的，因此逐步构建字符串会导致很多分配和复制。在这种情况下，使用bytes.Buffer类型将会更有效。
	3.strconv包	提供了布尔值、整型数、浮点数和对应字符串的相互转换。还提供了双引号转义相关的转换。
	4.unicode包	提供了IsDigit、IsLetter、IsUpper、IsLower等类似功能，他们勇于给字符分类。每个函数有一个单一的rune类型的参数，然后返回一个布尔值。而像ToUpper和ToLower之类的转换函数将用于rune字符的大小转换。

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

	s = "hello,世界"
	fmt.Println(len(s))						//长度12
	fmt.Println(utf8.RuneCountInString(s))	//长度8

	for i,r := range "hello,世界" {
		fmt.Printf("%d\t%q\t%d\n",i,r,r)
		/*
			0	'h'	104
			1	'e'	101
			2	'l'	108
			3	'l'	108
			4	'o'	111
			5	','	44
			6	'世'	19990
			9	'界'	30028
		*/
	}

	//截取字符串"."之前的和"/"之后的内容
	fmt.Println(basename1("a/b//d/ccc.go"))
	fmt.Println(basename2("aaaa/dd.go"))

	//利用递归将123456789每隔三位加一个逗号
	fmt.Println(comma("123456789"))

	var abc string = "abc"
	ABC := []byte(abc)						//[]byte()转换是分配了一个新的字节数组用于保存字符串数据的拷贝，然后引用这个底层的字节数组。
	fmt.Println(reflect.TypeOf(ABC))		//[]uint8
	fmt.Println(ABC)						//[97 98 99]
	fmt.Println(string(ABC))				//abc

	//将数字数组转换为字符串
	fmt.Println(intsToString( []int{1,2,3}) )

}

//bytes.Buffer变量：bytes包提供了Buffer类型用于字节slice的缓存。
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i,v := range values{
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf,"%d",v)
	}
	buf.WriteString("]")
	return buf.String()
}

//将字符串每隔三位加一个逗号（comma）
func comma(s string) string {
	len := len(s)
	if len <= 3 {
		return s
	}
	return comma(s[:len-3]) + "," + s[len-3:]
}

//截取字符串"."之前的和"/"之后的内容
func basename1(str string) string {
	//discard last '/' and everything before.
	for i := len(str) - 1; i >= 0 ; i-- {
		if str[i] == '/' {
			str = str[i+1:]
			break	//一定要结束当前循环
		}
	}

	//preserve everything before last '.'
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '.' {
			str = str[:i]
			break
		}
	}

	return str
}
//用strings.LastIndex函数优化
func basename2(s string) string {
	index := strings.LastIndex(s,"/")
	s = s[index+1:]

	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}

	return s
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