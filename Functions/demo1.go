package functions

import "fmt"

func Topla(a int, b int) int /* Dönüş tipi */ {
	sum := a + b
	fmt.Println("Fonk çalışınca gelen değer:", sum)
	return sum
}

func SelamVer(kullanici_adi string) {
	fmt.Println("Werhaba,", kullanici_adi)
}

	/*main
		functions.SelamVer("Salih")
		var sonuc = functions.Topla(2, 5)
		sonuc *= 10
		fmt.Println("Main çalışınca gelen değer:", sonuc)
	*/