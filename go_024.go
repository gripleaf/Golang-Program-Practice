// First project First.go

package main

import (
	"fmt"
)

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func kgcd(x, y int) int {
	conv := func(x int) bool {
		return x != 0
	}
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	if conv(x&1) && conv(y&1) {
		return kgcd(abs(x-y), min(x, y))
	} else if conv(x & 1) {
		return kgcd(x, y>>1)
	} else if conv(y & 1) {
		return kgcd(x>>1, y)
	} else {
		return kgcd(x>>1, y>>1) << 1
	}
}

type fraction struct {
	fenzi, fenmu int
}

// 分数相加，同时进行通分处理【此处要小心分子乘分母并相加会超过int32】
func (x *fraction) add(y fraction) {
	x.fenzi = x.fenzi*y.fenmu + y.fenzi*x.fenmu
	x.fenmu = x.fenmu * y.fenmu
	cc := kgcd(x.fenzi, x.fenmu)
	x.fenzi /= cc
	x.fenmu /= cc
}

func (x fraction) value() float64 {
	return float64(x.fenzi) / float64(x.fenmu)
}

func main() {
	var n int
	x, y := 1, 2
	fmt.Scanf("%d", &n)
	sum := fraction{y, x}
	for i := 1; i < n; i++ {
		y, x = y+x, y
		sum.add(fraction{y, x})
	}
	fmt.Println(sum.fenzi, "/", sum.fenmu, "=", sum.value())
}
