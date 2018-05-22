package main

import (
	"reflect"
	"strings"
	"time"
	"fmt"
)

func main()  {
	user := User{5, "zhangsan", "pwd", time.Now()}
	data := Struct2Map(user)
	fmt.Println(data)

	info := Info{"paddy","男","深圳市"}
	fmt.Println( Struct2MapString(info) )
}

type User struct {
	Id        int64
	Username  string
	Password  string
	Logintime time.Time
}

type Info struct {
	Name 	string
	Sex 	string
	Address string
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}

func Struct2MapString(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		data[strings.ToLower(t.Field(i).Name)] = v.Field(i).String()
	}
	return data
}