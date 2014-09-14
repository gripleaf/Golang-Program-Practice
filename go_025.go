// First project First.go

package main

import (
	"fmt"
)

func NumFact(x int) int {
	sum := 1
	for i := 2; i <= x; i++ {
		sum *= i
	}
	return sum
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += NumFact(i)
	}
	fmt.Println(sum)
}
