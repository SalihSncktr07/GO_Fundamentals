package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	sayi := 50

	if tahmin > 100 || tahmin < 1 {
		return "", errors.New("1 - 100 arasında bir sayı giriniz")
	}

	if tahmin > sayi {
		return "Daha küçük bir sayı giriniz", nil
	}

	if tahmin > sayi {
		return "Daha büyük bir sayı giriniz", nil
	}

	return "Bildiniz", nil //hata(nil), hata yok demek
}

func Demo2() {
	mesaj, _ := TahminEt(50)
	fmt.Println(mesaj)
	fmt.Println(TahminEt(50))
	fmt.Println(TahminEt(101))
	fmt.Println(TahminEt(-10))
}
