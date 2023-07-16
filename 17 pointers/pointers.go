package main

import "fmt"

func zeroVal(ival int) {
	fmt.Println("before:", ival)
	ival = 0
	fmt.Println("after:", ival)
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroVal(i)
	fmt.Println("zeroval:", i)

	zeroPtr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
