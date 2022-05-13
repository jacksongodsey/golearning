package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 16; i++ {
		for j := 1; j < 16; j++ {
			sum := i * j
			fmt.Printf("%4d ", sum)
		}
		fmt.Println()
	}
}
