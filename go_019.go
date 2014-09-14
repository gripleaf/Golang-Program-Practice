// First project First.go

package main

import (
	"fmt"
	"math"
)

func getSumofitem(n int) int {
	var sum int = 1
	var st int = int(math.Sqrt(float64(n)))
	for i := 2; i < st; i++ {
		if n%i == 0 {
			sum += i
			sum += n / i
		}
	}
	if st != 1 && n%st == 0 {
		sum += st
	}
	return sum
}

func main() {
	for i := 1; i < 1000; i++ {
		if i == getSumofitem(i) {
			fmt.Println(i)
		}
	}
}
