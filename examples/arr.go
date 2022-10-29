package examples

import (
	"fmt"
)

func modify1(arr *[3]int) {
	(*arr)[0] = 90
}

func modify2(sls []int) {
	sls[0] = 100
}

// 90 90 91
// 100 90 91
func TestArr1() {
	a := [3]int{89, 90, 91}
	modify1(&a)
	fmt.Println(a)

	modify2(a[:])
	fmt.Println(a)
}
