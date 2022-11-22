package interfaces

import "fmt"

func verify(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo3() {
	var sayi1 interface{} = 5
	verify(sayi1)

	var sayi2 interface{} = "Salih"
	verify(sayi2)
}
