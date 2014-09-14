// First project First.go

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	// letter number
	var lt_cnt int = 0
	// space number
	var sp_cnt int = 0
	// number number
	var num_cnt int = 0
	// others number
	var ot_cnt int = 0
	// 汉字个数
	var 汉字_cnt int = 0
	fmt.Println("|", str, "|")
	for _, cc := range str {
		fmt.Println(cc)
		if cc >= 'a' && cc <= 'z' || cc >= 'A' && cc <= 'Z' {
			lt_cnt++
		} else if cc >= '0' && cc <= '9' {
			num_cnt++
		} else if cc == ' ' {
			sp_cnt++
		} else if cc >= ox4e00 && cc < ox9fa5 {
			汉字_cnt++
		} else {
			ot_cnt++
		}
	}
	fmt.Println("letter: ", lt_cnt, " space: ", sp_cnt, " number: ", num_cnt, " others: ", ot_cnt, "汉字: ", 汉字_cnt)
}
