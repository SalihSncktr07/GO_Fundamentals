package goroutines

import (
	"fmt"
	"time"
)

func CiftSayi() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Çift sayı:", i)
		time.Sleep(time.Second)
	}
}

func TekSayi() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Tek sayı:", i)
		time.Sleep(time.Second)
	}
}

/*
	go goroutines.TekSayi()
	go goroutines.CiftSayi()
	time.Sleep(5 * time.Second) //5 sn (eğer 5 yazılmasaydı default 1sn olurdu)
	fmt.Println("Main bitti")
*/
