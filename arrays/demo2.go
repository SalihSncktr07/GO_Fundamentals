package arrays

import (
	"fmt"
)

func Demo2() {
	sayilar := [5]int{5, 8, 64, 32, 2}
	enBuyuk := sayilar[0]

	fmt.Println(sayilar)

	for i := 0; i < len(sayilar); i++ {
		if sayilar[i] > enBuyuk {
			enBuyuk = sayilar[i]
		}
	}

	fmt.Println("En Büyük Sayı:", enBuyuk)
}
