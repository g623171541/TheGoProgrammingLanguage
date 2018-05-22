package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "F123456789"
	s := string([]byte(str)[:1])
	if s >= "A" && s <= "Z" {
		s = strings.ToLower(s)
	}else {
		fmt.Println(str)
	}

	content := str[1 : len(str)]
	fmt.Println(s + content)

}