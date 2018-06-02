package main

// 导入依赖包
import "fmt"

// 常量定义
const NAME string = "imooc"

// 全局变量的声明和赋值
var a string = "慕课网"

// 一般类型声明
type imoocInt int

// 结构的声明
type Learn struct {

}

// 声明接口
type Ilearn interface {
	// 一般声明接口的名字首字母用大写的I
}

// 函数定义
func learmImooc()  {
	fmt.Println("learn imooc")
}

// main()函数
func main()  {
	learmImooc()
}