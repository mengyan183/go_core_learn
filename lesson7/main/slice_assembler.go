package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := make([]int, 1)
	fmt.Println(unsafe.Pointer(&a))
	b := append(a, 1, 2)
	fmt.Println(unsafe.Pointer(&b))
}
