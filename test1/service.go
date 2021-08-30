//package test		服务端
package main

import (
	"fmt"
	"net"
)

var count = 1

func functionInput1() {
	fmt.Println("functionInput1")
	//conn1, err1 := redis.Dial("tcp", "127.0.0.1:6379")
	//defer conn1.Close()
	//if err1 != nil {
	//	fmt.Println("错误conn = ", conn1)
	//	fmt.Println("错误err = ", err1)
	//}else {
	//	coon, err := conn1.Do("Get", "name", )
	//
	//}
}
func functionInput2() {
	fmt.Println("functionInput2")
}
func functionInput3() {
	fmt.Println("functionInput3")
}
func functionInput4() {
	fmt.Println("functionInput4")
	//fmt.Println("新建用户指令")
	//conn1 ,err1:=redis.Dial("tcp" ,"127.0.0.1:6379 ")
	//defer conn1.Close()
	//if err1 !=nil{
	//	fmt.Println("Redis连接错误 =" , err1)
	//}
	//conn1 , err2  := conn1.Do("Set" , "name" ,"tom")
	//defer conn1.Close()
	//if err2 !=nil{
	//	fmt.Println("Redis连接错误 =" , err1)
	//}
}
func functionInputErr() {

}
func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		//fmt.Println("服务器在等待客户端发送的信息:" + conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器端的read err =", err)
			return
		}
		fmt.Println(conn.RemoteAddr().String(), " : ", string(buf[:4]), "输入完毕")
		//fmt.Println(reflect.TypeOf(string(buf)))
		//a1, erra1 := strconv.ParseFloat(string(buf))
		//if erra1!=nil{
		//	fmt.Println("类型转换出错P28 err=" , erra1[:n]))
		//}
		//int64(string(buf[:n]))
		//a1 := string(buf[:n])
		//a2 :=a1[:3]
		//fmt.Printf("%v" , a2)
		switch string(buf[:n]) {
		case "1\n":
			fmt.Println("进入1")
			functionInput1()
		case "2":
			fmt.Println("进入2")
			functionInput2()
		case "3\n":
			fmt.Println("进入3")
			functionInput3()
		case "4\n":
			fmt.Println("进入4")
			functionInput4()
		default:
			fmt.Println("未匹配到合适的值")
			functionInputErr()
		}
	}
}
func main() {
	fmt.Println("Starting connect....")
	listen, err := net.Listen("tcp", "0.0.0.0:8888") //使用tcp协议，表示本地监听0.0.0.0:8080
	if err != nil {
		fmt.Println("listen err=//  ", err, " /////err local =26")
		//fmt.Fprint(listen, "get/HTTP/1.0\r\n\r\n")
		//status ,err:=bufio.NewReader(listen).ReadString('\n')
		return
	}
	defer listen.Close() //defer，反向第一位操作
	for { //循环等待链接
		if count == 1 {
			fmt.Println("等待连接中....")
		}
		conn, err := listen.Accept() //监听端口的接收，由于只执行一次，因此写在for循环里循环等待连接
		if err != nil {
			fmt.Println("Accept err", err)
			return
		} else {
			if count == 1 {
				fmt.Printf("连接已建立，连接地址为=%v 客户端为ip=%v \n", conn, conn.RemoteAddr())
				count++
			}
		}
		go process(conn) //连接的输入，使用go进行强行跳转，否则只输入一次后消失
		//准备协程，为后续客户端服务
	}
	//fmt.Println()
	//fmt.Printf("listen suc=%v", listen)
}
