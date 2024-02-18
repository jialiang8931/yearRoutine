package main

import "fmt"

func main() {
	i := 100
	var ptrI *int

	fmt.Println(i)
	fmt.Println(ptrI)
	fmt.Println("============")

	ptrI = &i
	fmt.Println(i)
	fmt.Println(ptrI)
	fmt.Println(*ptrI)
	fmt.Println("============")

	*ptrI = 999
	fmt.Println(i)
	fmt.Println(*ptrI)

}
