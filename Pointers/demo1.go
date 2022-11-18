package pointers

import "fmt"

func Demo1(sayi *int) {
	*sayi += 5
	fmt.Println("Demodaki sayı:", *sayi)

	/*
	sayi := 20
	pointers.Demo1(&sayi)
	fmt.Println("Maindeki sayı: ", sayi)
	*/
}