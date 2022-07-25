package examples

import "fmt"

func Defer() {
	var n int

	f1 := func() {
		fmt.Println("def: ", n)
		n++
	}

	f := func() {
		defer f1()
		fmt.Println("var: ", n)
	}

	defer f()
	// 0 0
}
