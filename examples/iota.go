package examples

import "fmt"

func TestIota() {
	const (
		// A1 старт с 0
		A1 = 1 << iota // сдвиг бита влево на 0
		A2             // сдвиг влево на 1
		A3             // сдвиг влево на 2
		A4             // сдвиг влево на 3
		A5             // сдвиг влево на 4
	)
	fmt.Println(A1, A2, A3, A4, A5)
	// 1 2 4 8 16
	// вспоминаем двоичную систему счисления и сдвиги
}
