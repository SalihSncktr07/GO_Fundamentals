package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilen float64 = 50.57

	if cekilen > hesap {
		fmt.Print("Hesaptaki para yetersiz.")
		fmt.Print("Hesaptaki para -> ", hesap)
	}

	if cekilen < hesap {
		fmt.Print("Paranız hazırlanıyor...")
		fmt.Print("Hesaptaki para -> ", hesap - cekilen)
	}
}
