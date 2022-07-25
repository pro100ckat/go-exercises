package examples

import "fmt"

func TestMap() {
	scores := map[string]int{
		"anna":   11,
		"andrew": 20,
	}

	foo(scores, "anna")
	fmt.Println("Score:", scores["anna"])
}

func foo(scores map[string]int, player string) {
	scores[player] = scores[player] * 2
}
