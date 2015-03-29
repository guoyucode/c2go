//输入一行字符,分别统计出其中英文字母、空格、数字和其它字符的个数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	letters, space, digit, others := 0, 0, 0, 0
	fmt.Println("Please input some characters:")
	inputReader := bufio.NewReader(os.Stdin)
	str, _ := inputReader.ReadString('\n')
	fmt.Println(str)
	s := []byte(str)
	for _, n := range s {
		if n >= byte('a') && n <= byte('z') {
			letters += 1
		} else if n >= byte('0') && n <= byte('9') {
			digit += 1
		} else if n == byte(' ') {
			space += 1
		} else {
			others += 1
		}
	}
	fmt.Println("letters:", letters, "\tspace:", space, "\tdigit:", digit, "\tothers:", others)
}
