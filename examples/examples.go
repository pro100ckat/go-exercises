package examples

import (
	"fmt"
)

func Example1() {
	for i, bar := 0, foo1(); i < 3; i++ {
		fmt.Println(bar())
	}
}

func foo1() func() int {
	var n int
	return func() int {
		n += 3
		return n
	}
}
