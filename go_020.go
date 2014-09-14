// First project First.go

package main

import (
	"fmt"
)

func main() {
	s := 100.0
	h := s / 2
	for i := 2; i <= 10; i++ {
		s += 2 * h
		h /= 2
	}
	fmt.Printf("the total of road is %f\n", s)
	fmt.Printf("the tenth is %f meter\n", h)
}
