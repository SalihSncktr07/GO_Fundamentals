package conditionals

import "fmt"

func Ws() {
	var sayi1, sayi2, sayi3 int = 5, 20, 10

	if sayi1 > sayi2 && sayi1 > sayi3 {
		fmt.Println(sayi1)
	} else if sayi2 > sayi1 && sayi2 > sayi3 {
		fmt.Println(sayi2)
	} else {
		fmt.Println(sayi3)
	}
}
