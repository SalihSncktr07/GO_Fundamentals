package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Bilgisayar", 10000, "X"})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
