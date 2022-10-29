package examples

import (
	"fmt"
	"strings"
)

// a b c
func TestString1() {
	a := []byte{'a', 'b', 'c'}
	b := string(a)
	a[0] = 'z'
	fmt.Printf("%s\n", b)
}

// true
func TestString2() {
	fmt.Println(string([]byte("")))
	b := strings.Contains("", "")
	fmt.Println(b)
}
