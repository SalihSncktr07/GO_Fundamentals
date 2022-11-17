package functions

import "fmt"

func Topla(a int, b int) int {
	sum := a + b
	fmt.Println("Fonk çalışınca gelen değer:", sum)
	return sum
}

func SelamVer(kullanici_adi string) {
	fmt.Println("Werhaba,", kullanici_adi)
}
