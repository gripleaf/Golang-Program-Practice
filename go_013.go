// First project First.go

package main

import (
	"fmt"
)

func main() {
	for i := 100; i < 1000; i++ {
		var st int = (i/100)*(i/100)*(i/100) + (i%10)*(i%10)*(i%10) + ((i/10)%10)*((i/10)%10)*((i/10)%10)
		//fmt.Println(st, i)
		if st == i {
			fmt.Println(st)
		}
	}
	fmt.Println("")
}
