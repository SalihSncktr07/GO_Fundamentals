package main

import (
	//"golesson/variables"
	"fmt"
	functions "golesson/Functions"
)

func main() {
	//conditionals.Ws()
	functions.SelamVer("Salih")
	var sonuc = functions.Topla(2, 5)
	sonuc *= 10
	fmt.Println("Main çalışınca gelen değer:", sonuc)
}
