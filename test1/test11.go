package main

import "fmt"

type CircleQueue struct {
	MaxSize int
	Array   [4]int
	Head    int
	tail    int
}

//加入元素
func (this *CircleQueue) Push(val int) (errPush error) {

}

//取出元素
func (this *CircleQueue) Pop() (val int, errPop error) {

}

//判断环形队列是否满栈
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.MaxSize == this.Head
}

//判断环形队列是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.Head
}

//取出环形队列有多少元素
func (this *CircleQueue) Size() int {
	return (this.tail + this.MaxSize - this.Head) % this.MaxSize
}
func main() {
	var key string
	var val int
	for {
		fmt.Println("1.输入add\t表示添加数据队列")
		fmt.Println("2.输入get\t表示获取数据队列")
		fmt.Println("3.输入show\t表示显示数据队列")
		fmt.Println("4.输入exit\t表示显示队列")
		fmt.Scan(&key)
		switch key {
		}
	}
}
