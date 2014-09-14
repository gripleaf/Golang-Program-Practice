// First project First.go

package main

import (
	"fmt"
)

func NumFact(x int) int {
	if x == 1 {
		return 1
	}
	return NumFact(x-1) * x
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(NumFact(n))
}
