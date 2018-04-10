package main

/*
	文本和HTML：
		text/template  或者  html/template 包提供
		
 */

func main()  {
	
}

const templ = `{{.TotalCount}} issues:
{{range .Items}}-------------------------------
Number:	{{.Number}}
User:	{{.User.Login}}
Title:	{{.Title | printf "%.64s"}}
Age:	{{.CreateAt | daysAgo}} days
{{end}}
`