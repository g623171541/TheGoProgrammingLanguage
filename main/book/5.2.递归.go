package main

import (
	"golang.org/x/net/html"
	"os"
	"fmt"
)

/*
	递归：
		函数递归意味着函数可以间接或直接的调用自身。
 */

func main()  {
	getHtmlHref()		//解析 html，获取所有a标签的 href
}


//解析 html，获取所有a标签的 href
func getHtmlHref()  {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr,"findlinks1:%v\n",err)
		os.Exit(1)
	}

	for _, link := range visit(nil,doc){
		fmt.Println(link)
	}
}
// 遍历HTML节点数，从每一个anchor元素的href属性获得link，并存入links字符串数组，然后返回
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr{
			if a.Key == "href" {
				links = append(links,a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links,c)
	}
	return links
}
