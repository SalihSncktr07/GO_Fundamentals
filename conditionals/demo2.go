package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilen float64 = 50.67

	if hesap > cekilen {
		fmt.Println("Paranız hazırlanıyor...")
		fmt.Println("Hesapta kalan para -> ", hesap - cekilen)
	} else if cekilen == hesap {
		fmt.Println("Paranız hazırlanıyor...")
		fmt.Println("Dikkat hesapta para kalmadı!\nBakiye ->", 0)
	} else {
		fmt.Println("Hesaptaki para yetersiz.")
		fmt.Println("Hesaptaki para -> ", hesap)
	}

	//else if sadece 2 den fazla olan durumlarda kullanılır, eğer sadece iki durum varsa if ve else kullanılır.
}
