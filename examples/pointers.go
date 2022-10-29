package examples

import "fmt"

// b это не ссылка а значение в виде указателя, b сверху не изменится
func ChangePointer1(b *int) {
	a := 3
	b = &a
}

// идет обращение к участку паямти, исходная ячейка памяти перезатрется
func ChangePointer2(b *int) {
	a := 3
	*b = a
}

type dur int

func (d dur) update() {
	d = 5
}

// нет указателя на dur, значение выше не изменится
func TestPointer1() {
	d := dur(42)
	d.update()

	fmt.Println(d)
}

// 4 4 4 4
func TestPointer4() {
	a := []int{1, 2, 3, 4}
	result := make([]*int, len(a))
	for i, v := range a {
		result[i] = &v
	}
	fmt.Println(result)
}
