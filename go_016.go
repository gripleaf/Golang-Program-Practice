// First project First.go

package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func conv(a int) bool {
	return a != 0
}

func kgcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if conv(a&1) && conv(b&1) {
		return kgcd(a>>1, b>>1) << 1
	} else if conv(b & 1) {
		return kgcd(a>>1, b)
	} else if conv(a & 1) {
		return kgcd(a, b>>1)
	} else {
		return kgcd(abs(a-b), min(a, b))
	}
}

func main() {
	var n, m int
	fmt.Scanf("%d%d", &n, &m)
	var cc int = kgcd(n, m)
	fmt.Println("最大公约数是 ", cc)
	fmt.Println("最大公倍数是 ", n*m/cc)
}
