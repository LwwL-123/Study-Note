package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(solve("123","127"))


}
func solve( s string ,  t string ) string {
	len1,len2 := len(s)-1,len(t)-1
	ifAdd := 0
	res := ""
	for len1 >= 0 || len2 >= 0 || ifAdd == 1 {
		sum := 0
		if len1 >= 0 {
			sum += int(s[len1]-'0')
			len1--
		}

		if len2 >= 0 {
			sum += int(t[len2]-'0')
			len2--
		}

		if ifAdd != 0 {
			sum += 1
			ifAdd = 0
		}

		if sum >= 10 {
			sum = sum - 10
			ifAdd = 1
		}

		res = strconv.Itoa(sum) + res
	}

	return res
}