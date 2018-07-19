package main

/*
	功能：简单的web服务：浏览器输入 http://localhost:8080/hello  会在页面上输出  URL.path = "/hello"
*/

import (
	"net/http"
	"fmt"
)

func main()  {
	// 注册路由
	http.HandleFunc("/",handler)

	// 开始监听，处理请求，返回响应
	http.ListenAndServe(":8888",nil)
}

func handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"URL.path = %q\n",r.URL.Scheme)
	w.Write([]byte("hello Write\n"))
	fmt.Fprintln(w, "hello world")
}

/*
	客户端：client --> request
	服务端：server --> response
		服务端在接受request的过程中最重要的是 路由router，即实现一个 Multiplexer器。
		Go中既可以使用内置的mutilplexer --- DefaultServeMux 也可以自定义。
		Multiplexer 路由的目的就是为了找到处理函数（handler）

	简单流程：Client -> Request -> Multiolexer(router) -> handler -> Response -> Client

	Golang 中的 Multiplexer 基于ServeMux 结构，同时也实现了handler接口。

*/
