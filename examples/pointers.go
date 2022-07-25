package examples

import "fmt"

func ChangePointer1(b *int) {
	a := 3
	// b это не ссылка а значение в виде указателя, b сверху не изменится
	b = &a
}

func ChangePointer2(b *int) {
	a := 3
	// идет обращение к участку паямти, исходная ячейка памяти перезатрется
	*b = a
}

type dur int

func (d dur) update() {
	// нет указателя на dur, значение выше не изменится
	d = 5
}

func TestPointer1() {
	d := dur(42)
	d.update()

	fmt.Println(d)
}
