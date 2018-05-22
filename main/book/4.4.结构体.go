package main

import (
	"time"
	"fmt"
	"reflect"
)

/*
	结构体：
		是一种聚合的数据类型。

 */


 // 如果结构体成员名字是以大写字母开头的，那么该成员是导出的。
type Employee struct {
	ID			int
	Name		string
	Address 	string
	DoB			time.Time
	Positon		string
	Salary		int
	ManagerID	int			// 每个值称为结构体的成员
}
 
func main() {
	struct1()				// 访问变量
	struct2()				// 结构体面值
	scale(Point{X:1,Y:2},3)		//接头体作为函数的参数和返回值

	struct3()				// 结构体的比较
	struct4()				// 结构体的嵌入
	struct5()				// 匿名成员
}

func struct1()  {
	var dibert Employee					// 声明一个 Employee类型的变量
	dibert.Salary = 5000				// 变量的成员可以用点来访问

	position := &dibert.Positon
	*position = "Senior" + *position	//成员也可以用指针访问

	//因为结构体通常通过指针处理，可以用&创建并初始化一个结构体变量，并返回结构体的地址
	pp := &Point{1,1}
	fmt.Println("pp的值为：",pp,"类型为：",reflect.TypeOf(pp))
	fmt.Println("----------------------------")
}


type Point struct {
	X 	int
	Y 	int
}

// 结构体面值
func struct2() {

	// 结构体面值 第一种写法
	p := Point{1,2}
	fmt.Println(p)

	// 结构体面值 第二种写法
	q := Point{X:3,Y:4}
	fmt.Println(q)
}

// 作为函数的参数和返回值
func scale(p Point,factor int) Point {
	q := Point{p.X * factor,p.Y * factor}
	fmt.Println(q)
	return q
}

// 考虑效率，较大的接头体通常会用指针的方式传入和返回。
// 指针的传入也方便在函数内部修改结构体成员。
func bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

// 结构体的比较
func struct3()  {
	// 如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，用 == 或 !=
	p := Point{1,2}
	q := Point{X:1,Y:2}
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p == q)


	// 可比较的结构体类型和其他可毕竟的类型一样，可以用于map的key类型
	type address struct {
		hostname 	string
		port 		int
	}
	hits := make(map[address]int)
	hits[address{"golang.org",443}] ++

}


//嵌套结构体，利用前面的Point
type Circle struct {
	Center Point
	Radius int
}
type Wheel struct {
	Circle Circle
	Spokes int
}
// 结构体的嵌入
func struct4()  {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
}


// 匿名成员
func struct5()  {

	w := Wheel{Circle{Point{0,0},2},3}
	fmt.Println(w)

	// w.X = 8 编译不通过
	// 匿名函数测试中并未像书中写的一样，待验证

}