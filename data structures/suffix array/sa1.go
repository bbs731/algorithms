package main

import (
	"fmt"
	"index/suffixarray"
	"reflect"
	"unsafe"
)

/*
直接使用标准库里面 suffix array 的实现！
 */
func main() {

	s := "vamamadn"
	n := len(s)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s))).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}

	fmt.Println(sa)
	fmt.Println(rank)

}
