package main

import (
	"crypto/hmac"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

func main() {

	//HmacMD5
	str := []byte("123")
	key := []byte("KEHTHJHJ4IWJ3I9J")
	h := hmac.New(md5.New, key)
	h.Write(str)
	fmt.Println(hex.EncodeToString(h.Sum(nil)))

}