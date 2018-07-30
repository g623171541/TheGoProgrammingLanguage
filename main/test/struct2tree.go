package main

import (
	"fmt"
	"encoding/json"
)

type Data struct {
	Id           int32
	ParentId     int32
	Appid        int32
	ProductCode  string
	Code         string
}

type Result struct {
	Id          int32
	Appid       int32
	ProductCode string
	Code        string
	Child		[]Result
}

//
func main()  {
	data1 := Data{Id:1,ParentId:0,Appid:10000,ProductCode:"code1",Code:"code1"}
	data2 := Data{Id:2,ParentId:0,Appid:10000,ProductCode:"code2",Code:"code2"}
	data3 := Data{Id:3,ParentId:1,Appid:10000,ProductCode:"code3",Code:"code3"}
	data4 := Data{Id:4,ParentId:2,Appid:10000,ProductCode:"code4",Code:"code4"}
	data5 := Data{Id:5,ParentId:2,Appid:10000,ProductCode:"code5",Code:"code5"}

	var dataList []Data
	dataList = append(dataList,data1,data2,data3,data4,data5)


	fmt.Println(ToParentChild(dataList))
	j,_ := json.Marshal(ToParentChild(dataList))
	fmt.Println(string(j))
}

func ToParentChild(data []Data) (result []Result) {
	resData := data
	var tree []Result

	// 首先找到最小的ParentId
	minParentId := int32(0)
	for _,item := range data{
		if minParentId > item.ParentId {
			minParentId = item.ParentId
		}
	}

	// 取出第一层的
	for i := 0; i < len(resData) ; i++ {
		if resData[i].ParentId == minParentId {
			var empty []Result
			obj := Result{resData[i].Id,resData[i].Appid,resData[i].ProductCode,resData[i].Code,empty}

			tree = append(tree,obj)
		}
	}

	run(data,tree)

	return tree

}

// 递归放值
func run(data []Data,chiArr []Result) {
	for i := 0 ; i < len(chiArr) ; i++ {
		for j := 0 ; j < len(data) ; j++ {
			if chiArr[i].Id == data[j].ParentId {
				var empty []Result
				obj := Result{data[j].Id,data[j].Appid,data[j].ProductCode,data[j].Code,empty}
				chiArr[i].Child = append(chiArr[i].Child,obj)
			}
		}
		run(data,chiArr[i].Child)
	}
}
