package main

import (
	"fmt"
	"runtime"
)

func main() {
	//创建一个goroutine
	go func(s string) {
		for i := 0; i < 12; i++ {
			fmt.Println(s)
		}
	}("world")

	for i := 0; i < 2; i++ {
		runtime.Gosched() //import "runtime"
		/*
		   屏蔽runtime.Gosched()运行结果如下：
		       hello
		       hello

		   没有runtime.Gosched()运行结果如下：
		       world
		       world
		       hello
		       hello
		*/
		fmt.Println("hello")
	}

}
