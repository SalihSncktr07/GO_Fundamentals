package main

import (
	"fmt"
	"golesson/functions"
)

func main() {
	fmt.Println(functions.Variadic(1, 4, 9, 34))
	fmt.Println(functions.Variadic(1, 4))
	fmt.Println(functions.Variadic(1, 4, 9, 34, 21, 7))
}
