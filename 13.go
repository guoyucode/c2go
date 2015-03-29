//一个 5 位数,判断它是不是回文数。即 12321 是回文数,个位与万位相同,
//十位与千位相同。

package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Println("请输入一个5位数:")
	fmt.Scanln(&str)
	input := []byte(str)
	if input[0] == input[4] && input[1] == input[3] {
		fmt.Println("这是回文数！")
	} else {
		fmt.Println("这并不是回文数哦。")
	}
}
