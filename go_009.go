// First project First.go

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				fmt.Printf("■")
			} else {
				fmt.Printf("□")
			}
		}
		fmt.Printf("\n")
	}
}
