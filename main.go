package main

import (
	"fmt"
	"golesson/pointers"
)

func main() {
	sayilar := []int{1, 2, 3}
	pointers.Demo2(sayilar)
	fmt.Println("Maindeki sayı: ", sayilar[0])
}
