// First project First.go

package main

import (
	"fmt"
)

func main() {
	var i, j, k int16 /*i 是 a 的对手，j 是 b 的对手，k 是 c 的对手*/
	for i = 'x'; i <= 'z'; i++ {
		for j = 'x'; j <= 'z'; j++ {
			if i != j {
				for k = 'x'; k <= 'z'; k++ {
					if i != k && j != k {
						if i != 'x' && k != 'x' && k != 'z' {
							fmt.Printf("order is a--%c b--%c c--%c\n", i, j, k)
						}
					}
				}
			}
		}
	}
}
