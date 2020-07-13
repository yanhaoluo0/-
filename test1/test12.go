package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func Creat(value int) *TreeNode{
	return &TreeNode{value,nil,nil}
}
func main() {
	//var i = []int{3,3,5,0,0,3,1,4}
	//r:=TreeNode{3,nil,nil}
	r:=Creat(5)
	r.Left=Creat(3)
	r.Right=Creat(4)
	fmt.Println("*********",r)
	//r.Right=TreeNode{5,nil,nil}
	//r.Left=TreeNode{4,nil,nil}
	//*r=TreeNode{5,nil,nil}
	//i:=TreeNode{Val: 1,}
	//a := levelOrder(i)
	a:=levelOrder(r)
	fmt.Println(a)
}
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}