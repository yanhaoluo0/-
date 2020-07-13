package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
)

var name string

func main() {
	conn1, _ := redis.Dial("tcp", "127.0.0.1:6379")
	fmt.Println("1,", conn1)
	fmt.Println("输入name")
	_, err1 := fmt.Scan(&name)

	if err1 != nil {
		fmt.Println("&name error--------", err1)
	}
	m0, n0 := redis.String(conn1.Do("Hmset", name, "age", "30"))
	if n0 != nil {
		fmt.Println("n0 err----------", n0)
	}
	fmt.Println("m0 ----------", m0)
	//m1, n1 := redis.String(redis.Values(conn1.Do("Hgetall", name)))
	m1, n1 := redis.StringMap(conn1.Do("Hgetall", name))
	fmt.Println("打印类型如下")
	fmt.Println(reflect.TypeOf(m1))
	if n1 != nil {
		fmt.Println("n1 err----------", n1)
	}
	fmt.Println("m1=", m1, "\nn1", n1)

}
