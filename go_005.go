// First project First.go
package main

import (
	"fmt"
)

const SIZE = 100

var tmp [SIZE]int

/* quick sort algorithm*/
func quick_sort(num_list *[SIZE]int, left, right int) {
	if right-left <= 1 {
		return
	}
	var left_inx int = left
	var right_inx int = right - 1
	for i := left + 1; i < right; i++ {
		if num_list[i] < num_list[left] {
			tmp[left_inx] = num_list[i]
			left_inx++
		} else {
			tmp[right_inx] = num_list[i]
			right_inx--
		}
	}
	tmp[left_inx] = num_list[left]
	for i := left; i < right; i++ {
		num_list[i] = tmp[i]
	}
	//fmt.Println(left, right, num_list)
	quick_sort(num_list, left, left_inx)
	quick_sort(num_list, left_inx+1, right)
}

/* please don't use 'enter',because fmt.scanf can not ignore it! */
func main() {
	var n int
	var num_list [SIZE]int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &num_list[i])
	}
	quick_sort(&num_list, 0, n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", num_list[i])
	}
	fmt.Println("")
}
