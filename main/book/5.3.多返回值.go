package main

import "fmt"

/*
	多返回值：
		在Go中，一个函数可以返回多个值。多个返回值的时候要用括号括起来。
 */

func main()  {
	a, err := fun1()
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("值为：",a)
	}


	a,b := fun2()
	fmt.Println(a,b)
}

func fun1() (int, error) {
	return 1, fmt.Errorf("错误信息:%d 解析失败",1)

	// 用 fmt.Errorf 输出详细的错误信息。
}


// 如果一个函数将所有的返回值都显示的变量名，那么该函数的return语句可以省略操作数，这称之为 bare return
func fun2() (a int, b bool) {
	return
}