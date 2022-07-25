package examples

import "fmt"

func TestString1() {
	a := []byte{'a', 'b', 'c'}
	b := string(a)
	a[0] = 'z'
	fmt.Printf("%s\n", b)
	// a b c
}
