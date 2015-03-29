//打印出如下图案(菱形)
//*
//***
//******
//********
//******
//***
//*

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 2*i+1; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	for i := 3; i > 0; i-- {
		for j := 2*i - 1; j > 0; j-- {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
