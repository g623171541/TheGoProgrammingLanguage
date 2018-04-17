package main

import (
	"strings"
	"fmt"
	"sort"
)

/*
	拥有函数名的函数只能在包级语法块中被声明，通过函数字面量，我们可以绕过这一限制，在任何表达式中表示一个函数值。
	函数字面量的语法和函数声明相似，区别在于func关键字后没有函数名。
	函数字面量是一种表达式，它的值被称为匿名函数。
 */



func main()  {
	// 函数字面量允许我们在使用函数时，再定义它。通过这种技巧，我们可以在改写之前对 strings.Map的调用。
	strings.Map(func(r rune) rune {
		return r + 1
	}, "HAL-9000" )

	f := squares()
	fmt.Println(f())		// 1
	fmt.Println(f())		// 4
	//例子证明，函数值不仅仅是一串代码，还记录了状态。


}


// 返回值的类型为：func() int 是一个匿名函数
func squares() (func() int) {
	var x int				// x变量的生命周期不由它的作用域决定：变量x仍然隐式的存在于 f 中

	// 每次调用匿名函数时都会先让 x 的值加1，再返回x的平方。
	// 第二次调用squares时，会生成第二个x变量，并返回一个新的匿名函数。新匿名函数操作的是第二个x变量。
	return func() int {
		x++
		return x * x
	}
}




// ---------------------- 拓扑排序 ---------------
// []string 是切片slice类型
var prereqs = map[string][]string{
	"algorithms":	{"data structures"},
	"calulus":		{"linear algebra"},
	"compilers":	{
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":		{"discrete math"},
	"databases":			{"data structures"},
	"discrete math":		{"intro to programming"},
	"formal languages": 	{"discrete math"},
	"networks":				{"operating systems"},
	"operating systems":	{"data structures","computer organization"},
	"programming languages":{"data structures","computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)	//当匿名函数需要被递归调用时，必须首先声明一个变量，再将匿名函数赋值给这个变量
	visitAll = func(items []string) {	//如果不分成两部分，函数字面量无法与visitAll绑定，我们也无法递归调用该匿名函数
		for _, item := range items{
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order,item)
			}
		}
	}
	var keys []string
	for key := range m{
		keys = append(keys,key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return  order
	
}