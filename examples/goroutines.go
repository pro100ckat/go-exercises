package examples

import (
	"fmt"
	"runtime"
	"time"
)

func TestGoroutines1() {
	var ch = make(chan int)
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}

	fmt.Println("result:", <-ch)
	//time.Sleep(2 * time.Second)
}

// что мы увидим в stdout?
func TestGoroutines2() {
	ch := make(chan string, 0)
	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2"

	// рандом
}

func TestGoroutines3() {
	data := []string{"one", "two", "three"}
	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(3 * time.Second)
	// three three three
	// потому цикл пробежит быстрее чем запустятся горутины, и при этом горутины смотрят на одну область памяти яйчейки памяти  v.
	// Loop variables captured by 'func' literals in 'go' statements might have unexpected values. Подсказка IDE.
}

func TestGoroutines4() {
	var num int

	for i := 0; i < 10000; i++ {
		go func(i int) {
			num = i
		}(i)
	}
	fmt.Printf("NUM is %d", num)
	// любое число от 0 до 9999
}

func TestGoroutines5() {
	dataMap := make(map[string]int)
	for i := 0; i < 10000; i++ {
		go func(d map[string]int, num int) {
			d[fmt.Sprintf("%d", num)] = num
		}(dataMap, i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(len(dataMap))
	// нельзя так писать в map
}

func TestGoroutines6() {
	runtime.GOMAXPROCS(1)

	x := 0
	go func(p *int) {
		for i := 1; i <= 20000000000; i++ {
			*p = i
			// если закоментить runtime.Gosched() x будет = 20000000000.
			// Gosched позволяют запускать другие горутины если GOMAXPROCS = 1, а в частнгсти, передать вызов горутине main. Из за того что GOMAXPROCS 1
			runtime.Gosched()
			// Число будет любое или 20000000000
		}
	}(&x)

	time.Sleep(5 * time.Second)
	fmt.Printf("x = %d.\n", x)
}

func TestGoroutines7() {
	c := make(chan int)
	d := make(chan struct{})

	go func(ch chan<- int, x int) {
		fmt.Println(x)
		ch <- x * x
	}(c, 3)

	go func(ch <-chan int) {
		n := <-ch
		fmt.Println(n)
		d <- struct{}{}
	}(c)

	<-d
	// 3 9
}
