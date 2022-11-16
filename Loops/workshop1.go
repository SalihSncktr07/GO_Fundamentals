package loops

import "fmt"

func Ws1() {
	tahmin := 0
	sayi := 96
	tahmin_sayisi := 0

	fmt.Println("Tahmininizi yazınız.")
	fmt.Scanln(&tahmin)

	for sayi != tahmin {
		if tahmin < sayi {
			fmt.Println("Lütfen daha büyük bir sayı giriniz..")
			fmt.Scanln(&tahmin)
			tahmin_sayisi++
		} else if tahmin > sayi {
			fmt.Println("Lütfen daha küçük bir sayı giriniz..")
			fmt.Scanln(&tahmin)
			tahmin_sayisi++
		}
	}
	fmt.Println("Bravo, doğru sayıyı buldunuz!! Tahmin sayınız:", tahmin_sayisi)
}
