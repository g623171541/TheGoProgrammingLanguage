package main

import (
	"fmt"
)

/*
	slice：切片
		代表边长的序列，序列中每个元素都有相同的类型。一个slice类型一般写作 []T，其中T代表slice中元素的类型。
		slice的语法和数组很像，只是没有规定的长度。

	一个slice有三部分组成：
		指针：指向第一个slice元素对应的底层数组元素的地址。注意：slice的第一个元素不一定就是数组的第一个元素。
		长度：对应slice元素的数目。长度 == len()
		容量：长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置。容量 == cap()
 */


func main()  {

	slice1()

	a := []int{1,2,3,4,5,6,7} 	// slice类型的变量a，初始化的时候没有指明序列的长度，会隐式的创建一个大小合适的长度
	fmt.Println( reverse(a) )	//倒序排列

	slice2()		//slice的比较，不能比较


}

func slice1()  {

	//定义数组
	months := [...]string{1:"January",2:"February",3:"March",4:"April",5:"May",6:"June",7:"July",8:"August",9:"September",10:"October",11:"November",12:"December"}
	// 默认数组的第一个元素下标是0，但是months跳过第0个元素，从第一个开始初始化，则第0个自动化初始化为空字符串
	fmt.Println("长度：",len(months))
	fmt.Println("容量：",cap(months))
	fmt.Println("一月份为：",months[1])
	fmt.Println("切片操作：",months[1:3],"只有 3-1 个元素，不包括下标是3的元素")

	Q2 := months[4:7]			//slice的切片操作，用于创建一个新的slice，多个slice之间可以共享底层的数据
	summer := months[6:9]		//长度是 3 容量是 7，因为是起始位置到底层数据的结尾位置
	fmt.Println(Q2)
	fmt.Println(summer)

	//找到都包含的元素
	for _,q := range Q2{
		for _,s := range summer{
			if q == s {
				fmt.Printf("%s appears in both\n",s)
			}
		}
	}

	// 如果切片操作超出cap(s)的上限将报panic错误，但是超出len(s)则意味着扩展了slice，新的slice长度会变大
	// fmt.Println(summer[:20])		//panic：slice bounds out of range
	fmt.Println(summer[:5])

}

//数组的倒序排列
func reverse(s []int) []int {
	for i,j := 0,len(s)-1 ; i < j ; i,j = i+1,j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func slice2()  {
	// 和数组不同的是，slice之间不能比较，不能使用 == 操作符来判断两个slice是否全部相等
	// 不过标准库提供了高度优化的bytes.Equal函数来判断两个字节型slice是否相等 （[]byte）

	a := []string{"a","b","c"}
	b := []string{"a","b"}
	fmt.Println(equal(a,b))

}

//两个slice的比较
func equal(x,y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x{
		if x[i] != y[i] {
			return false
		}
	}
	return true
}