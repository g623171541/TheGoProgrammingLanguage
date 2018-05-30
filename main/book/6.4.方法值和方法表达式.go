package main

import "fmt"

/*

 */

func main()  {
	for i:=1; i <= 30; i++ {
		if i % 2 == 0{
			fmt.Print(i)
			fmt.Print(",")
		}
	}
}

