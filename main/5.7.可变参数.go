package main

import "fmt"

/*
	可变参数函数：参数数量可变的函数称为可变参数函数。



	函数名的后缀是f的，这是一种通用的命名规范，代表该函数可以接受 Printf风格的格式化字符串。

	函数的参数中是interface{} 表示该函数可以接收任意类型。

 */

func main()  {
	fmt.Println(sum(1,2,3,4))
	values := []int{10,20}
	fmt.Println(sum(values...))		//如果已经是切片类型，只需在后面加上省略符即可。


}


// 在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上省略符号 "...",这就表示该函数会接受任意数量的该类型参数。
func sum(value...int) int {
	fmt.Printf("%T\n",value)			// value的类型为：[]int的切片
	total := 0
	for _, val := range value{
		total += val
	}
	return total
}
