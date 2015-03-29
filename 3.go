// 打印出所有的“水仙花数”,所谓“水仙花数”是指一个三位数,其各位
// 数字立方和等于该数本身。例如:153 是一个“水仙花数”,因为 153=1 的三次方
// +5 的三次方+3 的三次方。
package main

import (
	"fmt"
)

func main() {
	for i := 100; i < 999; i++ {
		baiwei := i / 100
		shiwei := (i - baiwei*100) / 10
		gewei := i % 10
		if i == baiwei*baiwei*baiwei+shiwei*shiwei*shiwei+gewei*gewei*gewei {
			fmt.Println(i)
		}
	}
}
