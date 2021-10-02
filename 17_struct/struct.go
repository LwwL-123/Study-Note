package main

import "reflect"

type AAA struct {
	id   int64
	name []byte
	b    map[int]bool
}

func main() {
	a1 := AAA{id: 1, name: []byte("123"),b: map[int]bool{1:true}}
	a2 := AAA{id: 1, name: []byte("123"),b: map[int]bool{1:false}}

	if reflect.DeepEqual(a1, a2) {
		println("same")
	} else {
		println("not same")
	}

}
