package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func HttpGET11(u string) (result string, err error) {
	re1, err1 := http.Get(u)
	if err1 != nil {
		err=err1
	}
	defer re1.Body.Close()
	buf:=make([]byte,1024*4)
	for{
		n,e :=re1.Body.Read(buf)
		if n==0{
			fmt.Println("read出错= ",e)
			continue
		}
		result +=string(buf[:n])
	}
	return result,err
}
func Select(st, en int) {
	fmt.Printf("正在查询%d到%d的内容\n", st, en)
	for i := st; i <= en; i++ {
		url := "https://tieba.baidu.com/f?kw=%E8%A1%A2%E5%B7%9E%E5%AD%A6%E9%99%A2&ie=utf-8&pn=" +
			strconv.Itoa((i-1)*50)
		fmt.Println("url=\t", url)
		information, err := HttpGET11(url)
		if err != nil {
			fmt.Println("页面获取时出错")
			break
		}
		//内容写入到文件
		fileName1 :=strconv.Itoa(i) +".html"
		f ,err01 :=os.Create(fileName1)
		if err01!=nil{
			fmt.Println("文件错误",err01)
			continue
		}
		f.WriteString(information)
		f.Close()
	}
	fmt.Println("ok")
}
func main() {
	var start, end int
	fmt.Println("请输入起始页：")
	fmt.Scan(&start)
	fmt.Println("请输入中止页：")
	fmt.Scan(&end)
	Select(start, end)
}
