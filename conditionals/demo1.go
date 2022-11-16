package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilen float64 = 50.57

	if cekilen > hesap {
		fmt.Println("Hesaptaki para yetersiz.")
		fmt.Println("Hesaptaki para -> ", hesap)
	}

	if cekilen < hesap {
		fmt.Println("Paranız hazırlanıyor...")
		fmt.Println("Hesaptaki para -> ", hesap - cekilen)
	}
}
