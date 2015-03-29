//对 10 个数进行排序

package main

import (
	"fmt"
)

func main() {
	var n [10]int
	for i := 0; i < 10; i++ { //输入
		fmt.Scan(&n[i])
	}
	for i := 0; i < 10-1; i++ { //冒泡排序法
		for j := i + 1; j < 10; j++ {
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
	fmt.Println(n)
}
