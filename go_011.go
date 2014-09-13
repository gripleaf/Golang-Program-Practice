// First project First.go

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	if n <= 2 {
		fmt.Println("the total num is 1")
	} else {
		var one int = 1
		var two int = 1
		var ant int
		for i := 2; i < n; i++ {
			ant = one + two
			one = two
			two = ant
		}
		fmt.Println("the total num is", ant)
	}
}
