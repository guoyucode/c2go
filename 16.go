//有一个已经排好序的数组。现输入一个数,要求按原来的规律将它插入数
//组中。

//简化了一下，仅保留功能。。。

package main

import (
	"fmt"
)

const N = 5

func main() {
	s1 := make([]int, 10)
	s2 := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}
	copy(s1, s2)
	for i := 8; i >= 0; i-- {
		if N <= s1[i] && N >= s1[i-1] {
			s1[i+1], s1[i] = s1[i], N
			break //插入后退出
		} else {
			s1[i+1] = s1[i]
		}
	}
	fmt.Println(s1)
}
