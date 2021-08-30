package main

import "fmt"

func idxOf(str string, bank []string) int {		//判断是否找到目标值
	for i, s := range bank {					//i代表角标，s代表结果值，str为目标值，通过角标返回的数，最后确定得到的最终结果
		if str == s {
			return i
		}
	}
	return -1
}

func hasOneDiff(x, y string) bool {				//判断是否有环路？？	，x用以表示queue的长度，y表示字典的值
	count := 0
	for i := 0; i < len(x); i++ {				//找寻在队列内的存在可能
		if x[i] != y[i] {						//若队列的值和字典值的相同角标下，目标值不相同，则继续遍历
			count++
		}
		if count > 1 {						  	 //但若count使用一次后，依然不存在
			return false
		}
	}
	if count == 1 {								//若
		return true
	}
	return false
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if idxOf(endWord, wordList) == -1 {		//查看目标值是否在待列表里
		return 0
	}
	step := 0								//
	used := make([]bool, len(wordList))		//制作切片used，用以判断是否有值，长度等于需要判断的字典
	queue := []string{beginWord}			//新建一个列表，并将起始值赋给它
	for len(queue) > 0 {					//遍历，列表是否还有值？
		step++
		l := len(queue)							//将列表长度赋给L
		for i := 0; i < l; i++ {				//遍历，L内所有的值
			if queue[i] == endWord {			//若，链表内值与目标值相同，则停止
				return step
			}
			for j, w := range wordList {		//若不同，二次遍历，将字典内角标与值相应赋给j、w
				if !used[j] && hasOneDiff(queue[i], w) {		//若制作的used的角标不为空，并且
					queue = append(queue, w)	//在数组后，添加	
					used[j] = true
				}
			}
		}
		queue = queue[l:]
	}
	return 0
}
func main() {
	a :=[] string{"hot","dot","dog","lot","log","cog"}
	p:=ladderLength("hit","cog",a)
	fmt.Println(p)
}
