// First project First.go
package main

import (
	"fmt"
)

var mons = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func isLeap(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func main() {
	var year, month, day int
	var ant int = 0
	fmt.Scanf("%d%d%d", &year, &month, &day)
	if isLeap(year) {
		mons[1] = 29
	} else {
		mons[1] = 28
	}
	for i := 1; i < month; i++ {
		ant += mons[i-1]
	}
	ant += day
	fmt.Printf("this day is %d\n", ant)
}
