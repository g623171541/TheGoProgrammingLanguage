package main

import "fmt"

//功能：声明一个常量、一个函数、两个变量


const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n\n", f, c)

	//-------------------------------------------------------------
	const freezingF, boilingF = 32.0, 212.0
	//调用函数
	fmt.Printf("%g°F = %g°C \n",freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C \n",boilingF, fToC(boilingF))

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}