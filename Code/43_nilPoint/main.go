package main

import (
	"fmt"
	"strings"
)

type test struct {
	name string
	m    map[int]string
	num  int64
}

type Age struct {
	a int
}

func main() {
	//r := getT(1)
	//
	//if r.age.a == 10 {
	//	fmt.Println("yes")
	//}

	//a := (int64(0) >> uint(2)) & int64(1)
	//fmt.Println(a)
	//
	//if a == int64(1) {
	//	fmt.Println(111)
	//}
	//
	//aaa := &test{name: ""}
	//if aaa == nil {
	//	fmt.Println("111")
	//}

	res := make(map[int]*test, 5)
	info, ok := res[0]
	if !ok {
		tmp := make(map[int]string)
		info = &test{
			name: "123",
			m:    tmp,
			num:  0,
		}
		res[0] = info
	}

	info.m[0] = "test"

	fmt.Println(res[0].m)

	keys := []string{"test1", "test2"}

	var (
		sqls []string
		args []interface{}
	)
	for _, k := range keys {
		sqls = append(sqls, "?")
		args = append(args, k)
	}
	a := fmt.Sprintf("SELECT aid, biz_type, biz_value, is_deleted FROM archive_extra_biz WHERE aid=? AND biz_type IN (%s)", strings.Join(sqls, ","))
	fmt.Println(a)
	fmt.Println(args)
}

//func getT(num int) (reply *t) {
//	tmp := &t{name: "111", age: &Age{a: 10}}
//
//	if num == 1 {
//		return
//	}
//
//	reply = tmp
//	return
//}
