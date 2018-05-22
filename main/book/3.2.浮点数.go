package main

import (
	"fmt"
	"math"
)

/*
	两种精度的浮点数：float32 float64

	最大的float32：math.MaxFloat32
	最大的float64：math.MaxFloat64

	通常使用float64，因为float32类型的累计计算误差很容易扩散，能精确表示的正整数并不是很大
		：float32有效bit位只有23个，其他bit位用于指数和符号，当整数大于23bit能表达的范围时，就会出现误差。


 */

func main()  {
	//float32最大值
	fmt.Println("float32最大值 : ", math.MaxFloat32 )
	//float64最大值
	fmt.Println("float64最大值 : " , math.MaxFloat64 )

	var f float32 = 16777216	// 1<<24
	fmt.Println(f == f+1)		// true

	//测试一个数是否是非数NaN
	nan := math.NaN()
	fmt.Println( math.IsNaN(nan) )
	fmt.Println(nan == nan,nan > nan,nan < nan)	//NaN和任何数都不相等


}