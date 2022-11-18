package main

import (
	"fmt"
	"golesson/goroutines"
	"time"
)

func main() {
	go goroutines.CiftSayi()
	go goroutines.TekSayi()
	time.Sleep(5 * time.Second) //5 sn (eğer 5 yazılmasaydı default 1sn olurdu)
	fmt.Println("Main bitti")
}
