// First project First.go

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a, n int
	fmt.Scanf("%d%d", &a, &n)
	var sum int = 0
	var x int = a
	bits := int(math.Pow10(len(strconv.Itoa(a))))
	for i := 0; i < n; i++ {
		sum += x
		x = x*bits + a
	}
	fmt.Println(sum)
}
