// First project First.go
package main

import (
	"fmt"
	"math"
)

func isSqrt(x int) bool {
	var cc float64 = math.Sqrt(float64(x))

	if int(cc)*int(cc) == x {
		return true
	}
	return false
}

func judge(x int) bool {
	if isSqrt(x+100) && isSqrt(x+268) {
		return true
	}
	return false
}

func main() {
	fmt.Println("the answer is:")
	for i := 0; i < 10000; i++ {
		if judge(i) {
			fmt.Println(i)
		}
	}
}
