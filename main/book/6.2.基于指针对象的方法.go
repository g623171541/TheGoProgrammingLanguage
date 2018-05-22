package main

import "fmt"

/*
	基于只针对象的方法：
 */

func main()  {
	r := &Point6{1,2}
	r.ScaleBy(2)
	fmt.Println(*r)


}

type Point6 struct {
	x, y float64
}

// 这个方法的名字是 (*Point6).ScaleBy
func (p *Point6) ScaleBy( factor float64 )  {
	p.x *= factor
	p.y *= factor
}