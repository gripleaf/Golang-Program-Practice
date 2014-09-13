// First project First.go
package main

import (
	"fmt"
	"math"
	"os"
)

func ParseInt(bs string) int {
	res := 0
	for _, item := range bs {
		res = res * 10
		switch item {
		case '0':
			res += 0
			break
		case '1':
			res += 1
			break
		case '2':
			res += 2
			break
		case '3':
			res += 3
			break
		case '4':
			res += 4
			break
		case '5':
			res += 5
			break
		case '6':
			res += 6
			break
		case '7':
			res += 7
			break
		case '8':
			res += 8
			break
		case '9':
			res += 9
			break
		}
	}
	return res
}

func main() {
	cost := float64(ParseInt(os.Args[1]))
	need := 0.0
	need += math.Min(10*10000, cost*0.1)
	if cost > 10*10000 {
		need = need + math.Min(10*10000.0, cost-10*10000.0)*0.075
	}
	if cost > 20*10000 {
		need = need + math.Min(20*10000.0, cost-20*10000.0)*0.05
	}
	if cost > 40*10000 {
		need = need + math.Min(20*10000.0, cost-40*10000.0)*0.03
	}

	if cost > 60*10000 {
		need = need + math.Min(40*10000.0, cost-60*10000.0)*0.015
	}

	if cost > 100*10000 {
		need = need + (cost-100*10000.0)*0.01
	}
	fmt.Println(need)
}
