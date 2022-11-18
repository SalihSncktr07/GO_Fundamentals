package goroutines

import "fmt"

func CiftSayi() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Çift sayı:", i)
	}
}

func TekSayi() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Tek sayı:", i)
	}
}

/*
go goroutines.CiftSayi()
	go goroutines.TekSayi()
	time.Sleep(5 * time.Second) //5 sn (eğer 5 yazılmasaydı default 1sn olurdu)
	fmt.Println("Main bitti")
*/
