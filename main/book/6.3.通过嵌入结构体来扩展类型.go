package main

import (
	"image/color"
	"fmt"
	"encoding/json"
)

/*
	
*/

func main()  {
	Color()

	Color2()
}

func Color()  {
	red := color.RGBA{255,0,0,255}
	blue := color.RGBA{0,0,255,255}
	var p = ColoredPoint{Point63{1,1},red}
	var q = ColoredPoint{Point63{5,4},blue}
	pJson,_ := json.Marshal(p)

	fmt.Println(string(pJson))
	fmt.Println(q)

}

type Point63 struct {
	X float64
	Y float64		// 这个必须大写才能被访问到，才能在打印ColoredPoint结构体的时候打印出来
}
type ColoredPoint struct {
	Point63
	Color color.RGBA
}
// 在 ColoredPoint 中完全可以定义三个字段，我们却将 Point63 这个类型嵌入其中，使ColoredPoint 来提供x和y两个字段。
// 我们可以直接认为嵌入的字段就是ColoredPoint自身的字段，而完全不需要在调用时支出Point63




type ColorPoint2 struct {
	Point63
	color.RGBA
}
// 这种类型的值变回拥有Point63 和RGBA类型的所有方法

func Color2()  {
	cp := ColorPoint2{Point63{1,1},color.RGBA{0,255,255,1}}
	fmt.Println(cp)
	js,_ := json.Marshal(cp)
	fmt.Println(string(js))


}