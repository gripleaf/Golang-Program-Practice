// First project First.go

package main

import (
	"fmt"
)

func main() {
	var i, j, k int
	for i = 0; i <= 3; i++ {
		for j = 0; j <= 2-i; j++ {
			fmt.Printf(" ")
		}
		for k = 0; k <= 2*i; k++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	for i = 0; i <= 2; i++ {
		for j = 0; j <= i; j++ {
			fmt.Printf(" ")
		}
		for k = 0; k <= 4-2*i; k++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
