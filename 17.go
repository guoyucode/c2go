//将一个数组逆序输出。

package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	lenth := len(s)
	for i := 0; i < lenth/2; i++ {
		s[i], s[lenth-i-1] = s[lenth-i-1], s[i]
	}
	fmt.Println(s)
}
