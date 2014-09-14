// First project First.go

package main

import (
	"fmt"
)

func main() {
	var x1, x2, day int = 0, 1, 9
	for day > 0 {
		x1 = (x2 + 1) * 2
		x2 = x1
		day--
	}
	fmt.Printf("the total is %d\n", x1)
}
