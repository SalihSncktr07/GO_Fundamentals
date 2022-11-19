package goroutines

import (
	"fmt"
	"time"
)

func CiftSayilar(ciftCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Çift sayı toplama çalışıyor:", i)
		time.Sleep(time.Second)
	}

	ciftCn <- toplam
}

func TekSayilar(tekCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Tek sayı:", i)
		time.Sleep(time.Second)
	}

	tekCn <- toplam
}
