package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name 	string
	Age 	int32
	Class 	string
}

func main()  {

	student := Student{Name:"小明",Age:10}

	studentByte,_ := json.Marshal(student)
	
	fmt.Println(string(studentByte))
}
