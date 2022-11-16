package main

import "fmt"

func main() {
	var metin string = "Hello World!"
	var sayi int = 15
	var sayi2 float32 = 1.6
	sayi3 := 30
	var durum bool
	metin1 := "salih"
	metin2 := "salih"
	durum = metin1 == metin2

	fmt.Println("Hello GO!")
	fmt.Println(60)

	fmt.Println(metin)

	fmt.Println(sayi)
	fmt.Println(sayi + 35)

	fmt.Println(sayi2 + 5)
	fmt.Printf("sayi2 -> Veri tipi: %T\n", sayi2)

	fmt.Println(sayi3)
	fmt.Printf("sayi3 -> Veri tipi: %T\n", sayi3)

	fmt.Println("bool türü: ", durum)
}
