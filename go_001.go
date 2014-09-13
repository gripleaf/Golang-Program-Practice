// First project First.go
package main

import (
	"fmt"
	"strconv"
)

var mp = make(map[string]bool)
var cnt int
var enum = []int{1, 2, 3, 4}
var limit int

func judge(x, y string) bool {
	for it, _ := range y {
		if int(y[it]) == int(x[0]) {
			return false
		}
	}
	return true
}

func dfs(res string) {
	if len(res) == limit {
		if res[0] == '0' {
			return
		}
		if mp[res] == false {
			cnt++
			mp[res] = true
			fmt.Println(res)
		}
		return
	}
	for _, val := range enum {
		if judge(strconv.Itoa(val), res) {
			dfs(strconv.Itoa(val) + res)
		}
	}
}

func main() {
	limit = 3
	cnt = 0
	for _, val := range enum {
		dfs(strconv.Itoa(val))
	}
	fmt.Printf("total num is %d\n", cnt)
}
