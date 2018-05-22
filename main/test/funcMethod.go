package main

import "fmt"

//定义一个结构体,相当于OC中的定义一个类
type DivError struct {
	div1 int
	div2 int
}

//实现接口：就是实现这个类中对象的方法
func (div DivError) Error() string {
	str := `
		不能计算，因为被除数为0
		除数：%d
		被除数：0
	`
	return fmt.Sprintf(str,div.div1)
}

//实现除法运算的函数
func Divide(varDiv1 int,varDiv2 int) (result int,errMsg string) {
	if varDiv2 == 0 {
		err := DivError{
			div1:varDiv1,
			div2:varDiv2,
		}
		errMsg = err.Error()
		return
	}else {
		return varDiv1 / varDiv2,""
	}
}

func main()  {

	if result , errMsg := Divide(100,10); errMsg != "" {
		fmt.Println(result)
		fmt.Println(errMsg)
	}else {
		fmt.Println("123")
	}
}
