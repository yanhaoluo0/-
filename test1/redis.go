package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	//"github.com/garyburd/redigo/redis"

)

func main() {
	conn1, err1 := redis.Dial("tcp", "127.0.0.1:6379")
	if err1 != nil {
		fmt.Println("错误conn = ", conn1)
		fmt.Println("错误err = ", err1)
	}
	//defer close
	fmt.Println("conn = ", conn1)
	fmt.Println("err = ", err1)
	coon, err := conn1.Do("Set", "name", "tom")
	if err != nil {
		fmt.Println("set err = ", err)
		return
	}
	fmt.Println("finish")
	fmt.Println("添加结果conn = ", coon)
	fmt.Println("返回结果err = ", err)
	n1, err := redis.String(conn1.Do("Get", "name"))
	fmt.Println("添加结果conn = ", n1)
	fmt.Println("返回结果err = ", err)
	conn1.Close()
	fmt.Println("------------------------------------------")
	var pool01 *redis.Pool
	pool01 = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //最大连接数目，0表示不设置上限
		IdleTimeout: 100, //最大空闲时间	单位S
		Dial: func() (redis.Conn, error) { //初始化连接，需要连接具体哪个ip
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	c := pool01.Get() //获取连接池pool01里的一个连接
	conn2, err3 := redis.String(c.Do("set", "name", "tom"))
	fmt.Println()
	fmt.Printf("conn2： \n    %v", conn2)
	if err3 !=nil{
		fmt.Println("error=",err3)
	}
	defer c.Close() //关闭连接池pool01，关闭后将不能获取
}
