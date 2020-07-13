package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

//使用结构体
type Queue struct {
	Maxsize int
	Array   [5]int //模拟队列
	Front   int    //指向队列首部
	Rear    int    //指向队列尾部
}

func (this *Queue) AddQueue(val int) (err error) {
	fmt.Println("剩余",this.Maxsize-2-this.Rear,"个空位")
	var key1 string
	if this.Rear == this.Maxsize-1 {
		return errors.New("队列已满")
	}
	this.Array[this.Rear] = val
	//循环该函数，做到循环输入增加值
	this.Rear++ //后移
	fmt.Println("请继续输入需要添加的值")
	fmt.Println("输入exit结束添加")
	fmt.Scan(&key1)
	switch key1 {
	case "exit":
		return
	default:
		//val, errKey := strconv.ParseInt(key1, 10, 8)
		//if errKey != nil {
		//	fmt.Println("字节转换错误")
		//	fmt.Println(key1)
		//}
		val, errStrconv := strconv.Atoi(key1)
		if errStrconv != nil {
			fmt.Println("字节转换错误 ，err=", errStrconv)
		}
		if this.Rear == this.Maxsize-1 {
			fmt.Println("this.Rear=",this.Rear)
			return errors.New("队列已满 ")
		}
		if this.Rear == this.Maxsize-1{
			return errors.New("队列已满")
		}
		this.AddQueue(val)
	}
	return
}

//显示队列，找到队首，遍历队列
func (this *Queue) ShowQueue() {
	for i := this.Front + 1; i <= this.Rear; i++ {
		fmt.Printf("array[%d]=%d\n", i, this.Array[i])
	}
}
func main() {
	//创建一个队列
	a := &Queue{
		Maxsize: 5,
		Front:   -1,
		Rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入add\t表示添加数据队列")
		fmt.Println("2.输入get\t表示获取数据队列")
		fmt.Println("3.输入show\t表示显示数据队列")
		fmt.Println("4.输入exit\t表示显示队列")
		fmt.Scan(&key)
		switch key {
		case "add":
			fmt.Println("输入需要的队列数")
			fmt.Scan(&val)
			err := a.AddQueue(val)
			if err != nil {
				//fmt.Println(val)
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列完成")
			}
		case "get":
			fmt.Println("get")
		case "show":
			a.ShowQueue()
		case "exit":
			os.Exit(0)
		}

	}
}
