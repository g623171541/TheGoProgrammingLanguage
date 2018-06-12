package main

import (
	"encoding/json"
	"fmt"
)

func main() {


	type Person struct {
		Name   string
		Age    int
		Gender bool
	}
	//unmarshal to struct
	var p Person
	var str = `{"Name":"junbin", "Age":21, "Gender":true}`
	json.Unmarshal([]byte(str), &p)
	//result --> junbin : 21 : true
	fmt.Println(p.Name, ":", p.Age, ":", p.Gender)

	// unmarshal to slice-struct
	var ps []Person
	var aJson = `[{"Name":"junbin", "Age":21, "Gender":true},
				{"Name":"junbin", "Age":21, "Gender":true}]`
	json.Unmarshal([]byte(aJson), &ps)
	//result --> [{junbin 21 true} {junbin 21 true}] len is 2
	fmt.Println(ps, "len is", len(ps))

	// unmarshal to map[string]interface{}
	var obj interface{} // var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)
	m := obj.(map[string]interface{})
	//result --> map[name_id:junbin Age:21 Gender:true]
	fmt.Println(m)

	// 如果想转换成 map[string]string
	var newMap map[string]string
	for k, v := range m {
		switch v.(type) {
		case float64:
			newMap[k] = fmt.Sprintf("%d", int64(v.(float64)))
		default:
			newMap[k] = fmt.Sprintf("%s", v)
		}
	}

	// 这样做就是在解析请求参数的时候更加通用，直接解析为 map[string]interface{} ,通过内部函数再次转换


	//unmarshal to slice
	var strs string = `["Go", "Java", "C", "Php"]`
	var aStr []string
	json.Unmarshal([]byte(strs), &aStr)
	//result --> [Go Java C Php]  len is 4
	fmt.Println(aStr, " len is", len(aStr))
}