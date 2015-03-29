//求 s=a+aa+aaa+aaaa+aa...a 的值,其中 a 是一个数字。例如
package main

import "fmt"

func main() {

	var a, n, s, sum int
	fmt.Println("Please input a and n :")
	fmt.Scanln(&a, &n)
	for i := 0; i <= n; i++ {
		s = 10*s + a
		sum += s
	}
	fmt.Println("sum=", sum)
}
