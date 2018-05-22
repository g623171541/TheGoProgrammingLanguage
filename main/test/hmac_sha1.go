package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"encoding/base64"
)

func main() {

	//hmac ,use sha1
	key := []byte("11111111111111111")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte("22222222222222222"))
	fmt.Println(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

}