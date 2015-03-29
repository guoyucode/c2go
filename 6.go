//输入两个正整数 m 和 n,求其最大公约数和最小公倍数。
package main

import "fmt"

func main() {
	var num1, num2, multi int
	fmt.Println("Please input 2 numbers")
	fmt.Scanf("%d %d", &num1, &num2)
	multi = num1 * num2
	if num1 < num2 {
		num1, num2 = num2, num1
	}
	for num2 != 0 {
		num1, num2 = num2, num1%num2
	}
	fmt.Println("Greatest Common Divisor is ", num1)
	fmt.Println("lowest common multiple ", multi/num1)
}
