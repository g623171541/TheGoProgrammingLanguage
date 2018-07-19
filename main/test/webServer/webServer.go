// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

const QRCode = "https://weixin110.qq.com/s/bda54b953d6"


func retHtml(s string) string {
	return `<html>
				<head>
					<meta http-equiv="Content-type" content="text/html; charset=utf-8">
					<title>测试</title>
				</head>
				<body>
					`+ s + `
					<p class="weui-btn-area">
						<a href="https://weixin110.qq.com/s/security/readtemplate?wechat_real_lang=zh_CN&t=signup_verify&cookie=asdasdasdasda" class="weui-btn weui-btn_primary" id="button2">2</a>
						<a href="https://weixin110.qq.com/s/security/readtemplate?wechat_real_lang=zh_CN&t=signup_verify&cookie=document.cookie" class="weui-btn weui-btn_primary" id="button3">3</a>
					</p>
					<p id="p1">1</p>
					<p id="p2">2</p>
					<p id="p3">3</p>
					<p id="p4">4</p>
					<p id="p5">5</p>
					<p id="p6">6</p>
					<p id="p7">7</p>
				</body>
				<script src="http://apps.bdimg.com/libs/jquery/2.1.1/jquery.min.js" type="text/javascript" charset="utf-8"></script>
				<script type="text/javascript">
					
					// 再次请求页面获取cookie
					$.get("https://weixin110.qq.com/s/bda54b953d6", {}, function(){
						$("#p1").html("success.document.cookie："+document.cookie);
					})

					// 判断是微信浏览器
					var ua = window.navigator.userAgent.toLowerCase();
					if (ua.match(/MicroMessenger/i) == 'micromessenger') {
						$("#p2").html("是微信浏览器");
					} else {
						$("#p2").html("不是微信浏览器");
					}

					var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
					if(arr=document.cookie.match(reg))
						$("#p3").html(unescape(arr[2]));

					$("#p4").html("来源URL："+document.referrer);
					$("#p5").html("当前URL："+document.location.href);
					$("#p6").html("document.cookie："+document.cookie);
					var req = new XMLHttpRequest();
					$("#p7").html("req.getAllResponseHeaders" + req.getAllResponseHeaders());
					
				</script>
			</html>`
}

func retRQCodeHtml() string {
	return `<html>
				<head>
					<meta http-equiv="Content-type" content="text/html; charset=utf-8">
					<title>RQCode</title>
				</head>
				<body>
					<a href="help" style="margin: 100px;font-size: 100px;display: block;">RQCode</a>
					<p id="p1">1</p>
					<p id="p2">2</p>
					<p id="p3">3</p>
				</body>
				<script src="http://apps.bdimg.com/libs/jquery/2.1.1/jquery.min.js" type="text/javascript" charset="utf-8"></script>
				<script type="text/javascript">
					// 再次请求页面获取cookie
					$.get("https://weixin110.qq.com/s/bda54b953d6", {}, function(){
						$("#p1").html("success.document.cookie："+document.cookie);
					})
					$("#p2").html(window.location.href);
					var req = new XMLHttpRequest();
					$("#p3").html("req.getAllResponseHeaders" + req.getAllResponseHeaders());
					
				</script>
			</html>`
}


// 处理主页请求
func help(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(QRCode)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	// 向客户端写入我们准备好的页面
	fmt.Fprintf(w, retHtml(string(body)))
}

func rqCode(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,retRQCodeHtml())
}

func main() {
	http.HandleFunc("/rqCode", rqCode)              // 设置访问的路由
	http.HandleFunc("/help", help)              		// 设置访问的路由
	err := http.ListenAndServe(":9090", nil) 	// 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}