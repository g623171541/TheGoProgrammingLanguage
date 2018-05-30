package main

import (
	"fmt"
	"math"
	"time"
)

/*

 */

func main()  {

	pointDistance()

	rocket()

	distance()
}

func distance()  {
	p := Point64{1,2}
	q := Point64{4,6}

	distanc := Point64.Distance
	fmt.Println(distanc(p,q))
	fmt.Printf("%T\n",distanc)
}

// --------------------
func rocket()  {
	r := new(Rocket)

	// time.AfterFunc 函数是在指定的延迟时间之后来执行另一个函数

	//time.AfterFunc(1 * time.Second, func() {
	//	r.launch()
	//} )

	// 其中上面的函数可以更改为
	time.AfterFunc(1 * time.Second, r.launch )

	time.Sleep(2 * time.Second) //要保证主线比子线“死的晚”，否则主线死了，子线也等于死了
	// 运行过程：运行了1秒后，打印出 "launch..." ，又过了1秒，程序退出
}
type Rocket struct {
	
}

func (r *Rocket) launch()  {
	fmt.Println("launch...")
}

// ---------------------------------
func pointDistance()  {
	p := Point64{1,2}
	q := Point64{4,6}

	// p.Distance 叫做选择器，选择器会返回一个方法值（一个将方法绑定到特定接收器变量的函数）
	// 这个函数可以不通过指定其接收器即可被调用；即调用时不需要指定接收器，只需传入函数的参数即可。
	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))
}

type Point64 struct {
	x,y float64
}
func (p Point64) Distance(q Point64) float64  {
	return math.Hypot(q.x-p.x, q.y-p.y)
}