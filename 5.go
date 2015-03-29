// 学习成绩>=90 分的同学用 A 表示,60
// -89 分之间的用 B 表示,60 分以下的用 C 表示。
package main

import (
	"fmt"
)

func main() {
	var score int
	fmt.Println("Please input the score:")
	fmt.Scanln(&score)
	if score >= 90 {
		fmt.Println("A")
	} else if score < 60 {
		fmt.Println("C")
	} else {
		fmt.Println("B")
	}
}
