package main		//客户端

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//loop := true
	count := 1
	//a:=net.Listener().Accept()
	//fmt.Println("a :" , a)
	for {
		conn, err := net.Dial("tcp", "127.0.0.1:8888")
		if err != nil {
			fmt.Println("client dial err=", err)
			return
		}
		if count == 1 {
			fmt.Println("连接已建立，请输入需要发送的信息")
			count++
		}
		reader := bufio.NewReader(os.Stdin) //os.stdin是终端的标准输入方式
		//reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err==", err)
		}
		//fmt.Println("conn success=", conn, "客户端为：", conn.RemoteAddr())
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err ==", err)
		}
		fmt.Printf("发送了了%d个字符", n)
		fmt.Println()
	}
}
