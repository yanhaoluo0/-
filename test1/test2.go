package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn ,err:=redis.Dial("tcp" ,"127.0.0.1:6379")
	if err!=nil{
		fmt.Println("错误conn = ",conn)
		fmt.Println("错误err = ",err)
	}
	fmt.Println("conn = ",conn)
	fmt.Println("err = ",err)
	fmt.Println("-----------")
}
