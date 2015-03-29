//求 100 之内的素数

package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 2; i <= 100; i++ {
		prime, mid := true, int(math.Sqrt(float64(i)))
		for j := 2; j <= mid; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			fmt.Println(i)
		}
	}
}
