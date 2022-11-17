package main

import (
	"fmt"
	"golesson/Functions"
)

func main() {
	//conditionals.Ws()
	/*
		functions.SelamVer("Salih")
		var sonuc = functions.Topla(2, 5)
		sonuc *= 10
		fmt.Println("Main çalışınca gelen değer:", sonuc)
	*/
	sonuc1, _ , sonuc3, _ := functions.DortIslem(10, 4)
	fmt.Println("Toplam", sonuc1)
	//fmt.Println("Çıkarma", sonuc2)
	fmt.Println("Çarpım", sonuc3)
	//fmt.Println("Bölüm", sonuc4)
}
