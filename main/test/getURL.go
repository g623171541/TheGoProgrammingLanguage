package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
)

/*
	作用是获取到URL值的html
	用法： go run getURL.go http://www.baidu.com
*/

func main()  {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch:%v\n",err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch reading %s : %v\n",url,err)
			os.Exit(1)
		}
		fmt.Printf("%s",b)
	}

}

