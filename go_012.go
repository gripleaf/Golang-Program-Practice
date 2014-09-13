// First project First.go

package main

import (
	"fmt"
)

const SIZE = 10000

var mark [SIZE]bool

func main() {
	var n int
	list := make([]int, 0)
	for i := 2; i < SIZE; i++ {
		if mark[i] == false {
			list = append(list, i)
			for ct := i + 1; ct*i < SIZE; ct++ {
				mark[ct*i] = true
			}
		}
	}
	fmt.Scanf("%d", &n)
	for _, it := range list {
		if it > n {
			break
		} else {
			fmt.Println(it)
		}
	}
}
