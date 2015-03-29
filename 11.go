//一只猴子摘了 N 个桃子第一天吃了一半又多吃了一个,第二天又吃了余下的
//一半又多吃了一个,到第十天的时候发现还有一个.
package main

import "fmt"

func main() {
	fmt.Println("N =", peach(10))
}

func peach(n int) int {
	if n == 1 {
		return 1
	} else {
		return (peach(n-1) + 1) * 2
	}
}
