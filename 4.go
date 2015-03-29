// 将一个正整数分解质因数。例如:输入 90,打印出 90=2*3*3*5。
package main

import (
	"fmt"
)

func main() {
	var n, i int
	fmt.Printf("\nPlease input a number:\n")
	fmt.Scanf("%d", &n)
	fmt.Printf("%d=", n)
	for i = 2; i <= n; i++ {
		for n != i {
			if n%i == 0 {
				fmt.Printf("%d*", i)
				n = n / i
			} else {
				break
			}
		}
	}
	fmt.Printf("%d\n", n)
}
