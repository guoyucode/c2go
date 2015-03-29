//一球从 100 米高度自由落下,每次落地后反跳回原高度的一半;再落下,
//求它在第 10 次落地时,共经过多少米?第 10 次反弹多高?
package main

import "fmt"

const COUNT = 10

func main() {
	var high, sum float64 = 100, -100
	for i := 0; i < COUNT; i++ {
		sum += 2 * high
		high /= 2
	}
	fmt.Println("The total of road is ", sum)
	fmt.Println("The tenth is ", high, " meter")
}
