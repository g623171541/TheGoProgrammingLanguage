package main

import (
	"encoding/json"
	"log"
	"fmt"
)

/*
	JSON:
		是一种用于发送和接受结构化信息的标准协议，但不是唯一的一个标准协议，还有 XML ASN.1 Google的Protocol Buffers

		基本的JSON类型有数字、布尔值、字符串、和中字符串


 */


func main()  {
	json1()				// 编码：结构体slice 转 JSON （可带格式）
	json2()				// 解码：JSON 转 GO语言的数据结构
	json3()				// Decoder Encoder 用来支持JSON数据了的读写
}

type Movie struct {
	Title	string
	Year	int `json:"released"`				// 构体成员tag，使其转为json的时候Year变为 released
	Color 	bool `json:"color,omitempty"`		// tag可以是任意的字符串面值，第一个值用于指定JSON对象的名字
	Actors	[]string							// omitempty 表示Go语言结构体成员为空或零值时不生成JSON对象
}

// 结构体slice
var movies = []Movie{
	{Title:"Casablanca",Year:1942,Color:false,Actors:[]string{"Humphrey","Ingrid"}},
	{Title:"Cocol",Year:1967,Color:true,Actors:[]string{"Paul"}},
}

// 编码：结构体slice 转 JSON
func json1()  {

	//这样的数据结构特别适合JSON格式，结构体slice与JSON互转：json.Marshal函数
	//*************** 结构体slice 转 JSON *********************
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed %s",err)
	}
	fmt.Printf("-------  结构体slice 转 JSON:\n%s\n",data)

	//*************** 带格式化的JSON *********************
	data2, err := json.MarshalIndent(movies,"","	")		// ""：每一行输出的前缀；"		"：每一个层级的缩进
	if err != nil {
		log.Fatalf("JSON marshaling failed %s",err)
	}
	fmt.Printf("-------  带格式的JSON:\n%s\n",data2)

}

// 解码：
func json2()  {
	data, _ := json.Marshal(movies)

	var titles []struct{		// 定义合适的Go语言数据结构，选择性的解码JSON中感兴趣的成员。
		Title 	string			// slice将被只含有Title信息值填充，其他JSON成员被忽略
		}
	if err := json.Unmarshal(data,&titles); err != nil {
		log.Fatalf("JSON unmarshaling failed : %s",err)
	}
	fmt.Println("-------  JSON 转 结构体slice")
	fmt.Println(titles)			// [{Casablanca} {Cocol}]


	type Message struct {
		Name string
		Body string
		Time int64
	}
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, _ := json.Marshal(m)
	var bm Message				// 声明一个变量用于存放解码后的数据
	json.Unmarshal(b, &bm)
	fmt.Println(bm)

}

// Decoder Encoder 用来支持JSON数据了的读写。
func json3()  {

}
