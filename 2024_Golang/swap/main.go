package main

import "fmt"

func main() {
	a := 100
	b := 10
	new_a, new_b := swap(a, b)
	fmt.Println(new_a, new_b)

	a2 := 10
	b2 := 100
	new_a, new_b = swap(a2, b2)
	fmt.Println(new_a, new_b)
}

func swap(a, b int) (int, int) {
	if a < b {
		fmt.Println("進入切換")
		return swap(b, a)
	}
	return b, a
}
