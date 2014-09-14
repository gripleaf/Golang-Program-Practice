// First project First.go

package main

import (
	"fmt"
)

const SIZE = 10000

var mark [SIZE]bool

func main() {
	var n int
	fmt.Scanf("%d", &n)
	// switch case
	switch {
	case n >= 90 && n <= 100:
		fmt.Println("A")
		break
	case n < 60:
		fmt.Println("C")
		break
	case n < 90:
		fmt.Println("B")
		break

	default:
		fmt.Println("Unknow Score!")
	}

	// a <operator> b ? process1 : process2 ---> which can use in C
	//n<=90?(n<=60?fmt.Println("C"):fmt.Println("B")):(n>100?fmt.Println("Unknow Score!"):fmt.Println("A"))
}
