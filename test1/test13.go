package main

import (
	"fmt"
	"net/http"
)

//处理器函数
func handler (w http.ResponseWriter ,r *http.Request){
	fmt.Println("hello")
	fmt.Println(w,"hello world ",r.URL.Path)
}
func main() {
	http.HandleFunc("/",handler)
	//创建路由
	http.ListenAndServe("127.0.0.1:8082",nil)

}
