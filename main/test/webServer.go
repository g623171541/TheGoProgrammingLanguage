package main

/*
	功能：简单的web服务：浏览器输入 http://localhost:8080/hello  会在页面上输出  URL.path = "/hello"
*/

import (
	"net/http"
	"log"
	"fmt"
)

func main()  {
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}

func handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"URL.path = %q\n",r.URL.Scheme)
}

