package examples

import (
	"fmt"
	"strings"
)

func modify1(arr *[3]int) {
	(*arr)[0] = 90
}

func modify2(sls []int) {
	sls[0] = 100
}

func TestArr1() {
	a := [3]int{89, 90, 91}
	modify1(&a)
	fmt.Println(a)
	fmt.Println(string([]byte("")))
	b := strings.Contains("", "")
	fmt.Println(b)
	// 90 90 91

	modify2(a[:])
	// 100 90 91
	fmt.Println(a)
}
