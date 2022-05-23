package main

import (
	"fmt"
	"hash/crc32"
	"strings"
)

type aaa struct {
	sections []*section
}

type section struct {
	age int
}

func main() {
OutLoop:
	for i := 0; i < 100; i++ {
		for j := 1000; j < 2000; j++ {
			switch j {
			case 1005:
				break OutLoop
			}
			fmt.Println(i)
			fmt.Println(j)
		}

	}

	//sec := section{age: 20}
	//
	//test := aaa{sections: []*section{&sec}}
	//
	//fmt.Printf("%+v", test.sections)
	//
	//test.edit()
	//
	//for _, i := range test.sections {
	//	fmt.Println(i.age)
	//}
	//fmt.Println("========================")
	//fmt.Println(test)

	//buvid := "22"
	//a := ClarityGrayControl(false, buvid)
	//fmt.Println(a)
	//test := &aaa{
	//	name: "lzw",
	//	age:  22,
	//}
	//fmt.Println(test)
	//change(test)
	//fmt.Println(test)
	//a := trans("This is a sample",16)
	//fmt.Println(a)
	//l := list.New()
	//for i:=1;i<=6;i++ {
	//	l.PushBack(i)
	//}
	//
	//l2 := list.New()
	//
	//cur := l.Front()
	//for cur != nil {
	//	tmp := cur.Next()
	//	l2.PushBack(tmp.Value)
	//
	//	l.Remove(cur.Next())
	//	cur = cur.Next()
	//}
	//
	//l.PushBackList(l2)
	//
	//
	//for e:=l.Front();e!=nil;e=e.Next(){
	//	fmt.Println(e.Value)
	//}

	//cpunum := flag.Int("cpunum", 0, "cpunum")
	//flag.Parse()
	//fmt.Println("使用的cpunum:", *cpunum)
	//println("系统的CPU核数",runtime.NumCPU())
	//runtime.GOMAXPROCS(*cpunum)
	//for i := 0; i < *cpunum - 1; i++ {
	//	go func() {
	//		for {
	//
	//		}
	//	}()
	//}
	//for {
	//
	//}

	//a := []int{-1,5,4,6,8,10,5,3,2,1}
	//sort.Ints(a)
	//fmt.Println(a)
	//
	//b := a[:1]
	//c := a[:2]
	//fmt.Println(b)
	//fmt.Println(c)
	//
	//c[0] = 100
	//
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//
	//fmt.Printf("bool align: %d\n", unsafe.Alignof(bool(true)))
	//fmt.Printf("int32 align: %d\n", unsafe.Alignof(int32(0)))
	//fmt.Printf("int8 align: %d\n", unsafe.Alignof(int8(0)))
	//fmt.Printf("int64 align: %d\n", unsafe.Alignof(int64(0)))
	//fmt.Printf("byte align: %d\n", unsafe.Alignof(byte(0)))
	//fmt.Printf("string align: %d\n", unsafe.Alignof("EDDYCJY"))
	//fmt.Printf("map align: %d\n", unsafe.Alignof(map[string]string{}))

	//b := test{
	//	1,
	//}
	//c := b.p()
	//fmt.Println(c)
	//a := 1
	//reflect.TypeOf(a)

}

func (v *aaa) edit() {

	for _, s := range v.sections {
		s.age = 100009
	}

	fmt.Println(v)
}

type T interface {
	p()
}

func trans(s string, n int) string {
	stack := make([]string, n)
	num := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			num++
		} else {
			if s[i] >= 'a' && s[i] <= 'z' {
				stack[num] = fmt.Sprintf("%s%s", stack[num], strings.ToUpper(string(s[i])))
			} else {
				stack[num] = fmt.Sprintf("%s%s", stack[num], strings.ToLower(string(s[i])))
			}

		}

	}
	stack = stack[:num+1]
	reverse(&stack)
	res := strings.Join(stack, " ")

	return res
}

func reverse(s *[]string) {
	length := len(*s)
	for i := 0; i < length/2; i++ {
		(*s)[i], (*s)[length-i-1] = (*s)[length-i-1], (*s)[i]
	}

}

//type test

//package main
//
//type T interface {
//	Foo()
//}
//
//type S struct{}
//
//func (s *S) Foo() {}
//
//func main() {
//	s := new(S)
//	T(s).Foo()
//}

// ClarityGrayControl 白名单和灰度控制
func ClarityGrayControl(ok bool, buvid string) bool {
	// 白名单
	// 灰度控制
	group := crc32.ChecksumIEEE([]byte(buvid)) % 1000
	fmt.Println(group)
	return ok || group < uint32(1)
}
