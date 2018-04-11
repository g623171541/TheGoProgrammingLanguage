package main

import (
	"time"
	"html/template"
	"log"
)

/*
	文本和HTML：
		text/template  或者  html/template 包提供

	并没有看懂，可后续用到的时候再研究
		
 */

func main()  {
	// 创建并分析定义的模板 templ
	report, err := template.New("report").			//创建并返回一个模板
		Funcs(template.FuncMap{"daysAgo":daysAgo}).			//将daysAgo等自定义的函数注册到模板中
			Parse(templ)									//Parse函数分析模板
	if err != nil {
		log.Fatal(err)
	}

	result, err := github.Search
}

const templ = `{{.TotalCount}} issues:
{{range .Items}}-------------------------------
Number:	{{.Number}}
User:	{{.User.Login}}
Title:	{{.Title | printf "%.64s"}}
Age:	{{.CreatedAt | daysAgo}} days
{{end}}
`
// 模板中 {{rang .Items}} 和 {{end}} 对应一个循环action
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}