package main

import (
	"fmt"
	"sort"
	"reflect"
)

/*
	map：
		类型：map[K]V；  K：Key V：Value
 */

func main()  {

	map1()			// map的创建
	map2()			// map的操作
	map3()			// map的value类型可以是map或者slice
}

func map1()  {
	//内置的 make 函数创建map
	ages := make(map[string]int)

	//创建方式二
	age := map[string]int{
		"alice":	31,
		"charlie":	34,
	}

	//方式二相当于
	a := make(map[string]int)
	a["alice"] = 31
	a["charlie"] = 34

	fmt.Println(ages,age)
}

func map2()  {
	ages := map[string]int{
		"alice":	31,
		"charlie":	34,
	}

	// 删除元素
	delete(ages,"alice")
	delete(ages,"abc")		//没有对应的key也不会报错，查找失败将返回value类型对应的零值
	fmt.Println(ages)

	// 增加元素
	ages["paddy"] = 26
	ages["jack"] = 25

	//禁止对map的元素进行取地址操作，因为map可能随着元素数量的增长而重新分配更大的内存空间，从而导致之前的地址无效

	// 遍历：顺序是随机的
	for name, age := range ages{
		fmt.Println(name,age)
	}

	fmt.Println("-------------遍历---按key排序-------------")
	// 对遍历出来的key 进行排序
	var names []string
	for name := range ages{
		names = append(names,name)
	}
	sort.Strings(names)
	for _,name := range names{
		fmt.Println(name,ages[name])
	}

	// map类型对应的零值是nil，想一个nil值的map存入元素会导致panic异常
	var m map[string]int
	fmt.Println(m == nil)		// true
	fmt.Println(len(m) == 0)	// true
	// m["a"] = 1					// panic异常

	// map的下标语法产生两个值：①值 ②布尔值：用于报告元素是否存在，一般命名为OK，特别适合马上用于if条件哦按段部分。
	if age, ok := ages["paddy"]; !ok {
		fmt.Println("不存在",age)
	}
}

// map的value类型也可以是一个聚合类型，如map和slice
func map3()  {

	var graph = make(map[string]map[string]bool)
	graph["a"] = map[string]bool{
		"aa":true,
		"bb":false,
	}

	fmt.Println(reflect.TypeOf(graph))
	fmt.Println(graph["a"]["aa"])
}
