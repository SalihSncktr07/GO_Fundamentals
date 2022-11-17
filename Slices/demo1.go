package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)

	fmt.Println(isimler)
	isimler[0] = "salih"
	isimler[1] = "engin"
	isimler[2] = "deniz"
	isimler = append(isimler, "büşra")
	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
