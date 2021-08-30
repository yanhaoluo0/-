package main

import (
	"fmt"
	"net/http"
)

func prin(a string) {
	n := 10
	for {
		fmt.Print(a)
		n--
		if n == 0 {
			break
		}
	}
}

type superme struct {
	head string
	body string
}

func main() {
	res1, err1 := http.Get("https://www.baidu.com")
	if err1 != nil {
		fmt.Println("get err1=", err1)
		return
	}
	prin("*")
	fmt.Print("information")
	prin("*")
	fmt.Println()
	fmt.Println("res1.Status\t=\t", res1.Status)
	fmt.Println("res1.Header\t=\t", res1.Header)
	fmt.Println("res1.Body\t=\t", res1.Body)
	buf := make([]byte, 1024)
	var temp string
	for {
		n, _ := res1.Body.Read(buf)
		if n == 0 {
			//fmt.Println("body内容读取错误 ：", err)
			break
		}
		temp += string(buf[:n])
	}
	fmt.Println("temp=", temp)
}
