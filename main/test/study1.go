package main

//fmt 包实现了格式化 IO（输入/输出）的函数
import (
	"unsafe"
	"fmt"
	"math"
	"errors"
)



func main()  {

	var name = "go"
	age := 10

	println(name,age)

	const (
		UnKnown = 0
		Male = 1
		Female = 2
	)

	const (
		aaa = "a"
		bbb = len(aaa)
		ccc = unsafe.Sizeof(aaa)
	)
	println("aaa bbb ccc =",aaa,bbb,ccc)

	const(
		a = iota
		b
		c = 100
		d = iota
		e

	)
	println(a,b,c,d,e)


	SelectMethod()

	//创建数组(声明长度)
	var array1 = [5]float32{1.0,2.2,3.3,4.4}
	fmt.Printf("array1--- type:%T %d\n", array1,len(array1))

	//创建数组(不声明长度)
	var array2 = [...]int{6, 7, 8}
	fmt.Printf("array2--- type:%d %d\n", array2,len(array2))

	println("数组的指针为 ",&array1)

	//for循环
	forCirculate()


	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(100))

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())

	//指针
	printZhizhen(age)

	//结构体
	type Books struct {
		title string
		money float32
		subject string
		book_id int
	}
	//Book := Books{"Go",98.00,"学习",2017038032348}
	Book := Books{title:"go",money:198.99,subject:"study",book_id:100000}
	println(Book.subject,Book.title)

	//切片：直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3
	s := []int{1,2,3}
	println("--------------",s[0])

	numbers := make([]int,2,3)
	printSlice(numbers)

	var abc []int
	printSlice(abc)
	if abc == nil {
		println("切片为空")
	}
	//创建切片
	num := []int{1,2,3,4,5,6,7,8}
	fmt.Println("num = ",num)		//打印原始切片
	fmt.Println("num[1:4] = ",num[1:4])  //前开后闭区间
	fmt.Println("num[:2] = ",num[:2])

	num1 := num[:3]
	fmt.Println("num1 = ",num1)

	var num2 []int
	fmt.Println("num2 = ",num2)

	num2 = append(num2,0) 	/* 向切片添加一个元素 */
	fmt.Println("num2 = ",num2)

	num2 = append(num2,1,2,3,4) /* 同时添加多个元素 */
	fmt.Println("num2 = ",num2)

	num3 := make([]int,len(num1),cap(num1)*2)  /* 创建切片 num3 是之前num1切片的两倍容量*/
	printSlice(num3)

	//range 的用法
	num4 := []int{2,3,4}
	sum := 0
	for i, num := range num4{	//不需要序号时 可将i换成 "_"下划线代替
		sum += num
		fmt.Print(i)
	}
	fmt.Println("sum = ",sum)

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "Aa" {
		fmt.Println(i, c)
	}

	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//map 无序的键值对集合，类似于OC中的字典。无序是因为map是用hash表来实现的
	var countryMap = make(map[string]string)   	//如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	//countryMap := make(map[string]string)		//也可以这样初始化
	countryMap["A"] = "北京"
	countryMap["B"] = "上海"
	for key , city := range countryMap{
		fmt.Println(key,"对应的城市为：",city)
	}
	//for key := range countryMap{
	//	fmt.Println(key,"对应的城市为：",countryMap[key])
	//}

	city,ok := countryMap["C"]
	/* 如果 ok 是 true, 则存在，否则不存在 ，city对应value，ok==false时city为空*/
	fmt.Println("city:",city,"OK:",ok)

	//创建map，直接赋值，
	map1 := map[string]string {"A":"beijing","B":"shanghai","C":"c"}
	delete(map1, "C")			//删除map元素
	fmt.Println(map1)

	fmt.Println("计算阶乘结果为：",jiecheng(4))

	var phone Phone
	phone = new(iPhone)
	phone.call()

	phone = new(AndroidPhone)
	phone.call()


	//go语言的错误处理，err为一个接受错误信息的参数，result是计算结果的值
	result, err:= Sqrt(-1)

	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}

}



type error interface {
	Error() string
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// 实现
	return 999,nil
}


//interface类型：它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
type Phone interface {
	call()
}
type AndroidPhone struct {}

func (phone AndroidPhone) call()  {
	fmt.Println("AndroidPhone")
}

type iPhone struct {}

func (phone iPhone) call()  {
	fmt.Println("iPhone")
}



//计算阶乘：利用递归运算
func jiecheng(n int)  int{
	if n > 0 {
		return n * jiecheng(n - 1)
	}
	return 1
}

//切片的打印
func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
	// 	len() 方法获取长度   	cap() 可以测量切片最长可以达到多少
}


//打印指针
func printZhizhen(age int)  {
	//声明指针：
	var ip *int			//指针指向整形
	ip = &age
	println("打印ip（地址）为：",ip)
	println("打印age的地址为：",&age)
	println("打印age的值为：",age)
	println("打印*ip的值为：",*ip)
}

/* 定义函数 */     //---------------------------------函数的三种形式
type Circle struct {
	radius float64
}
//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {

	if (num1 > num2) {
		return num1
	} else {
		return num2
	}
}

//闭包
func getSequence() func() int {
	i:=0
	return func() int {
		i+=1
		return i
	}
}


func SelectMethod() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	//-------------------------------------------------------看不懂的 = <- 符号`
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

//for循环
func forCirculate()  {

	numbers := [4]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a <= len(numbers); a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
}