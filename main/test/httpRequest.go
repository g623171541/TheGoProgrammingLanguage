package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"strings"
	"net/url"
)

func main()  {
	httpGet()
}

// get
func httpGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

// post
func httpPost()  {
	//第二个参数要设置成”application/x-www-form-urlencoded”，否则post参数无法传递。
	resp,err := http.Post("http://www.baidu.com", "application/x-www-form-urlencoded", strings.NewReader("name=cjb"), )
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		 fmt.Println(err)
	}

	fmt.Println(string(body))
}

// postForm
func httpPostForm()  {
	resp,err := http.PostForm("http://www.baidu.com",url.Values{
		"key": {"Value"}, "id": {"123"},
	})
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

// post-----复杂的请求（设置头参数、cookie之类的数据），可以使用http.Client的Do()方法
func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.baidu.com", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
