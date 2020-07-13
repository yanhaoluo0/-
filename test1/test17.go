package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	i := 0
	file, err := os.OpenFile("./TestLog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("文件写入错误：",err)
	}
	log.SetOutput(file)
	for {
		log.Println("test ", i, " 完成输入")
		//time.Sleep(time.Second)
		i++
		if i == 10 {
			break
		}
	}

	defer fmt.Println("8888:",nil)
}
