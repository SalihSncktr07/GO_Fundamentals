package loops

import "fmt"

func Ws2() {
	sayi := 2
	check := 2

	fmt.Println("Lütfen asal bir sayı giriniz..")
	fmt.Scanln(&sayi)

	if sayi <= 1 {
		fmt.Println("Lütfen asal bir sayı giriniz..")
		fmt.Scanln(&sayi)
	}
	for check <= sayi/2 {
		if sayi%check == 0 {
			fmt.Println("Lütfen asal bir sayı giriniz..")
			fmt.Scanln(&sayi)
		} else {
			check++
		}
	}
	fmt.Println("Girdiğiniz sayı asaldır.")
}
